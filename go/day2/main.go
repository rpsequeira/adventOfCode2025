package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/GRbit/go-pcre"
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

	if len(file) != 1 {
		log.Fatal("expected single line input")
	}

	prods := strings.Split(file[0], ",")
	// part 1
	counter := 0

	for _, rng := range prods {
		bounds := strings.Split(rng, "-")
		left, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		iter := left
		for iter <= right {
			sIter := strconv.Itoa(iter)
			if len(sIter)%2 != 0 {
				iter++
				continue
			}
			halfLen := len(sIter) / 2
			firstHalf := sIter[0:halfLen]
			secondHalf := sIter[halfLen:]
			if firstHalf == secondHalf {
				counter += iter
			}
			iter++
		}
	}
	fmt.Println(counter)

	// part 2
	counter = 0
	r := pcre.MustCompile(`^(.+?)\1+$`, 0)

	for _, rng := range prods {
		bounds := strings.Split(rng, "-")
		left, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		iter := left
		for iter <= right {
			matcher := r.NewMatcherString(strconv.Itoa(iter), 0)
			if matcher.Matches {
				counter += iter
			}
			iter++
		}
	}
	fmt.Println(counter)

}
