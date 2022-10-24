package game

import "testing"

func Test_values(t *testing.T) {

	cases := []struct {
		input string
		want  []int
	}{
		{"", []int{}},
		{"6895", []int{6895}},
		{"-6895", []int{6895}},
		{"62+988", []int{62, 988}},
		{"62-988", []int{62, 988}},
		{"62/988", []int{62, 988}},
		{"62*988", []int{62, 988}},
		{"62+3-1", []int{62, 3, 1}},
		{"62*3/1", []int{62, 3, 1}},
	}
	for _, test := range cases {
		got := values(test.input)
		if !sliceIntEquals(got, test.want) {
			t.Fatalf("Expected %d, got %d", test.want, got)
		}
	}
}

func Test_operators(t *testing.T) {

	cases := []struct {
		input string
		want  []string
	}{
		{"", []string{}},
		{"6895", []string{}},
		{"62+988", []string{"+"}},
		{"62-988", []string{"-"}},
		{"62/988", []string{"/"}},
		{"62*988", []string{"*"}},
		{"62+3-1", []string{"+", "-"}},
		{"62*3/1", []string{"*", "/"}},
	}
	for _, test := range cases {
		got := operators(test.input)
		if !sliceStringEquals(got, test.want) {
			t.Fatalf("Expected %s, got %s", test.want, got)
		}
	}
}

func Test_eval(t *testing.T) {

	cases := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"1", 0},
		{"1+1", 2},
		{"42*2", 84},
		{"42/2", 21},
		{"42-2", 40},
		{"42+2", 44},
		{"42+2", 44},
		{"42+2-5", 39},
		{"42-2+5", 45},
		{"42+2*5", 52},
		{"42+2/5", 42},
		{"42*2-1", 83},
		{"42/2+5", 26},
	}
	for _, test := range cases {
		got := eval(test.input)
		if got != test.want {
			t.Fatalf("Expected %d, got %d", test.want, got)
		}
	}
}

func Test_Hints(t *testing.T) {

	g, err := New("84/2+0")
	if err != nil {
		t.Fatalf("Error when creating game")
	}

	cases := []struct {
		input string
		want  string
	}{
		{"", ""},
		{" ", ""},
		{"123456", ""},
		{"8*55/5", ""},

		{"111-69", "XXXXXX"},
		{"84/2+0", "CCCCCC"},

		{"042+00", "XCTTXC"},
		{"20*2+2", "XTXCCX"},
	}
	for _, test := range cases {
		got, _ := g.Hints(test.input)
		if got != test.want {
			t.Fatalf("Expected %s, got %s", test.want, got)
		}
	}
}

func sliceStringEquals(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func sliceIntEquals(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
