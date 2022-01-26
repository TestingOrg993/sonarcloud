package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name  string
		argsa int32
		argsb int32
		want  int32
	}{
		{
			name:  "a+b",
			argsa: 20,
			argsb: 30,
			want:  50,
		},
		{
			name:  "a minus",
			argsa: -3,
			argsb: 30,
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.argsa, tt.argsb)
			assert.Equal(t, tt.want, got)
		})
	}
}
