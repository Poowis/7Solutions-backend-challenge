package utils

import (
	"maps"
	"testing"
)

var baseString = "Fatback\tt-bone t-bone, pastrami  ..   t-bone.pork,meatloaf jowl\nenim.  Bresaola t-bone  t"
var wantMap = map[string]int{
	"bresaola": 1,
	"enim":     1,
	"fatback":  1,
	"jowl":     1,
	"meatloaf": 1,
	"pastrami": 1,
	"pork":     1,
	"t-bone":   4,
	"t":        1,
}
var tests = []struct {
	name string
	arg  string
	want map[string]int
}{
	{
		name: "base",
		arg:  baseString,
		want: wantMap,
	}, {
		name: "base+.",
		arg:  baseString + ".",
		want: wantMap,
	}, {
		name: "base+.,",
		arg:  baseString + ".,",
		want: wantMap,
	}, {
		name: "base+., ",
		arg:  baseString + "., ",
		want: wantMap,
	},
}

func TestFunction(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Summary(test.arg)
			if !maps.Equal(got, test.want) {
				t.Errorf("got = %v, want = %v", got, test.want)
			}
		})
	}
}

func BenchmarkFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Summary(baseString)
	}
}

func TestSplit(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Split(test.arg)
			if !maps.Equal(got, test.want) {
				t.Errorf("got = %v, want = %v", got, test.want)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split(baseString)
	}
}
