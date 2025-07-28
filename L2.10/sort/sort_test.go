package sort

import (
	"L2-10/args"
	"reflect"
	"testing"
)

func TestLines(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		flags args.Flags
		want  []string
	}{
		{
			name:  "Empty input",
			lines: []string{},
			flags: args.Flags{},
			want:  []string{},
		},
		{
			name:  "Single line",
			lines: []string{"onlyline"},
			flags: args.Flags{},
			want:  []string{"onlyline"},
		},
		{
			name:  "Sort lexicographically",
			lines: []string{"b", "a", "c"},
			flags: args.Flags{},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "Sort reversed lexicographically",
			lines: []string{"b", "a", "c"},
			flags: args.Flags{
				ReverseSort: true,
			},
			want: []string{"c", "b", "a"},
		},
		{
			name:  "Sort numerically",
			lines: []string{"10", "2", "1"},
			flags: args.Flags{
				SortByNumber: true,
			},
			want: []string{"1", "2", "10"},
		},
		{
			name:  "Sort numerically reversed",
			lines: []string{"10", "2", "1"},
			flags: args.Flags{
				SortByNumber: true,
				ReverseSort:  true,
			},
			want: []string{"10", "2", "1"},
		},
		{
			name:  "Sort by column",
			lines: []string{"a\t2", "a\t1", "a\t3"},
			flags: args.Flags{
				SortByColumn: true,
				Column:       2,
			},
			want: []string{"a\t1", "a\t2", "a\t3"},
		},
		{
			name:  "Sort by column numeric",
			lines: []string{"a\t10", "a\t2", "a\t1"},
			flags: args.Flags{
				SortByColumn: true,
				Column:       2,
				SortByNumber: true,
			},
			want: []string{"a\t1", "a\t2", "a\t10"},
		},
		{
			name:  "Sort unique",
			lines: []string{"a", "b", "a", "c", "b"},
			flags: args.Flags{
				UniqueSort: true,
			},
			want: []string{"a", "b", "c"},
		},
		{
			name:  "All flags combined",
			lines: []string{"a\t10", "a\t2", "a\t1", "a\t2"},
			flags: args.Flags{
				SortByColumn: true,
				Column:       2,
				SortByNumber: true,
				UniqueSort:   true,
				ReverseSort:  false,
			},
			want: []string{"a\t1", "a\t2", "a\t10"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Lines(tt.lines, tt.flags)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lines() = %v, want %v", got, tt.want)
			}
		})
	}
}
