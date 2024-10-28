package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(Decoder(strings.TrimSpace(input)))
}

func Decoder(s string) string {
	output := make([]int, len(s)+1)
	for i, c := range s {
		switch c {
		case 'L':
			output[i+1] = 0
			if output[i] <= 0 {
				output[i]++
				for j := i - 1; j >= 0 && s[j] != 'R'; j-- {
					if s[j] == '=' {
						output[j] = output[j+1]
					} else if s[j] == 'L' && output[j] <= output[j+1] {
						output[j]++
					} else {
						break
					}
				}
			}
		case 'R':
			output[i+1] = output[i] + 1
		case '=':
			output[i+1] = output[i]
		}
	}
	outputStr := ""
	for _, n := range output {
		outputStr += strconv.Itoa(n)
	}
	return outputStr
}
