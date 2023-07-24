package pathfinder

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/agnivade/levenshtein"
)

type Response struct {
	Body       []byte
	StatusCode int
	Words      int
	Runner     *Runner
	Headers    http.Header
	Cookies    []string
	Request    string
}

var defaultHeaders = []string{"Content-Type", "Server", "X-Powered-By", "Location"}

func NewResponse(resp *http.Response, runner *Runner, request string) Response {
	var response Response

	response.Runner = runner
	response.Request = request
	response.StatusCode = resp.StatusCode
	response.Headers = resp.Header
	response.Words = 0
	response.Body, _ = ioutil.ReadAll(resp.Body)
	for _, c := range resp.Cookies() {
		response.Cookies = append(response.Cookies, c.Name)
	}
	resp.Body.Close()

	return response
}

func (resp *Response) ParseBody() bool {
	body := string(resp.Body)
	if len(body) <= 0 {
		return false
	}

	resp.Words = len(strings.Split(body, " "))
	return true
}

func (resp *Response) Distance() int {
	if len(string(resp.Body)) > 0 && len(string(resp.Runner.Base.Body)) > 0 {
		return levenshtein.ComputeDistance(string(resp.Body), string(resp.Runner.Base.Body))
	}

	return -1
}

func CookieDifference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func (resp *Response) Compare() {
	var msg string
	confidence := 0
	base := resp.Runner.Base

	if base == nil || resp == nil {
		return
	}

	if base.StatusCode != resp.StatusCode {
		msg += fmt.Sprintf("[*] StatusCode mismatch:\n\t%d != %d\n", base.StatusCode, resp.StatusCode)
		confidence++
	}

	if resp.ParseBody() {
		if resp.Words != resp.Runner.Base.Words {
			msg += fmt.Sprintf("[*] Levenshtein distance:\n\t%d\n", resp.Distance())
			msg += fmt.Sprintf("[*] Word count:\n\t%d != %d\n", resp.Words, resp.Runner.Base.Words)
			confidence++
		}
	}

	for _, h := range defaultHeaders {
		baseHeaders := base.Headers.Get(h)
		currHeaders := resp.Headers.Get(h)

		if baseHeaders == "" && currHeaders == "" {
			continue
		}

		if baseHeaders != currHeaders {
			confidence++
			msg += fmt.Sprintf("[*] Header mismatch: %s\n\t%s != %s\n", h, baseHeaders, currHeaders)
		}
	}

	baseDiff := CookieDifference(base.Cookies, resp.Cookies)
	currDiff := CookieDifference(resp.Cookies, base.Cookies)
	if len(baseDiff) > 0 {
		msg += fmt.Sprintf("[*] Cookie mismatch (missing):\n\t%v\n", baseDiff)
		confidence++
	}
	if len(currDiff) > 0 {
		msg += fmt.Sprintf("[*] Cookie mismatch (has new):\n\t%v\n", currDiff)
		confidence++
	}

	if confidence >= resp.Runner.Config.Confidence {
		if resp.Runner.Config.Verbose {
			fmt.Printf("%s\n", msg)
		}
		fmt.Printf("[*] Potential hit: %s\n\n", resp.Request)
	}

	return
}
