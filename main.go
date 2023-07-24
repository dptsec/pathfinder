package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/dptsec/pathfinder/pkg/pathfinder"
	"github.com/dptsec/pathfinder/pkg/utils"
)

func GetFlags(config *pathfinder.Config) {
	flag.IntVar(&config.Confidence, "c", 3, "Confidence level for a positive finding")
	flag.StringVar(&config.Cookie, "C", "", "Cookie data")
	flag.StringVar(&config.Method, "m", "GET", "HTTP method")
	flag.StringVar(&config.OutputFile, "o", "TODO", "Output file")
	flag.StringVar(&config.ProxyURL, "p", "", "Proxy URL")
	flag.BoolVar(&config.StopError, "e", false, "Stop when >=75% of requests have returned an error")
	flag.IntVar(&config.Rate, "r", 50, "Maximum requests per second")
	flag.IntVar(&config.Timeout, "T", 10, "HTTP request timeout (seconds)")
	flag.IntVar(&config.Threads, "t", 5, "Concurrent threads to use")
	flag.StringVar(&config.UserAgent, "u", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36", "User-Agent")
	flag.BoolVar(&config.Verbose, "v", false, "Verbose mode - show all confidence level checks")
	flag.Parse()
}

func main() {
	config := pathfinder.NewConfig()
	GetFlags(config)

	job := pathfinder.NewJob(config)

	config.Print()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		current, err := utils.CheckURL(input.Text())
		if err != nil {
			fmt.Printf("[-] Error: %v\n", err)
			continue
		}

		for _, c := range pathfinder.CreatePayloads(current) {
			job.Queue(c, "")
		}
	}
	if err := job.Start(); err != nil {
		fmt.Printf("[-] job.Start(): %v\n", err)
	}
}
