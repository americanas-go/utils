package strings

import (
	"testing"
)

func TestSliceContainsAll(t *testing.T) {
	type args struct {
		slice []string
		strs  []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "When strs arg is empty",
			args: args{
				slice: []string{"item1"},
				strs:  []string{},
			},
			want: false,
		},
		{
			name: "When slice arg has more items",
			args: args{
				slice: []string{"item1", "item2"},
				strs:  []string{"item1"},
			},
			want: false,
		},
		{
			name: "When strs arg doesn' have the same items",
			args: args{
				slice: []string{"item1", "item2"},
				strs:  []string{"item1", "item3"},
			},
			want: false,
		},
		{
			name: "When slice and strs are empty",
			args: args{
				slice: []string{},
				strs:  []string{},
			},
			want: true,
		},
		{
			name: "When slice is empty",
			args: args{
				slice: []string{},
				strs:  []string{"item1"},
			},
			want: true,
		},
		{
			name: "When slice and strs have the same items with 1 item",
			args: args{
				slice: []string{"item1"},
				strs:  []string{"item1"},
			},
			want: true,
		},
		{
			name: "When slice and strs have the same items with 2 items",
			args: args{
				slice: []string{"item1", "item2"},
				strs:  []string{"item1", "item2"},
			},
			want: true,
		},
		{
			name: "When strs has the same item and more",
			args: args{
				slice: []string{"item1"},
				strs:  []string{"item1", "item2"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContainsAll(tt.args.slice, tt.args.strs); got != tt.want {
				t.Errorf("SliceContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
