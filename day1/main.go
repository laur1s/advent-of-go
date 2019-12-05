package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	check(err)
	ints, _ := ReadInts(strings.NewReader(string(dat)))
	check(err)
	sum := 0
	for i := 0; i < len(ints); i++ {
		sum += getMass(ints[i])
	}
	fmt.Println(sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getMass(mass int) int {
	fuel := mass/3 - 2
	return fuel
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
