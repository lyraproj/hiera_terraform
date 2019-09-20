# Terraform Hiera Backend

This function allows hiera to query data from a Terraform [backend](https://www.terraform.io/docs/backends/types/index.html).
Any backend supported by Terraform can be queried.

## Installation
Build the plugin from the root directory of this module:
```
go build -buildmode=plugin -o terraform_backend.so
```
Then make the plugin available to Hiera. See
[Extending Hiera](https://github.com/lyraproj/hiera#Extending-Hiera) for info on how to do that.

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

* `backend` - The name of the backend.

* `config` - A map of options to configure the backend. See the Terraform documentation for each backend.

* `workspace` - The name of the workspace. If not set the `default` workspace will be used.

If the backend supports reading options from environment variables this will work as well.
