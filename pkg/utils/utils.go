package utils

import (
	"fmt"
	"net/url"
	"strings"
)

/* Take a URL and return the path and number of "levels" to the closest directory:
 *
 * https://www.example.com -> /
 * https://www.example.com/	-> /
 * https://www.example.com/test -> /
 * https://www.example.com/api/v1/search -> /api/v1/
 *
 */
func ParsePath(URL string) (int, string, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return 0, "", err
	}

	if len(u.Path) < 1 {
		u.Path = "/"
	}

	if strings.HasSuffix(u.Path, "/") {
		return strings.Count(u.Path, "/"), u.Path, nil
	}

	index := 0
	for i := 0; i < len(u.Path); i++ {
		if string(u.Path[i]) == "/" {
			index = i + 1
		}
	}

	path := u.Path[:index]
	return strings.Count(path, "/"), path, nil
}

func CheckURL(URL string) (string, error) {
	u, err := url.Parse(URL)
	if u.Scheme != "http" && u.Scheme != "https" {
		URL = "https://" + URL
	}

	_, path, err := ParsePath(URL)
	fmt.Printf("%s%s\n", URL, path)

	return URL, err
}

func InsertParams(URL string, params string) (string, error) {
	var err error

	if len(params) <= 0 {
		return "", err
	}

	query, err := url.ParseQuery(params)
	if err != nil {
		return "", err
	}

	done := make(map[string][]string, len(query))
	for k, v := range query {
		done[k] = v
	}

	val := url.Values{}
	for k := range done {
		for _, value := range query {
			fmt.Printf("%s %s\n", k, value)
			val.Set(value[0], "")
		}
	}
	return "", nil
}

func Trim(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func Add(s, suffix string) string {
	if !strings.HasSuffix(s, "%2f") && !strings.HasSuffix(s, suffix) {
		s = s + suffix
	}
	return s
}
