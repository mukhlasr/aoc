package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("part1:", nCharsBeforePacketMarkers(data))
	fmt.Println("part2:", nCharsBeforeMessageMarker(data))

}
func nCharsBeforePacketMarkers(data []byte) int {
	const n = 4
	return firstUniqueNCharsPos(n, data) + n
}

func nCharsBeforeMessageMarker(data []byte) int {
	const n = 14
	return firstUniqueNCharsPos(n, data) + n
}

func firstUniqueNCharsPos(n int, b []byte) int {
	for i := 0; i < len(b); i++ {
		if isUniqueChars(b[i : i+n]) {
			return i
		}
	}
	return -1
}

func isUniqueChars(b []byte) bool {
	for i, char := range b {
		for j := i + 1; j < len(b); j++ {
			if char == b[j] {
				return false
			}
		}
	}
	return true
}
