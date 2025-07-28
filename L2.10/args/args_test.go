package args

import (
	"reflect"
	"testing"
)

func TestParseFlagsFromArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Flags
		wantErr bool
	}{
		{
			name: "Only -n",
			args: []string{"-n"},
			want: Flags{
				SortByNumber: true,
				Column:       -1,
			},
		},
		{
			name: "Only -r",
			args: []string{"-r"},
			want: Flags{
				ReverseSort: true,
				Column:      -1,
			},
		},
		{
			name: "Only -u",
			args: []string{"-u"},
			want: Flags{
				UniqueSort: true,
				Column:     -1,
			},
		},
		{
			name: "Multiple short flags combined",
			args: []string{"-nru"},
			want: Flags{
				SortByNumber: true,
				ReverseSort:  true,
				UniqueSort:   true,
				Column:       -1,
			},
		},
		{
			name: "Sort by column with number",
			args: []string{"-k", "2"},
			want: Flags{
				SortByColumn: true,
				Column:       2,
			},
		},
		{
			name:    "Sort by column with missing number",
			args:    []string{"-k"},
			wantErr: false, // возвращается -1, ошибки нет
			want: Flags{
				SortByColumn: true,
				Column:       -1,
			},
		},
		{
			name:    "Sort by column with non-number",
			args:    []string{"-k", "abc"},
			wantErr: true,
		},
		{
			name: "All flags combined",
			args: []string{"-n", "-r", "-k", "3", "-u"},
			want: Flags{
				SortByNumber: true,
				ReverseSort:  true,
				SortByColumn: true,
				UniqueSort:   true,
				Column:       3,
			},
		},
		{
			name: "All flags combined in string",
			args: []string{"-nrku", "3"},
			want: Flags{
				SortByNumber: true,
				ReverseSort:  true,
				SortByColumn: true,
				UniqueSort:   true,
				Column:       3,
			},
		},
		{
			name: "May rewrite col or not",
			args: []string{"-nrku", "3", "2"},
			want: Flags{
				SortByNumber: true,
				ReverseSort:  true,
				SortByColumn: true,
				UniqueSort:   true,
				Column:       3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlagsFromArgs() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlagsFromArgs() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
