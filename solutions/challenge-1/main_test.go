package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestFunction(t *testing.T) {
	jsonFile, err := os.ReadFile("../../files/hard.json")
	if err != nil {
		t.Errorf("hard.json file is missing")
	}
	var data [][]int
	json.Unmarshal(jsonFile, &data)

	tests := []struct {
		name string
		arg  [][]int
		want int
	}{
		{
			name: "len = 0",
			arg:  [][]int{},
			want: 0,
		}, {
			name: "len = 1",
			arg:  [][]int{{1}},
			want: 1,
		}, {
			name: "len = 2",
			arg:  [][]int{{1}, {2, 3}},
			want: 4,
		}, {
			name: "easy data",
			arg:  [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}},
			want: 237,
		},
		{
			name: "hard data",
			arg:  data,
			want: 7273,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MaxSum(test.arg)
			if got != test.want {
				t.Errorf("got = %v, want = %v", got, test.want)
			}
		})
	}
}

func BenchmarkFunction(b *testing.B) {
	jsonFile, err := os.ReadFile("../../files/hard.json")
	if err != nil {
		b.Errorf("hard.json file is missing")
	}
	var data [][]int
	json.Unmarshal(jsonFile, &data)

	for i := 0; i < b.N; i++ {
		MaxSum(data)
	}
}
