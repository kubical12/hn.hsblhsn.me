package bionify

import (
	"testing"
)

// nolint:funlen // test cases are long.
func TestWord(t *testing.T) {
	t.Parallel()
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "word with length 1",
			args: args{
				word: "a",
			},
			want: "a",
		},
		{
			name: "word with length 2",
			args: args{
				word: "an",
			},
			want: "<b bionic-bold>a</b><span bionic-span>n</span>",
		},
		{
			name: "word with length 3",
			args: args{
				word: "the",
			},
			want: "<b bionic-bold>t</b><span bionic-span>he</span>",
		},
		{
			name: "word with length 8",
			args: args{
				word: "absolute",
			},
			want: "<b bionic-bold>abs</b><span bionic-span>olute</span>",
		},
		{
			name: "word with unicode",
			args: args{
				word: "привет",
			},
			want: "<b bionic-bold>пр</b><span bionic-span>ивет</span>",
		},
		{
			name: "word with bengali unicode",
			args: args{
				word: "আমি",
			},
			want: "<b bionic-bold>আ</b><span bionic-span>মি</span>",
		},
		{
			name: "word with bengali multichar unicode",
			args: args{
				word: "শ্রীমতি",
			},
			want: "<b bionic-bold>শ্রী</b><span bionic-span>মতি</span>",
		},
		{
			name: "word with bengali multichar unicode at the end",
			args: args{
				word: "মতিশ্রী",
			},
			want: "<b bionic-bold>মতি</b><span bionic-span>শ্রী</span>",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Word(tt.args.word); got != tt.want {
				t.Errorf("Word() = %v, want %v", got, tt.want)
			}
		})
	}
}
