module github.com/lyraproj/hiera_terraform

go 1.13

require (
	github.com/hashicorp/terraform v0.13.0
	github.com/lyraproj/dgo v0.4.4
	github.com/lyraproj/dgocty v0.4.4
	github.com/lyraproj/hierasdk v0.4.4
	k8s.io/client-go v11.0.0+incompatible // indirect
)

replace (
	github.com/Azure/go-autorest v11.1.2+incompatible => github.com/Azure/go-autorest v12.1.0+incompatible
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.13+incompatible
	github.com/ugorji/go => github.com/ugorji/go v0.0.0-20181204163529-d75b2dcb6bc8
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
