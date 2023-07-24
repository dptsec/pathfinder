package pathfinder

import (
	"fmt"
)

type Config struct {
	Confidence int
	Cookie     string
	Method     string
	OutputFile string
	ProxyURL   string
	Rate       int
	StopError  bool
	Timeout    int
	Threads    int
	UserAgent  string
	Verbose    bool
}

func NewConfig() *Config {
	var config Config

	return &config
}

func (c *Config) Print() {
	fmt.Printf("[*] Configuration:\n")
	fmt.Printf("- Confidence cut-off:\t%d\n", c.Confidence)
	fmt.Printf("- Request method:\t%s\n", c.Method)
	fmt.Printf("- Proxy URL:\t\t%s\n", c.ProxyURL)
	fmt.Printf("- Rate-limiting:\t%d requests/second\n", c.Rate)
	fmt.Printf("- Request timeout:\t%d seconds\n", c.Timeout)
	fmt.Printf("- Total threads:\t%d\n", c.Threads)
	fmt.Printf("- HTTP User-Agent:\t%s\n", c.UserAgent)
	fmt.Printf("- Stop on errors:\t%v\n\n", c.StopError)
}
