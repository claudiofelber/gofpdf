package gofpdf

import (
	"reflect"
	"testing"
)

func Test_splitPath(t *testing.T) {
	tests := []struct {
		path  string
		parts []string
	}{
		{
			"M14.7.998c.006-1.33 1.999-1.33 1.997-.003.002 1.333-1.997 1.333-1.997.003z",
			[]string{"M", "14.7", ".998", "c", ".006", "-1.33", "1.999", "-1.33", "1.997", "-.003", ".002", "1.333", "-1.997", "1.333", "-1.997", ".003", "z"},
		},
	}
	for _, tt := range tests {
		parts := splitPath(tt.path)
		if !reflect.DeepEqual(parts, tt.parts) {
			t.Errorf("splotPath(%s) = %v, want %v", tt.path, parts, tt.parts)
		}
	}
}
