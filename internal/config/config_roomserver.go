package config

type RoomServer struct {
	Matrix *Global `yaml:"-"`

	InternalAPI InternalAPIOptions `yaml:"internal_api"`

	Database DatabaseOptions `yaml:"database"`
}

func (c *RoomServer) Defaults() {
	c.InternalAPI.Listen = "http://localhost:7770"
	c.InternalAPI.Connect = "http://localhost:7770"
	c.Database.Defaults()
	c.Database.ConnectionString = "file:roomserver.db"
	c.InternalAPI.GRPCConnect = "localhost:0"
	c.InternalAPI.GRPCListen = "localhost:0"
}

func (c *RoomServer) Verify(configErrs *ConfigErrors, isMonolith bool) {
	checkURL(configErrs, "room_server.internal_api.listen", string(c.InternalAPI.Listen))
	checkURL(configErrs, "room_server.internal_ap.bind", string(c.InternalAPI.Connect))
	checkNotEmpty(configErrs, "room_server.database.connection_string", string(c.Database.ConnectionString))
	checkNotEmpty(configErrs, "room_server.internal_api.grpc_listen", c.InternalAPI.GRPCListen)
	checkNotEmpty(configErrs, "room_server.internal_api.grpc_connect", c.InternalAPI.GRPCConnect)
}
