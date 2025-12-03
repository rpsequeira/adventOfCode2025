package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)

	// read file
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
	counter := 0
	for i := range file {
		max := 0
		btts := strings.Split(file[i], "")
		// fmt.Printf("b - %v\n", btts)
		for j := range btts {
			k := j + 1
			for k < len(btts) {
				x, err := strconv.Atoi((btts[j] + btts[k]))
				if err != nil {
					log.Fatal(err)
				}
				// fmt.Printf("1 - %d\n", x)
				if x > max {
					max = x
				}
				k++
			}
		}
		// fmt.Printf("max - %d\n", max)
		counter += max
	}
	fmt.Println(counter)

	// part 2
	counter = 0
	for i := range file {
		max := []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
		btts := strings.Split(file[i], "")
		// fmt.Printf("b - %v\n", btts)
		size := 12
		idxK := 0

		for j, btt := range btts {
			x, err := strconv.Atoi(btt)
			if err != nil {
				log.Fatal(err)
			}
			if len(btts)-j < size {
				idxK = size - (len(btts) - j)
			}
			clear := false
			for k, m := range max {
				if k < idxK {
					continue
				}
				if clear {
					max[k] = "0"
					continue
				}
				// fmt.Printf("k - %d, btt - %s, max - %v, idx - %d, len - %d\n", k, btt, max, idxK, len(btts)-j)
				mInt, err := strconv.Atoi(m)
				if err != nil {
					log.Fatal(err)
				}
				if x > mInt {
					max[k] = btt
					clear = true
				}
			}
		}
		numStr := strings.Join(max, "")
		fmt.Println(numStr)
		maxInt, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("max - %d\n", max)
		counter += maxInt
	}
	fmt.Println(counter)
}
