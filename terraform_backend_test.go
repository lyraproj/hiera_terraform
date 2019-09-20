package main

import (
	"errors"
	"os"
	"testing"

	"github.com/lyraproj/hiera/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookup_TerraformBackend(t *testing.T) {
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
	inTestdata(func() {
		_, err := cli.ExecuteLookup(`--var`, `backend:something`, `--config`, `terraform_backend.yaml`, `test`)
		if assert.Error(t, err) {
			require.Equal(t, errors.New(`unknown backend type "something"`), err)
		}
	})
	inTestdata(func() {
		_, err := cli.ExecuteLookup(`--var`, `backend:local`, `--var`, `path:something`, `--config`, `terraform_backend.yaml`, `test`)
		if assert.Error(t, err) {
			require.Equal(t, errors.New(`RootModule called on nil State`), err)
		}
	})
	inTestdata(func() {
		_, err := cli.ExecuteLookup(`--var`, `backend:local`, `--config`, `terraform_backend_errors.yaml`, `test`)
		if assert.Error(t, err) {
			require.Equal(t, errors.New(`the given configuration is not valid for backend "local"`), err)
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
