module github.com/lyraproj/hiera_terraform

go 1.13

require (
	github.com/hashicorp/terraform v0.12.9
	github.com/lyraproj/hiera v0.0.0-20190820132249-b08bb003a3b6
	github.com/lyraproj/hierasdk v0.0.0-20191002210033-6ab2cc3bcf0e
	github.com/stretchr/testify v1.3.0
	github.com/zclconf/go-cty v1.1.0
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.13+incompatible
	github.com/lyraproj/hiera => github.com/thallgren/hiera v0.0.0-20191002212354-5dabab12ab67
	github.com/ugorji/go => github.com/ugorji/go v0.0.0-20181204163529-d75b2dcb6bc8
)
