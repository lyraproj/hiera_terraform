module github.com/lyraproj/hiera_terraform

go 1.13

require (
	github.com/hashicorp/terraform v0.12.9
	github.com/lyraproj/hiera v0.0.0-20190820132249-b08bb003a3b6
	github.com/lyraproj/issue v0.0.0-20190606092846-e082d6813d15
	github.com/lyraproj/pcore v0.0.0-20190918201925-7e14d50f3d7d
	github.com/stretchr/testify v1.3.0
	github.com/zclconf/go-cty v1.1.0
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.13+incompatible
	github.com/lyraproj/hiera => ../hiera
	github.com/lyraproj/pcore => ../pcore
	github.com/ugorji/go => github.com/ugorji/go v0.0.0-20181204163529-d75b2dcb6bc8
)
