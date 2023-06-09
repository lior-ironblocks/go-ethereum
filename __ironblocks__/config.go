package __ironblocks__

import "fmt"

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetHost(host string) *Config {
	c.Host = host
	return c
}

func (c *Config) SetPort(port int) *Config {
	c.Port = port
	return c
}

func (c *Config) getRPCURL() string {
	return fmt.Sprintf("http://%s:%d", c.Host, c.Port)
}
