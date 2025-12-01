package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)

	// read file as integers
	file := []string{}
	for reader.Scan() {
		var line string
		_, err := fmt.Sscanf(reader.Text(), "%s", &line)
		if err != nil {
			break
		}

		file = append(file, line)
	}

	// part 1
	dial := 50
	counter := 0
	for i := range file {
		direction := file[i][0:1]
		number, err := strconv.Atoi(file[i][1:])
		// fmt.Print(file[i])
		if err != nil {
			log.Fatal(err)
		}
		if direction == "R" {
			dial += number
			for dial >= 100 {
				dial -= 100
			}
		} else if direction == "L" {
			dial -= number
			for dial < 0 {
				dial += 100
			}
		} else {
			log.Fatal("unknown direction")
		}
		// fmt.Printf(" - %d\n", dial)
		if dial == 0 {
			counter++
		}
	}
	fmt.Println(counter)

	// part 2
	dial = 50
	counter = 0
	for i := range file {
		direction := file[i][0:1]
		number, err := strconv.Atoi(file[i][1:])
		// fmt.Print(file[i])
		if err != nil {
			log.Fatal(err)
		}

		if direction == "R" {
			dial += number
			for dial > 100 {
				dial -= 100
				counter++
				// fmt.Print(" * ")
			}
		} else if direction == "L" {
			if dial == 0 {
				counter--
				// fmt.Print(" % ")
			}
			dial -= number
			for dial < 0 {
				dial += 100
				counter++
				// fmt.Print(" * ")
			}
		} else {
			log.Fatal("unknown direction")
		}

		if dial == 100 {
			dial = 0
		}
		if dial == 0 {
			counter++
			// fmt.Print(" * ")
		}
		// fmt.Printf(" - %d\n", dial)
	}
	fmt.Println(counter)

}
