package main

import "testing"

func TestNormalize(t *testing.T) {
	testCases := []struct {
		input string
		want string
	}{
		{"+46(769)123-421", "0769123421"},
		{"+46769-123 422", "0769123422"},
		{"+46-769-123-423", "0769123423"},
		{"+46(769)123 424", "0769123424"},
		{"+46 768 987 656", "0768987656"},
		{"0768-987-657", "0768987657"},
		{"+46-768(987 658)", "0768987658"},
		{"+46768987659", "0768987659"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t * testing.T) {
			actual := normalize(tc.input)
			if actual != tc.want {
				t.Errorf("Got %s, wanted %s", actual, tc.want)
			}
		})
	}
}