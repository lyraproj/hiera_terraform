package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	backendInit "github.com/hashicorp/terraform/backend/init"
	"github.com/hashicorp/terraform/configs/hcl2shim"
	"github.com/lyraproj/hierasdk/hiera"
	"github.com/lyraproj/hierasdk/plugin"
	"github.com/lyraproj/hierasdk/register"
	"github.com/lyraproj/hierasdk/vf"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	register.DataHash(`terraform_backend`, TerraformBackendData)
	plugin.ServeAndExit()
}

var lookupLock sync.Mutex

// TerraformBackendData is a data hash function that returns values from a Terraform backend.
// The config can be any valid Terraform backend configuration.
func TerraformBackendData(ctx hiera.ProviderContext) vf.Data {
	// Hide Terraform's debug messages temporarily. A global mutex is required when doing
	// since only one Go routine can hide and restore at any given time.
	lookupLock.Lock()
	stdOut := log.Writer()
	log.SetOutput(ioutil.Discard)
	defer func() {
		log.SetOutput(stdOut)
		lookupLock.Unlock()
	}()

	backend, ok := ctx.StringOption(`backend`)
	if !ok {
		panic(fmt.Errorf(`missing required provider option 'backend'`))
	}
	workspace, ok := ctx.StringOption(`workspace`)
	if !ok {
		workspace = "default"
	}
	configMap := ctx.Option(`config`)
	if configMap == nil {
		panic(fmt.Errorf(`missing required provider option 'config'`))
	}
	if _, ok := configMap.(vf.Map); !ok {
		panic(fmt.Errorf("%q must be a map", "config"))
	}
	config := convertDataToCty(configMap)

	backendInit.Init(nil)
	f := backendInit.Backend(backend)
	if f == nil {
		panic(fmt.Errorf("unknown backend type %q", backend))
	}
	b := f()
	schema := b.ConfigSchema()
	configVal, err := schema.CoerceValue(config)
	if err != nil {
		panic(fmt.Errorf("the given configuration is not valid for backend %q", backend))
	}
	newVal, diags := b.PrepareConfig(configVal)
	if diags.HasErrors() {
		panic(diags.Err())
	}
	configVal = newVal
	diags = b.Configure(configVal)
	if diags.HasErrors() {
		panic(diags.Err())
	}
	state, err := b.StateMgr(workspace)
	if err != nil {
		panic(err)
	}
	err = state.RefreshState()
	if err != nil {
		panic(err)
	}
	remoteState := state.State()
	mod := remoteState.RootModule()
	output := make(vf.Map)
	for k, os := range mod.OutputValues {
		output[k] = ctx.ToData(hcl2shim.ConfigValueFromHCL2(os.Value))
	}
	return output
}

// convert vf.Data to cty.Value recursively
func convertDataToCty(v vf.Data) cty.Value {
	var cv cty.Value
	switch v := v.(type) {
	case vf.String:
		cv = cty.StringVal(string(v))
	case vf.Int:
		cv = cty.NumberIntVal(int64(v))
	case vf.Float:
		cv = cty.NumberFloatVal(float64(v))
	case vf.Bool:
		cv = cty.BoolVal(bool(v))
	case vf.Slice:
		cvs := make([]cty.Value, len(v))
		for i := range v {
			cvs[i] = convertDataToCty(v[i])
		}
		cv = cty.ListVal(cvs)
	case vf.Map:
		cvs := make(map[string]cty.Value, len(v))
		for k, v := range v {
			cvs[k] = convertDataToCty(v)
		}
		cv = cty.ObjectVal(cvs)
	default:
		cv = cty.NilVal
	}
	return cv
}
