package mqttx

import (
	"testing"
)

var homePath = ".."

func TestConectDefault(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{"demo", args{homePath+"/testutil/testdata/mqtt/config.yaml"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConectDefault(tt.args.configFile)
		})
	}
}
