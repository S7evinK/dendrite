package config

type EDUServer struct {
	Matrix *Global `yaml:"-"`

	InternalAPI InternalAPIOptions `yaml:"internal_api"`
}

func (c *EDUServer) Defaults() {
	c.InternalAPI.Listen = "http://localhost:7778"
	c.InternalAPI.Connect = "http://localhost:7778"
	c.InternalAPI.GRPCConnect = "localhost:9090"
	c.InternalAPI.GRPCListen = "localhost:9090"
}

func (c *EDUServer) Verify(configErrs *ConfigErrors, isMonolith bool) {
	checkURL(configErrs, "edu_server.internal_api.listen", string(c.InternalAPI.Listen))
	checkURL(configErrs, "edu_server.internal_api.connect", string(c.InternalAPI.Connect))
	checkNotEmpty(configErrs, "edu_server.internal_api.grpc_listen", c.InternalAPI.GRPCListen)
	checkNotEmpty(configErrs, "edu_server.internal_api.grpc_connect", c.InternalAPI.GRPCConnect)
}
