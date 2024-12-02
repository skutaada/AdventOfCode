package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkSafe(l *[]int, res chan<- bool, badIndex chan<- int, diffF func(d int) bool) {
	for i := range *l {
		if i == len(*l)-1 {
			res <- true
			return
		}
		diff := (*l)[i+1] - (*l)[i]
		if diffF(diff) {
			continue
		} else {
			res <- false
			if badIndex != nil {
				badIndex <- i + 1
			}
			return
		}
	}
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

		resIncChan := make(chan bool, 2)
		resDecChan := make(chan bool, 2)
		indexIncChan := make(chan int)
		indexDecChan := make(chan int)

		go checkSafe(&toCheck, resIncChan, indexIncChan, func(d int) bool { return d <= 3 && d >= 1 })
		go checkSafe(&toCheck, resDecChan, indexDecChan, func(d int) bool { return -d <= 3 && -d >= 1 })
		resInc := <-resIncChan
		resDec := <-resDecChan
		if resInc || resDec {
			count++
			countDamp++
		} else {
			index1 := <-indexIncChan
			fmt.Println(index1)
			index2 := <-indexDecChan
			l_new_inc := make([]int, 0)
			l_new_inc = append(l_new_inc, toCheck[:index1]...)
			l_new_inc = append(l_new_inc, toCheck[index1+1:]...)
			l_new_dec := make([]int, 0)
			l_new_dec = append(l_new_dec, toCheck[:index2]...)
			l_new_dec = append(l_new_dec, toCheck[index2+1:]...)
			fmt.Println(l_new_inc)
			fmt.Println(l_new_dec)
			go checkSafe(&l_new_inc, resIncChan, nil, func(d int) bool { return d <= 3 && d >= 1 })
			go checkSafe(&l_new_dec, resIncChan, nil, func(d int) bool { return d <= 3 && d >= 1 })
			go checkSafe(&l_new_inc, resDecChan, nil, func(d int) bool { return -d <= 3 && -d >= 1 })
			go checkSafe(&l_new_dec, resDecChan, nil, func(d int) bool { return -d <= 3 && -d >= 1 })
			for range 2 {
				resInc = <-resIncChan
				fmt.Println(resInc)
				resDec = <-resDecChan
				if resInc || resDec {
					countDamp++
					break
				}
			}

		}
	}

	fmt.Printf("Part one: %d\n", count)
	fmt.Printf("Part two: %d\n", countDamp)

}
