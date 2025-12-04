package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	file := [][]string{}
	for reader.Scan() {
		var line string
		_, err := fmt.Sscanf(reader.Text(), "%s", &line)
		if err != nil {
			break
		}

		file = append(file, strings.Split(line, ""))
	}

	// part 1
	counter := 0
	for i := range file {
		for j := range file[i] {
			innerCounter := 0
			if file[i][j] == "." {
				continue
			}
			if i > 0 {
				if j > 0 {
					if file[i-1][j-1] == "@" {
						innerCounter++
					}
				}
				if file[i-1][j] == "@" {
					innerCounter++
				}
				if j < len(file[i])-1 {
					if file[i-1][j+1] == "@" {
						innerCounter++
					}
				}
			}
			if j > 0 {
				if file[i][j-1] == "@" {
					innerCounter++
				}
			}
			if j < len(file[i])-1 {
				if file[i][j+1] == "@" {
					innerCounter++
				}
			}

			if i < len(file)-1 {
				if j > 0 {
					if file[i+1][j-1] == "@" {
						innerCounter++
					}
				}
				if file[i+1][j] == "@" {
					innerCounter++
				}
				if j < len(file[i])-1 {
					if file[i+1][j+1] == "@" {
						innerCounter++
					}
				}
			}
			if innerCounter < 4 {
				counter++
			}
		}
	}
	fmt.Println(counter)

	// part 2
	counter = 0
	snapshot := 0
	for {
		for i := range file {
			for j := range file[i] {
				innerCounter := 0
				if file[i][j] != "@" {
					continue
				}
				if i > 0 {
					if j > 0 {
						if file[i-1][j-1] == "@" {
							innerCounter++
						}
					}
					if file[i-1][j] == "@" {
						innerCounter++
					}
					if j < len(file[i])-1 {
						if file[i-1][j+1] == "@" {
							innerCounter++
						}
					}
				}
				if j > 0 {
					if file[i][j-1] == "@" {
						innerCounter++
					}
				}
				if j < len(file[i])-1 {
					if file[i][j+1] == "@" {
						innerCounter++
					}
				}

				if i < len(file)-1 {
					if j > 0 {
						if file[i+1][j-1] == "@" {
							innerCounter++
						}
					}
					if file[i+1][j] == "@" {
						innerCounter++
					}
					if j < len(file[i])-1 {
						if file[i+1][j+1] == "@" {
							innerCounter++
						}
					}
				}
				if innerCounter < 4 {
					file[i][j] = "x"
					counter++
				}
			}
		}
		if snapshot == counter {
			break
		}
		snapshot = counter
	}
	fmt.Println(counter)
}
