package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

const (
	DO = "do()"
	DONT = "don't()"
)

func handleMul(exp []byte) int {
	toStr := string(exp)
	var left, right int
	fmt.Sscanf(toStr, "mul(%d,%d)", &left, &right)
	return left*right
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	re2 := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\))`)
	finds := re.FindAll(bytes, -1)
	finds2 := re2.FindAll(bytes, -1)
	sum := 0
	sumEn := 0

	for _, exp := range finds {
		sum += handleMul(exp)
	}

	enabled := true
	for _, exp := range finds2 {
		if string(exp) == DO {
			enabled = true
		} else if string(exp) == DONT {
			enabled = false
		} else if enabled {
			sumEn += handleMul(exp)
		}
	}



	fmt.Printf("Part one: %d\n", sum)
	fmt.Printf("Part two: %d\n", sumEn)
}