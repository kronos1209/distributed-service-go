package agent_test

import (
	"github.com/kronos1209/proglog/internal/config"
	"github.com/stretchr/testify/require"
)

func TestAgent(t *testing.T){
	serverTLSConfig,err := config.SetupTLSConfig(config.TLSConfig{
		CertFile: config.ServerCertFile,
		KeyFile: config.ServerKeyFile,
		CAFile: config.CAFile,
		Server: true,
		ServerAddress: "127.0.0.1",
	})
	require.NoError(t,err)

	peerTLSConfig,err := config.SetupTLSConfig(config.TLSConfig{
		CertFile: config.RootClientCertFile,
		KeyFile:config.RootClientkeyFile,
		CAFile: config.CAFile,
		Server: false,
		ServerAddress: "127.0.0.1",
	})
	require.NoError(t,err)

	
}