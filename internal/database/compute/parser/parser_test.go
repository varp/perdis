package parser

import (
	"log/slog"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.vardan.dev/perdis/internal/test"
)

func TestParser_NewParser(t *testing.T) {
	t.Parallel()

	type args struct {
		logger *slog.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *Parser
		wantErr bool
	}{
		{
			name: "passing non-nil logger",
			args: args{
				logger: test.NewNopLogger(),
			},
			want: &Parser{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
		},

		{
			name: "passing nil logger",
			args: args{
				logger: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewParser(tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.IsType(t, tt.want, got) {
				t.Errorf("NewParser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseQuery(t *testing.T) {
	t.Parallel()

	type fields struct {
		logger *slog.Logger
		sm     *stateMachine
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "passing valid query",
			fields: fields{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
			args:    args{"GET key1"},
			want:    []string{"GET", "key1"},
			wantErr: false,
		},

		{
			name: "passing valid command and invalid args",
			fields: fields{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
			args:    args{"SET key1"},
			want:    []string{"SET", "key1"},
			wantErr: false,
		},

		{
			name: "passing invalid command",
			fields: fields{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
			args: args{"REPLACE key1 newValue"},
			want: []string{"REPLACE", "key1", "newValue"},
		},

		{
			name: "passing query with additional spaces",
			fields: fields{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
			args: args{"  SET    key1  	newValue		"},
			want: []string{"SET", "key1", "newValue"},
		},

		{
			name: "passing query with new lines",
			fields: fields{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
			args: args{`  SET
							key1
								newValue		`},
			want: []string{"SET", "key1", "newValue"},
		},

		{
			name: "passing query with invalid characters in key",
			fields: fields{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
			args:    args{"REPLACE k+ey1 newValue"},
			want:    nil,
			wantErr: true,
		},

		{
			name: "passing query with invalid characters in value",
			fields: fields{
				logger: test.NewNopLogger(),
				sm:     newStateMachine(),
			},
			args:    args{"REPLACE key1 new=Value"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Parser{
				logger: tt.fields.logger,
				sm:     tt.fields.sm,
			}
			got, err := p.ParseQuery(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}
