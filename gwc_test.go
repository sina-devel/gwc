package gwc

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "new",
			args: args{
				config: Config{
					Filenames: []string{".", "go.mod", "go.m"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = New(tt.args.config)
		})
	}
}
