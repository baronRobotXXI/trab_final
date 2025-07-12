package utils

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

func ReadLines (path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		lines = append(lines,num)
	}
	return lines
	
}
