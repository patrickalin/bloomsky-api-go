package bloomskyStructure

import (
	"strings"
	"testing"
)

func Test_funcName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Ok", "bloomsky-api-go.Test_funcName"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcName(); !strings.Contains(got, tt.want) {
				t.Errorf("funcName() = %v, want %v", got, tt.want)
			}
		})
	}
}
