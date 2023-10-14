package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	api "github.com/kronos1209/proglog/api/v1"
	"github.com/kronos1209/proglog/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	newClient := func(crtPath, keyPath string) (
		*grpc.ClientConn,
		api.LogClient,
		[]grpc.DialOption,
	) {
		tlsConfig, err := config.SetupTLSConfig(config.TLSConfig{
			CertFile: crtPath,
			KeyFile:  keyPath,
			CAFile:   config.CAFile,
			Server:   false,
		})
		if err != nil {
			log.Fatal(err)
		}
		tlsCreds := credentials.NewTLS(tlsConfig)
		opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCreds)}
		conn, err := grpc.Dial("127.0.0.1:35825", opts...)
		if err != nil {
			log.Fatal(err)
		}
		client := api.NewLogClient(conn)
		return conn, client, opts
	}
	rootConn, rootClient, _ := newClient(
		config.RootClientCertFile,
		config.RootClientkeyFile,
	)
	defer rootConn.Close()
	rRootClient := reflect.ValueOf(rootClient)
	callFunc := rRootClient.MethodByName("Produce")

	res := callFunc.Call([]reflect.Value{reflect.ValueOf(context.TODO()), reflect.ValueOf(&api.ProduceRequest{
		Record: &api.Record{
			Value: []byte("hello world!"),
		},
	})})
	fmt.Println(res[0])
}
