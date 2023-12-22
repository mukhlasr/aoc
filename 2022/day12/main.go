package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)

	var hill [][]byte
	var start, end position
	for y := 0; scn.Scan(); y++ {
		text := scn.Text()
		var row []byte
		for x, b := range text {
			c := byte(b)
			switch c {
			case 'S':
				start = position{
					x: x,
					y: y,
				}
			case 'E':
				end = position{
					x: x,
					y: y,
				}
			}
			row = append(row, c)
		}
		hill = append(hill, row)
	}
	p := hillClimbing(hill, start, end)
	fmt.Println(len(p))
}
