package html

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Hello", "<h1>Hello</h1>"},
		{"Good Bye!", "<h1>Good Bye!</h1>"},
		{"ハイ", "<h1>ハイ</h1>"},
		{"<b>Hello</b>", "<h1><b>Hello</b></h1>"},
		{"", "<h1></h1>"},
	}

	for _, c := range tests {
		got := Headline(c.s)
		if got != c.want {
			t.Errorf("Headline(%s) == %s, want %s", c.s, got, c.want)
		}
	}
}
