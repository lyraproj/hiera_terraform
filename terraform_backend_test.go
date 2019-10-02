package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/lyraproj/hiera/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookup_TerraformBackend(t *testing.T) {
	ensureTestPlugin(t)
	inTestdata(func() {
		result, err := cli.ExecuteLookup(`--var`, `backend:local`, `--var`, `path:terraform.tfstate`, `--config`, `terraform_backend.yaml`, `test`)
		require.NoError(t, err)
		require.Equal(t, "value\n", string(result))
	})
	inTestdata(func() {
		result, err := cli.ExecuteLookup(`--var`, `backend:local`, `--var`, `path:terraform.tfstate`, `--config`, `terraform_backend.yaml`, `--render-as`, `json`, `testobject`)
		require.NoError(t, err)
		require.Equal(t, `{"key1":"value1","key2":"value2"}`, string(result))
	})
}

func TestLookup_TerraformBackendErrors(t *testing.T) {
	ensureTestPlugin(t)
	inTestdata(func() {
		_, err := cli.ExecuteLookup(`--var`, `backend:something`, `--config`, `terraform_backend.yaml`, `test`)
		if assert.Error(t, err) {
			require.Regexp(t, `unknown backend type "something"`, err.Error())
		}
	})
	inTestdata(func() {
		_, err := cli.ExecuteLookup(`--var`, `backend:local`, `--var`, `path:something`, `--config`, `terraform_backend.yaml`, `test`)
		if assert.Error(t, err) {
			require.Regexp(t, `RootModule called on nil State`, err.Error())
		}
	})
	inTestdata(func() {
		_, err := cli.ExecuteLookup(`--var`, `backend:local`, `--config`, `terraform_backend_errors.yaml`, `test`)
		if assert.Error(t, err) {
			require.Regexp(t, `the given configuration is not valid for backend "local"`, err.Error())
		}
	})
}

func inTestdata(f func()) {
	cw, err := os.Getwd()
	if err == nil {
		err = os.Chdir(`testdata`)
		if err == nil {
			defer func() {
				_ = os.Chdir(cw)
			}()
			f()
		}
	}
	if err != nil {
		panic(err)
	}
}

var once = sync.Once{}

func ensureTestPlugin(t *testing.T) {
	once.Do(func() {
		t.Helper()
		pe := `terraform_backend`
		ps := pe + `.go`
		if runtime.GOOS == `windows` {
			pe += `.exe`
		}

		cmd := exec.Command(`go`, `build`, `-o`, filepath.Join(`testdata`, `plugin`, pe), ps)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
