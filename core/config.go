package core

import "time"

type Setting struct {
	Redis  RedisConfig  `yaml:"redis" json:"redis"`
	Server ServerConfig `yaml:"server" json:"server"`
}

type RedisConfig struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type ServerConfig struct {
	ServerPort   string        `yaml:"server_port" json:"server_port"`
	DebugPort    string        `yaml:"debug_port" json:"debug_port"`
	EnablePprof  bool          `yaml:"enable_pprof" json:"enable_pprof"`
	ReadTimeout  time.Duration `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout"`
	Mode         string        `yaml:"mode" json:"mode"`
	Domain       string        `yaml:"domain" json:"domain"`
}
