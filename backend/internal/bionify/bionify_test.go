package bionify

import (
	"testing"
)

func TestWord(t *testing.T) {
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
			want: "<b>a</b><span>n</span>",
		},
		{
			name: "word with length 3",
			args: args{
				word: "the",
			},
			want: "<b>t</b><span>he</span>",
		},
		{
			name: "word with length 8",
			args: args{
				word: "absolute",
			},
			want: "<b>abs</b><span>olute</span>",
		},
		{
			name: "word with unicode",
			args: args{
				word: "привет",
			},
			want: "<b>пр</b><span>ивет</span>",
		},
		{
			name: "word with bengali unicode",
			args: args{
				word: "আমি",
			},
			want: "<b>আ</b><span>মি</span>",
		},
		{
			name: "word with bengali multichar unicode",
			args: args{
				word: "শ্রীমতি",
			},
			want: "<b>শ্রী</b><span>মতি</span>",
		},
		{
			name: "word with bengali multichar unicode at the end",
			args: args{
				word: "মতিশ্রী",
			},
			want: "<b>মতি</b><span>শ্রী</span>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Word(tt.args.word); got != tt.want {
				t.Errorf("Word() = %v, want %v", got, tt.want)
			}
		})
	}
}
