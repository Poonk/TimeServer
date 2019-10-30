package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
)

type Server struct {
	Proto string
	Addr  string
}

type Time struct {
	Second int
}

type Config struct {
	Server *Server
	Time   *Time
}

var (
	confPath string
	Conf     *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "TimeServer.toml", "config path")
}

func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
