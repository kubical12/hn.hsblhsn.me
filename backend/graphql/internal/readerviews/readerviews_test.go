package readerviews

import "testing"

func TestToHNLink(t *testing.T) {
	t.Parallel()
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestToHNLink",
			args: args{
				link: "https://news.ycombinator.com/item?id=12345",
			},
			want: "/item?id=12345",
		},
		{
			name: "invalid item link",
			args: args{
				link: "https://news.ycombinator.com/item?id=1234134853845534324234",
			},
			want: "https://news.ycombinator.com/item?id=1234134853845534324234",
		},
		{
			name: "invalid item url link",
			args: args{
				link: "https://news.ycombinator.com/top",
			},
			want: "https://news.ycombinator.com/top",
		},
		{
			name: "invalid item url link",
			args: args{
				link: "https://example.com",
			},
			want: "https://example.com",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got, _ := toHNLink(tt.args.link); got != tt.want {
				t.Errorf("toHNLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
