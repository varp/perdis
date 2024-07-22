package memory

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.vardan.dev/perdis/internal/test"
	"log/slog"
	"testing"
)

func TestNewEngine(t *testing.T) {
	t.Parallel()

	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *Engine
		wantErr bool
	}{
		{
			name:    "passing nil logger",
			args:    args{logger: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "passing non-nil logger",
			args: args{logger: test.NewNopLogger()},
			want: &Engine{logger: test.NewNopLogger()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewEngine(tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEngine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.IsType(t, tt.want, got) {
				t.Errorf("NewEngine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_Set(t *testing.T) {
	t.Parallel()

	engine, _ := NewEngine(test.NewNopLogger())
	engine.Set("key", "value")
	require.Equal(t, engine.data["key"], "value")
}

func TestEngine_Get(t *testing.T) {
	t.Parallel()

	engine, _ := NewEngine(test.NewNopLogger())
	engine.Set("key", "value")
	require.Equal(t, engine.data["key"], "value")
}

func TestEngine_Del(t *testing.T) {
	t.Parallel()

	engine, _ := NewEngine(test.NewNopLogger())
	engine.Set("key", "value")
	engine.Del("key")

	value, found := engine.data["key"]
	require.Equal(t, value, "")
	require.Equal(t, found, false)
}
