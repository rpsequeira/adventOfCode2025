package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)

	// read file as integers
	file := []int{}
	for reader.Scan() {
		var number int
		_, err := fmt.Sscanf(reader.Text(), "%d", &number)
		if err != nil {
			break
		}

		file = append(file, number)
	}

	// part 1
	// previous := -1
	// counter := 0
	// for i := range file {
	// 	if previous != -1 && previous < file[i] {
	// 		counter++
	// 	}

	// 	previous = file[i]
	// }
	// fmt.Println(counter)

	// part 2
	// previous = -1
	// counter = 0
	// for i := 0; i < len(file)-2; i++ {
	// 	sum := file[i] + file[i+1] + file[i+2]
	// 	if previous != -1 && previous < sum {
	// 		counter++
	// 	}

	// 	previous = sum
	// }
	// fmt.Println(counter)

}
