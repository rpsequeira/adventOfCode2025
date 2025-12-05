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
	fresh := [][]int{}
	available := []int{}
	linebreak := false
	for reader.Scan() {
		var line string
		_, err := fmt.Sscanf(reader.Text(), "%s", &line)
		if err != nil {
			break
		}
		splited := strings.Split(line, "-")

		if line == "@" {
			linebreak = true
			continue
		}

		if linebreak {

			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			available = append(available, i)
		} else {
			left, err := strconv.Atoi(splited[0])
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.Atoi(splited[1])
			if err != nil {
				log.Fatal(err)
			}
			fresh = append(fresh, []int{left, right})
		}
	}

	// part 1
	counter := 0
	for _, ing := range available {
		for _, fre := range fresh {
			if ing >= fre[0] && ing <= fre[1] {
				counter++
				break
			}
		}
	}
	fmt.Println(counter)

	// part 2
	// allFresh := map[int]struct{}{}
	// for _, fre := range fresh {
	// 	for i := fre[0]; i <= fre[1]; i++ {
	// 		allFresh[i] = struct{}{}
	// 	}
	// }
	// fmt.Println(len(allFresh))
	optFresh := [][]int{fresh[0]}
	for _, fre := range fresh[1:] {
		newFresh := [][]int{fre}
		for _, ofre := range optFresh {
			merged := false
			for i, nfre := range newFresh {
				if ofre[0] <= nfre[1] && ofre[0] >= nfre[0] && ofre[1] >= nfre[1] {
					newFresh[i][1] = ofre[1]
					merged = true
					break
				}
				if ofre[1] >= nfre[0] && ofre[1] <= nfre[1] && ofre[0] <= nfre[0] {
					newFresh[i][0] = ofre[0]
					merged = true
					break
				}
				if ofre[0] <= nfre[0] && ofre[1] >= nfre[1] {
					newFresh[i][0] = ofre[0]
					newFresh[i][1] = ofre[1]
					merged = true
					break
				}
				if ofre[0] >= nfre[0] && ofre[1] <= nfre[1] {
					merged = true
					break
				}
			}
			if !merged {
				newFresh = append(newFresh, ofre)
			}
		}
		optFresh = newFresh
	}
	counter = 0
	for _, ing := range optFresh {
		counter += ing[1] - ing[0] + 1
	}
	fmt.Println(counter)

}
