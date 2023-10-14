package main

import (
	def_log "log"
	"net"
	"os"

	"github.com/kronos1209/proglog/internal/auth"
	"github.com/kronos1209/proglog/internal/config"
	"github.com/kronos1209/proglog/internal/log"
	"github.com/kronos1209/proglog/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:35825")
	if err != nil {
		def_log.Fatal(err)
	}

	serverTLSConfig, err := config.SetupTLSConfig(config.TLSConfig{
		CertFile:      config.ServerCertFile,
		KeyFile:       config.ServerKeyFile,
		CAFile:        config.CAFile,
		ServerAddress: l.Addr().String(),
		Server:        true,
	})
	if err != nil {
		def_log.Fatal(err)
	}

	serverCreds := credentials.NewTLS(serverTLSConfig)
	dir, err := os.MkdirTemp("", "server-test")
	if err != nil {
		def_log.Fatal(err)
	}
	clog, err := log.NewLog(dir, log.Config{})
	if err != nil {
		def_log.Fatal(err)
	}

	authorizer := auth.New(config.ACLModelFile, config.ACLPolicyFile)
	cfg := &server.Config{
		CommitLog:  clog,
		Authorizer: authorizer,
	}

	server, err := server.NewGRPCServer(cfg, grpc.Creds(serverCreds))
	if err != nil {
		def_log.Fatal(err)
	}
	if err := server.Serve(l); err != nil {
		def_log.Fatal(err)
	}
}
