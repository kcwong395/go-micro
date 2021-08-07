module github.com/kcwong395/go-micro/shippy-cli-consignment

go 1.16

replace github.com/kcwong395/go-micro/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/kcwong395/go-micro/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
)
