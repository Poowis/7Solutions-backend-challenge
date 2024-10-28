package main

import (
	"testing"
)

func TestFunction(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "LLRR=",
			arg:  "LLRR=",
			want: "210122",
		}, {
			name: "==RLL",
			arg:  "==RLL",
			want: "000210",
		}, {
			name: "=LLRR",
			arg:  "=LLRR",
			want: "221012",
		}, {
			name: "RRL=R",
			arg:  "RRL=R",
			want: "012001",
		}, {
			name: "RRRLL=L",
			arg:  "RRRLL=L",
			want: "01232110",
		}, {
			name: "RRLL=L",
			arg:  "RRRLLL=L",
			want: "012432110",
		}, {
			name: "=RRLL=L",
			arg:  "=RRRLLLLL=L",
			want: "001265432110",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Decoder(test.arg)
			if got != test.want {
				t.Errorf("got = %v, want = %v", got, test.want)
			}
		})
	}
}
