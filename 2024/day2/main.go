package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkSafe(l []int, resChan chan<- bool) []int {
	banList := make([]int, 0)
	var DiffF func(int) bool
	if l[0] < l[1] {
		DiffF = func(d int) bool {return d <= 3 && d >= 1}
	} else if l[0] > l[1] {
		DiffF = func(d int) bool {return -d <= 3 && -d >= 1}
	} else {
		if resChan != nil {
			resChan <- false
		}
		banList = append(banList, 0, 1, 2)
		return banList
	}
	for i := 0; i < len(l)-1; i++ {
		diff := l[i+1] - l[i]
		if DiffF(diff) {
			continue
		} else {
			if resChan != nil {
				resChan <- false
			}
			banList = append(banList, 0, i, i+1)
			return banList
		}
	}
	if resChan != nil {
		resChan <- true
	}
	return nil
}

func Map(l []string) []int {
	l_i := make([]int, len(l))
	for i, c := range l {
		parsed, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		l_i[i] = int(parsed)
	}
	return l_i
}

func removeElement(l []int, i int) []int {
	newL := make([]int, 0)
	newL = append(newL, l[:i]...)
	return append(newL, l[i+1:]...)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := bufio.NewScanner(f)

	count := 0
	countDamp := 0

	for b.Scan() {
		l := strings.Split(b.Text(), " ")
		toCheck := Map(l)

		banList := checkSafe(toCheck, nil)
		if banList != nil {
			for _, i := range banList {
				newBanList := checkSafe(removeElement(toCheck, i), nil)
				if newBanList == nil {
					countDamp++
					break
				}
			}
		} else {
			count++
		}
	}

	fmt.Printf("Part one: %d\n", count)
	fmt.Printf("Part two: %d\n", countDamp+count)

}
