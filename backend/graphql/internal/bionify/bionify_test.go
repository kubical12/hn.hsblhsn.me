package bionify

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/hsblhsn/hn.hsblhsn.me/featureflags"
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
			if got := word(tt.args.word); got != tt.want {
				t.Errorf("Word() = %q, want %q", got, tt.want)
			}
		})
	}
}

// nolint:paralleltest // t.Setenv is not parallelizable.
func TestFeatureFlag(t *testing.T) {
	t.Setenv(featureflags.FeatureBionify.String(), "on")
	out := Text("quick")
	if out != "<b bionic-bold>qu</b><span bionic-span>ick</span>" {
		t.Errorf("Text() = %q, want %q", out, "<b bionic-bold>qu</b><span bionic-span>ick</span>")
	}
	t.Setenv(featureflags.FeatureBionify.String(), "off")
	out = Text("quick")
	if out != "quick" {
		t.Errorf("Text() = %q, want %q", out, "quick")
	}
	t.Setenv(featureflags.FeatureBionify.String(), "")
	out = Text("quick")
	if out != "quick" {
		t.Errorf("Text() = %q, want %q", out, "quick")
	}
}

func FuzzText(f *testing.F) {
	f.Setenv("FEATURE_FLAG_BIONIFY", "on")
	f.Add("quick brown fox jumps over the lazy dog")
	f.Add("привет мир")
	f.Add("আমি")
	f.Add("শ্রীমতি")
	f.Add("মতিশ্রী")
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		t.Helper()
		out := Text(input)
		if out == "" && input != "" {
			t.Errorf("Text(%q) = %q, want non-empty", input, out)
			return
		}
		_, err := goquery.NewDocumentFromReader(strings.NewReader(out))
		if err != nil {
			t.Errorf("Text(%q) = %q, want valid HTML", input, out)
			return
		}
	})
}
