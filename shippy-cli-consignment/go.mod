module github.com/kcwong395/go-micro/shippy-cli-consignment

go 1.16

replace (
	github.com/kcwong395/go-micro/shippy-service-consignment => ../shippy-service-consignment
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/kcwong395/go-micro/shippy-service-consignment v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
)
