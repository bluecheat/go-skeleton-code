// Golang
//go:generate protoc -I. -I$GOPATH -I../vendor --go_opt=paths=source_relative --go_out=plugins=grpc:./generated --go_opt=paths=source_relative ./vehicle.proto
//go:generate protoc -I. -I$GOPATH  -I../vendor --grpc-gateway_out=logtostderr=true,paths=source_relative,allow_delete_body=true:./generated --swagger_out=logtostderr=true,allow_delete_body=true,repeated_path_param_separator=ssv:./generated ./vehicle.proto

package proto
