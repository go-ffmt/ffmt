package ffmt

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	type args struct {
		nested map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			args: args{
				nested: map[string]interface{}{},
			},
			want: map[string]interface{}{},
		},
		{
			args: args{
				nested: map[string]interface{}{
					"a": 1,
				},
			},
			want: map[string]interface{}{
				"a": 1,
			},
		},
		{
			args: args{
				nested: map[string]interface{}{
					"a": 1,
					"bb": map[string]interface{}{
						"ccc": 2,
					},
				},
			},
			want: map[string]interface{}{
				"a":      1,
				"bb.ccc": 2,
			},
		},
		{
			args: args{
				nested: map[string]interface{}{
					"a": 1,
					"bb": []interface{}{
						"cc", 2,
					},
				},
			},
			want: map[string]interface{}{
				"a":    1,
				"bb.0": "cc",
				"bb.1": 2,
			},
		},
		{
			args: args{
				nested: map[string]interface{}{
					"a": 1,
					"bb": []interface{}{
						map[string]interface{}{
							"ddd": 3,
						}, "cc",
					},
				},
			},
			want: map[string]interface{}{
				"a":        1,
				"bb.1":     "cc",
				"bb.0.ddd": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.nested); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
