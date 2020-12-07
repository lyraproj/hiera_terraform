module github.com/lyraproj/hiera_terraform

go 1.13

require (
	github.com/hashicorp/terraform v0.14.0
	github.com/lyraproj/dgo v0.4.4
	github.com/lyraproj/dgocty v0.4.4
	github.com/lyraproj/hierasdk v0.4.4
	k8s.io/client-go v11.0.0+incompatible // indirect
)

replace (
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.27.1
)
