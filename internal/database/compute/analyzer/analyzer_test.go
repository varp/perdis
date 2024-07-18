package analyzer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.vardan.dev/perdis/internal/test"
	"log/slog"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	t.Parallel()

	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *Analyzer
		wantErr bool
	}{
		{
			name: "passing non-nil logger",
			args: args{logger: test.NewNopLogger()},
			want: &Analyzer{logger: test.NewNopLogger()},
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
			t.Parallel()

			got, err := NewAnalyzer(tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAnalyzer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.IsType(t, tt.want, got) {
				t.Errorf("NewAnalyzer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalyzer_Analyze(t *testing.T) {
	t.Parallel()

	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		queryTokens []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Query
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "passing valid GET query",
			fields:  fields{logger: test.NewNopLogger()},
			args:    args{queryTokens: []string{"GET", "key1"}},
			want:    Query{commandId: CommandGetID, args: []string{"key1"}},
			wantErr: assert.NoError,
		},
		{
			name:    "passing valid DEL query",
			fields:  fields{logger: test.NewNopLogger()},
			args:    args{queryTokens: []string{"DEL", "key1"}},
			want:    Query{commandId: CommandDelID, args: []string{"key1"}},
			wantErr: assert.NoError,
		},
		{
			name:    "passing valid SET query",
			fields:  fields{logger: test.NewNopLogger()},
			args:    args{queryTokens: []string{"SET", "key1", "value"}},
			want:    Query{commandId: CommandSetID, args: []string{"key1", "value"}},
			wantErr: assert.NoError,
		},

		{
			name:    "passing valid command with invalid args",
			fields:  fields{logger: test.NewNopLogger()},
			args:    args{queryTokens: []string{"SET", "key1"}},
			want:    Query{},
			wantErr: assert.Error,
		},
		{
			name:   "passing invalid command",
			fields: fields{logger: test.NewNopLogger()},
			args:   args{queryTokens: []string{"REPLACE", "key1", "newValue"}},
			want:   Query{},
			wantErr: func(t assert.TestingT, err error, args ...interface{}) bool {
				return assert.ErrorIs(
					t,
					err,
					errAnalyzerUnknownCommand,
					"error must be errAnalyzerUnknownCommand",
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := &Analyzer{
				logger: tt.fields.logger,
			}
			got, err := a.Analyze(tt.args.queryTokens)
			if !tt.wantErr(t, err, fmt.Sprintf("Analyze(%v)", tt.args.queryTokens)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Analyze(%v)", tt.args.queryTokens)
		})
	}
}
