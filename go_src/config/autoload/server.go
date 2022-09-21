package autoload

type ServerConfig struct {
	Host string `ini:"host" yaml:"host"`
	Port uint16 `ini:"port" yaml:"port"`
}

var Server = ServerConfig{
	Host: "127.0.0.1",
	Port: 9999,
}
