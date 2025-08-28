package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	iterationCount, _ := strconv.Atoi(scanner.Text())

	results := make([]int, 0, iterationCount)
	for i := 0; i < iterationCount; i++ {
		scanner.Scan()
		inputList := strings.Fields(scanner.Text())
		x1, _ := strconv.Atoi(inputList[0])
		y1, _ := strconv.Atoi(inputList[1])
		r1, _ := strconv.Atoi(inputList[2])
		x2, _ := strconv.Atoi(inputList[3])
		y2, _ := strconv.Atoi(inputList[4])
		r2, _ := strconv.Atoi(inputList[5])

		distance := math.Hypot(float64(x1-x2), float64(y1-y2))
		rSum := float64(r1 + r2)
		rDiff := math.Abs(float64(r1 - r2))

		if rDiff < distance && distance < rSum {
			results = append(results, 2)
		} else if distance == rSum || (distance == rDiff && distance != 0) {
			results = append(results, 1)
		} else if distance == 0 && r1 == r2 {
			results = append(results, -1)
		} else {
			results = append(results, 0)
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}