package pathfinder

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Runner struct {
	Base   *Response
	Config *Config
	Client *http.Client
	Ready  bool
}

func NewRunner(config *Config) *Runner {
	var runner Runner
	proxyURL := http.ProxyFromEnvironment

	runner.Base = nil
	runner.Ready = false
	runner.Config = config
	timeout := time.Duration(config.Timeout) * time.Second
	if len(config.ProxyURL) > 0 {
		proxy, err := url.Parse(config.ProxyURL)
		if err == nil {
			proxyURL = http.ProxyURL(proxy)
		}
	}

	runner.Client = &http.Client{
		Timeout:       timeout,
		CheckRedirect: nil,
		Transport: &http.Transport{
			Proxy:               proxyURL,
			MaxConnsPerHost:     500,
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 500,
			ForceAttemptHTTP2:   false,
			DialContext: (&net.Dialer{
				Timeout: timeout,
			}).DialContext,
			TLSHandshakeTimeout: timeout,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				MinVersion:         tls.VersionTLS10,
				Renegotiation:      tls.RenegotiateOnceAsClient,
			},
		}}

	return &runner
}

func (runner *Runner) Baseline(URL string) error {
	var err error
	if runner.Base == nil {
		p, _ := url.Parse(URL)
		if runner.Base, err = runner.Fetch(p.Scheme + "://" + p.Host); err != nil {
			return err
		}
		runner.Base.ParseBody()
	}
	return nil
}

func (runner *Runner) CheckReady(URL string) error {
	if !runner.Ready {
		if err := runner.Baseline(URL); err != nil {
			return err
		}
		runner.Ready = true
	}
	return nil
}

func (runner *Runner) Fetch(URL string) (*Response, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(runner.Config.Timeout)*time.Second)

	req, err := http.NewRequestWithContext(ctx, runner.Config.Method, URL, nil)
	if err != nil {
		return nil, err
	}

	if len(runner.Config.Cookie) > 0 {
		req.Header.Set("Cookie", runner.Config.Cookie)
	}

	req.Header.Set("User-Agent", runner.Config.UserAgent)

	r, err := runner.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)

	resp := NewResponse(r, runner, URL)
	return &resp, err
}
