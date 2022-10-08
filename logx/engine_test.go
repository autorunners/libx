package logx

import (
	"log"
	"os"
	"testing"
)

func Test_defaultLogEngine_Output(t *testing.T) {
	engine := newEngine(os.Stdout, "[unit]", log.LstdFlags)
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{"demo1", args{"hello world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine.Output(tt.args.message)
		})
	}
}
