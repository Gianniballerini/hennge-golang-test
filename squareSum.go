package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var TotalTestCases int
var tot []int

func main() {
	// ammount of test cases
	var outerLoop int

	scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("How many test cases?")
	// read the first line
	scanner.Scan()

	// convert readed text to int
	outerLoop, err := strconv.Atoi(scanner.Text())
	// if conversion was unsuccessfull
	if err != nil {
		fmt.Println("outerLoop error")
	}

	// if reading error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	TotalTestCases = outerLoop
	processTestCase(outerLoop)
}

func printAnswer(n int, c int) int {
	// base case for recursion
	if n == 0 {
		return 0
	}
	fmt.Println(tot[c])
	return printAnswer(n-1, c+1)
}

func processTestCase(testCaseNumber int) int {
	//fmt.Println("processTestCase")
	// base case for recursion
	if testCaseNumber == 0 {
		printAnswer(TotalTestCases, 0)
		return 0
	}

	var numberOfIntegersToRead int
	var integerArray []int
	scanner := bufio.NewScanner(os.Stdin)
	// read first entry of test case (how many integers will have)
	//fmt.Println("How many integers in the test case?")
	scanner.Scan()
	// convert it to integer
	numberOfIntegersToRead, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("numberOfIntegersToRead error")
	}

	// this call will process all the integers for this test case
	processInteger(numberOfIntegersToRead, integerArray)

	// call this method again for the next test case
	return processTestCase(testCaseNumber - 1)
}

func processInteger(amountOfIntegersToRead int, arrayOfIntegers []int) int {
	//fmt.Println("processInteger")
	// meaning i have read all the numbers i needed
	if amountOfIntegersToRead == 0 {
		sumOfSquares(arrayOfIntegers, (len(arrayOfIntegers) - 1), 0)
		return 0
	}
	var i int
	//reads digit from std input separated by blanks and saves them in i
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return 0
	}
	//fmt.Println(i)

	arrayOfIntegers = append(arrayOfIntegers, i)

	// call this method again for the next case
	return processInteger(amountOfIntegersToRead-1, arrayOfIntegers)
}

func sumOfSquares(arr []int, n int, sum int) int {
	//fmt.Println("sumOfSquares")

	// base case, position -1 all number processed
	if n == -1 {
		//fmt.Println("TOT APPEND")
		// append total to final answer array
		tot = append(tot, sum)
		return -1
	}
	// if the number is negative do not sum and go to next number
	if (arr[n] <= 0) && n >= 0 {
		return sumOfSquares(arr, n-1, sum)
	} else {
		// sum of square of the number
		sum += (arr[n] * arr[n])
	}

	// recursion next number
	return sumOfSquares(arr, n-1, sum)

}
