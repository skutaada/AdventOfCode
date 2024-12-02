package day1

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
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	var left []int
	var right []int
	for reader.Scan() {
		str := reader.Text()
		splitted := strings.Split(str, "   ")
		leftv, err := strconv.ParseInt(splitted[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		rightv, err := strconv.ParseInt(splitted[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, int(leftv))
		right = append(right, int(rightv))
	}

	slices.Sort(left)
	slices.Sort(right)
	diff := 0
	simScore := 0
	for idx := range left {
		diff += int(math.Abs(float64(left[idx]) - float64(right[idx])))
	}
	for _, l := range left {
		for _, r := range right {
			if l == r {
				simScore += l
			}
		}
	}
	fmt.Printf("Part one: %d\n", diff)
	fmt.Printf("Part two: %d\n", simScore)
}
