package pathfinder

import (
	"reflect"
	"testing"
)

func TestCreatePayloads(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"Base URL", args{"https://www.example.com"},
			[]string{"https://www.example.com/../", "https://www.example.com/%2e%2e%2f", "https://www.example.com/..%2f", "https://www.example.com/;../"}},
		{"Directories", args{"https://www.example.com/api/v1/test"},
			[]string{"https://www.example.com/../", "https://www.example.com/%2e%2e%2f", "https://www.example.com/..%2f", "https://www.example.com/;../"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreatePayloads(tt.args.URL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePayloads() = %v, want %v", got, tt.want)
			}
		})
	}
}
