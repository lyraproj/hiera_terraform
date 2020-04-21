module github.com/lyraproj/hiera_terraform

go 1.13

require (
	github.com/hashicorp/terraform v0.12.24
	github.com/lyraproj/dgo v0.4.4
	github.com/lyraproj/dgocty v0.4.4
	github.com/lyraproj/hierasdk v0.4.4
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.13+incompatible
	github.com/ugorji/go => github.com/ugorji/go v0.0.0-20181204163529-d75b2dcb6bc8
)
