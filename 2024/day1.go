package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not read file")
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	left := []int{}
	right := []int{}
	for reader.Scan() {
		str := reader.Text()
		splitted := strings.Split(str, "   ")
		leftv, _ := strconv.ParseInt(splitted[0], 10, 32)
		rightv, _ := strconv.ParseInt(splitted[1], 10, 32)
		left = append(left, int(leftv))
		right = append(right, int(rightv))
	}

	slices.Sort(left)
	slices.Sort(right)
	diff := 0.0
	simScore := 0
	for idx := range left {
		diff += math.Abs(float64(left[idx]) - float64(right[idx]))
	}
	for _, l := range left {
		for _, r := range right {
			if l == r {
				simScore += l
			}
		}
	}
	fmt.Println(diff)
	fmt.Println(simScore)
}
