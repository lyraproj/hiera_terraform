# Terraform Hiera Backend

This function allows hiera to query data from a Terraform [backend](https://www.terraform.io/docs/backends/types/index.html).
Any backend supported by Terraform can be queried.

## Installation
Build the plugin from the root directory of this module:
```
go build -o terraform_backend
```
Then make the plugin available to Hiera. See
[Extending Hiera](https://github.com/lyraproj/hiera#Extending-Hiera) for info on how to do that.

#### A Note about debugging
When debugging remotely from an IDE like JetBrains goland, use `-gcflags 'all=N -l'` to ensure that all symbols are present in the
final binary.
```
go build -o terraform_backend -gcflags 'all=-N -l'
```

## Examples
Example using a local backend:

    ---
    version: 5
    defaults:
      datadir: hiera
      data_hash: yaml_data

    hierarchy:
    - name: common
      path: common.yaml
    - name: terraform_backend_local
      data_hash: terraform_backend
      options:
        backend: local
        config:
          path: /tfdir/terraform.tfstate

Example using a remote S3 backend:

    ---
    version: 5
    defaults:
      datadir: hiera
      data_hash: yaml_data

    hierarchy:
    - name: common
      path: common.yaml
    - name: terraform_backend_s3
      data_hash: terraform_backend
      options:
        backend: s3
        config:
          bucket: mybucket
          key: path/to/my/key
          region: us-east-1

Example using a root key:

    ---
    version: 5
    defaults:
      datadir: hiera
      data_hash: yaml_data

    hierarchy:
    - name: common
      path: common.yaml
    - name: terraform_backend_local
      data_hash: terraform_backend
      root_key: terraform
      options:
        backend: local
        config:
          path: /tfdir/terraform.tfstate

    ---
    some_value: "%{lookup('terraform.some_output')}"

* `backend` - The name of the backend.

* `config` - A map of options to configure the backend. See the Terraform documentation for each backend.

* `workspace` - The name of the workspace. If not set the `default` workspace will be used.

* `root_key` - If set then the state outputs will be wrapped in a map with this value as the root key.

If the backend supports reading options from environment variables this will work as well.
