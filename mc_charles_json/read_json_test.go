package mc_charles_json

import (
	"testing"
)

func TestReadJson(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{"./charles.json"}},
		//{"2", args{"./test.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadJson(tt.args.fileName)
		})
	}
}
