package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"EmployeeLearning/Employee/protos/gen"

	"EmployeeLearning/Employee/server/controller"
	"EmployeeLearning/Employee/server/db"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func RunGrpcServer() {
	fmt.Println("Starting GRPC Server...")

	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}

	// GRPC Server
	grpcserver := grpc.NewServer()
	gen.RegisterEmployeeServiceServer(grpcserver, &controller.EmployeeServiceServerImplementation{})

	if err := grpcserver.Serve(listen); err != nil {
		log.Fatal(err)
	}

	fmt.Println("GRPC Server Serving Request on localhost:9090")
}

func RunGatewayServer() {
	fmt.Println("Starting Gateway Server...")

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Marshaling and unmarshling request of http request
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	err = gen.RegisterEmployeeServiceHandlerServer(ctx, grpcMux, &controller.EmployeeServiceServerImplementation{})
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	err = http.Serve(listen, mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gateway Server Serving Request on localhost:8080")
}
func main() {

	// Connection to db
	db.ConnectDB()
	go RunGrpcServer()
	RunGatewayServer()

}
