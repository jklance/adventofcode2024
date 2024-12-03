package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func integerizeFields(data [][]string) [][]int {
	var inData [][]int 

	for _, row := range data {
		row = strings.Fields(row[0])

		var newRow []int
		for _, col := range row {
			tmp, _ := strconv.Atoi(col)
			newRow = append(newRow, tmp)
		}

		inData = append(inData, newRow)
	}

	return inData
}

func getDataFields() [][]int { 
	f, _ := os.Open("day02_input.txt")
	defer f.Close()
	
	inData := csv.NewReader(f)
	inData.Comma = '\t'

	dataRecords, _ := inData.ReadAll()

	dataFields := integerizeFields(dataRecords)

	return dataFields
}

func isSafeChange(prior int, current int, direction int) bool {
	if prior == current {
		return false
	}

	if direction > 0 {
		if current - prior > 3 || current < prior {
			return false
		}
	} else {
		if prior - current > 3 || prior < current {
			return false
		}
	}

	return true
}
func testReport(report []int) bool {
	safeReport := true
	direction := 1
	
	for i, level := range report {
		// Establish which direction the levels are going; -1 for down and 1 for up
		// (non-change will be covered by safe gap test)
		if i == 1 && (report[i] - report[i - 1]) < 0 {
			direction = -1
		}

		if i > 0 {
			if !isSafeChange(report[i - 1], level, direction) {
				safeReport = false
				//fmt.Println(i, level, safeReport)
				break
			}

		}
	}

	return safeReport
}

func testDataSet(dataSet [][]int) int {
	safeReportCount := 0

	for _, report := range dataSet {
		if testReport(report) {
			safeReportCount++
		}
	}

	return safeReportCount
}

func main() {
	inData := getDataFields()

	fmt.Println(inData)

	safeReportCount := testDataSet(inData)

	fmt.Println(safeReportCount)
}