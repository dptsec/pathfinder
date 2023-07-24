package pathfinder

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	type args struct {
		resp    *http.Response
		runner  *Runner
		request string
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(tt.args.resp, tt.args.runner, tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_ParseBody(t *testing.T) {
	type fields struct {
		Body       []byte
		StatusCode int
		Words      int
		Runner     *Runner
		Headers    http.Header
		Cookies    []string
		Request    string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				Body:       tt.fields.Body,
				StatusCode: tt.fields.StatusCode,
				Words:      tt.fields.Words,
				Runner:     tt.fields.Runner,
				Headers:    tt.fields.Headers,
				Cookies:    tt.fields.Cookies,
				Request:    tt.fields.Request,
			}
			if got := resp.ParseBody(); got != tt.want {
				t.Errorf("Response.ParseBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Distance(t *testing.T) {
	type fields struct {
		Body       []byte
		StatusCode int
		Words      int
		Runner     *Runner
		Headers    http.Header
		Cookies    []string
		Request    string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				Body:       tt.fields.Body,
				StatusCode: tt.fields.StatusCode,
				Words:      tt.fields.Words,
				Runner:     tt.fields.Runner,
				Headers:    tt.fields.Headers,
				Cookies:    tt.fields.Cookies,
				Request:    tt.fields.Request,
			}
			if got := resp.Distance(); got != tt.want {
				t.Errorf("Response.Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCookieDifference(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CookieDifference(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CookieDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Compare(t *testing.T) {
	type fields struct {
		Body       []byte
		StatusCode int
		Words      int
		Runner     *Runner
		Headers    http.Header
		Cookies    []string
		Request    string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				Body:       tt.fields.Body,
				StatusCode: tt.fields.StatusCode,
				Words:      tt.fields.Words,
				Runner:     tt.fields.Runner,
				Headers:    tt.fields.Headers,
				Cookies:    tt.fields.Cookies,
				Request:    tt.fields.Request,
			}
			resp.Compare()
		})
	}
}
