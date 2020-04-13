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
	reader := bufio.NewReader(os.Stdin)
	// reader := fileReader("/Users/bmancini/code/go/src/github.com/bmancini55/codeforces/p1328a/input.txt")

	t, err := readInt(reader)
	if err != nil {
		log.Fatal(err)
	}

	tuples, err := readInts(reader, t)
	if err != nil {
		log.Fatal(err)
	}

	for _, tuple := range tuples {
		a := tuple[0]
		b := tuple[1]

		// fmt.Printf("a: %d, b %d, mod: %d, eq: %v\n", a, b, a%b, a%b == 0)
		if a%b == 0 {
			fmt.Printf("%d\n", 0)
		} else {
			r := b - (a % b)
			fmt.Printf("%d\n", r)
		}
	}

}

func fileReader(path string) *bufio.Reader {
	var (
		f   *os.File
		r   *bufio.Reader
		err error
	)

	f, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	r = bufio.NewReader(f)
	return r
}

func readInt(reader *bufio.Reader) (int, error) {
	var (
		result int
		err    error
	)

	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(line)
	result, err = strconv.Atoi(fields[0])

	if err != nil {
		return 0, err
	}

	return result, nil
}

func readInts(reader *bufio.Reader, n int) ([][]int, error) {
	results := [][]int{}

	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		nums, err := parseInts(line)
		if err != nil {
			return nil, err
		}

		results = append(results, nums)
	}
	return results, nil
}

func parseInts(s string) ([]int, error) {
	var (
		fields = strings.Fields(s)
		ints   = make([]int, len(fields))
		err    error
	)
	for i, f := range fields {
		ints[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}
