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
		//fmt.Println((len(arrayOfIntegers)-1))
		sumOfSquares(arrayOfIntegers, (len(arrayOfIntegers) - 1), 0)
		//fmt.Println("CHECK")
		return 0
	}
	var i int
	//reads digits from std input separated by blanks and saves them in i
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return 0
	}
	fmt.Println(i)

	arrayOfIntegers = append(arrayOfIntegers, i)

	return processInteger(amountOfIntegersToRead-1, arrayOfIntegers)
}

func sumOfSquares(arr []int, n int, sum int) int {
	//fmt.Println("sumOfSquares")

	if n == -1 {
		//fmt.Println("TOT APPEND")
		tot = append(tot, sum)
		return -1
	}
	if (arr[n] <= 0) && n >= 0 {
		return sumOfSquares(arr, n-1, sum)
	} else {
		//fmt.Println(arr[n])
		sum += (arr[n] * arr[n])
	}

	return sumOfSquares(arr, n-1, sum)

}
