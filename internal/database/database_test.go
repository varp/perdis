package database

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.vardan.dev/perdis/internal/database/compute/analyzer"
	"go.vardan.dev/perdis/internal/database/compute/parser"
	"go.vardan.dev/perdis/internal/database/storage/memory"
	"go.vardan.dev/perdis/internal/test"
)

func TestStart(t *testing.T) {
	type args struct {
		logger *slog.Logger
	}

	nopLogger := test.NewNopLogger()

	e, _ := memory.NewEngine(nopLogger)
	p, _ := parser.NewParser(nopLogger)
	a, _ := analyzer.NewAnalyzer(nopLogger)

	tests := []struct {
		name    string
		args    args
		want    *Database
		wantErr bool
	}{
		{
			name: "passing non-nil logger",
			args: args{logger: nopLogger},
			want: &Database{
				logger:   nopLogger,
				engine:   e,
				parser:   p,
				analyzer: a,
			},
		},

		{
			name:    "passing nil logger",
			args:    args{logger: nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Start(tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.IsType(t, tt.want, got) {
				t.Errorf("Start() got = %v, want %v", got, tt.want)
			}
		})
	}
}
