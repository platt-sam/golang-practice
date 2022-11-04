package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	var min, max, amount, sum int
	var numbers []int

	fmt.Printf("\nWELCOME TO THE RANDOM NUMBER GENERATOR")
	fmt.Printf("\n> Please enter the smallest possible number as an integer: ")
	fmt.Scan(&min)
	fmt.Print("\n> Please enter the largest possible number as an integer: ")
	fmt.Scan(&max)
	if min > max {
		t := max
		max = min
		min = t
	}
	fmt.Printf("The minimum permitted number is %v and the maximum permitted number is %v", min, max)
	fmt.Print("\n> Please set how many numbers should be generated: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		panic("The amount of numbers generated must be an integer greater than 0")
	}

	for i := 0; i < amount; i++ {
		// generate a random number between min and max and append it to numbers
		rand.Seed(time.Now().UnixNano())
		numbers = append(numbers, (rand.Intn(max+1-min) + min))
	}

	sort.Ints(numbers)

	// create a file to write to
	f, err := os.Create("random-numbers.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, v := range numbers {
		sum += v

		_, err := f.WriteString(strconv.Itoa(v) + "\n")
		if err != nil {
			panic(err)
		}
	}

	f.Close()

	fmt.Printf("A total of %d numbers between %d and %d were written to the file \"%v\"\n", amount, min, max, f.Name())
	fmt.Printf("The largest number generated was %v, the smallest number generated was %v, and the average number generated was %v\n", numbers[len(numbers)-1], numbers[0], float32(sum/amount))
}
