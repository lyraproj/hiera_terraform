package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/lyraproj/dgo/dgo"
	require "github.com/lyraproj/dgo/dgo_test"
	"github.com/lyraproj/dgo/vf"
	"github.com/lyraproj/hierasdk/register"
	"github.com/lyraproj/hierasdk/routes"
)

func TestLookup_TerraformBackend(t *testing.T) {
	testTerraformPlugin(t, url.Values{`options`: {`{"backend": "local", "config": {"path": "terraform.tfstate"}}`}},
		http.StatusOK,
		vf.Map("testobject", vf.Map("key1", "value1", "key2", "value2"), "test", "value"))
}

func TestLookup_TerraformBackendEmpty(t *testing.T) {
	testTerraformPlugin(t, url.Values{`options`: {`{"backend": "local", "config": {"path": "terraform_empty.tfstate"}}`}},
		http.StatusOK,
		vf.Map())
}

func TestLookup_TerraformBackendErrors(t *testing.T) {
	testTerraformPlugin(t, url.Values{`options`: {`{"backend": "something", "config": {"path": "terraform.tfstate"}}`}},
		http.StatusInternalServerError,
		`unknown backend type "something"`)

	testTerraformPlugin(t, url.Values{`options`: {`{"backend": "local", "config": {"something": "else"}}`}},
		http.StatusInternalServerError,
		`the given configuration is not valid for backend "local"`)
}

// testTerraformPlugin runs the plugin in-process using the "net/http/httptest" package.
func testTerraformPlugin(t *testing.T, query url.Values, expectedStatus int, expectedBody interface{}) {
	t.Helper()
	cw, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(`testdata`)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Chdir(cw)
	}()

	register.Clean()
	register.DataHash(`tf`, TerraformBackendData)
	path := `/data_hash/tf`
	r, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(query) > 0 {
		r.URL.RawQuery = query.Encode()
	}

	rr := httptest.NewRecorder()
	handler, _ := routes.Register()
	handler.ServeHTTP(rr, r)

	status := rr.Code
	if status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	var actualBody dgo.Value
	if status == http.StatusOK {
		expectedType := `application/json`
		actualType := rr.Header().Get(`Content-Type`)
		if expectedType != actualType {
			t.Errorf("handler returned unexpected content path: got %q want %q", actualType, expectedType)
		}
		actualBody, err = vf.UnmarshalJSON(rr.Body.Bytes())
		if err != nil {
			t.Fatal(err)
		}
	} else {
		actualBody = vf.String(strings.TrimSpace(rr.Body.String()))
	}

	// Check the response body is what we expect.
	require.Equal(t, expectedBody, actualBody)
}
