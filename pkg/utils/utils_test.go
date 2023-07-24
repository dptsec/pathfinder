package utils

import "testing"

func TestParsePath(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantErr bool
	}{
		{"Base URL", args{"https://www.google.com/"}, 1, "/", false},
		{"Base URL no slash", args{"https://www.google.com"}, 1, "/", false},
		{"File", args{"https://www.google.com/search"}, 1, "/", false},
		{"Directory with slash", args{"https://www.google.com/search/"}, 2, "/search/", false},
		{"Directory with params", args{"https://www.google.com/api/v1/search?a=b&c=d"}, 3, "/api/v1/", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParsePath(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathLevels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathLevels() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PathLevels() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInsertParams(t *testing.T) {
	type args struct {
		URL    string
		params string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertParams(tt.args.URL, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Base URL", args{"https://www.google.com/", "/"}, "https://www.google.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
