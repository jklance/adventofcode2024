package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getDataFields() [][]string { 
	f, _ := os.Open("day01a_input.txt")
	defer f.Close()
	
	inData := csv.NewReader(f)
	inData.Comma = '\t'

	dataRecords, _ := inData.ReadAll()

	return dataRecords
}

func flipDataFields(data [][]string) [][]int {
	var inData [][]int 
	var rowA []int
	var rowB []int

	for _, row := range data {
		row = strings.Fields(row[0])

		a, _ := strconv.Atoi(row[0])
		b, _ := strconv.Atoi(row[1])

		rowA = append(rowA, a)
		rowB = append(rowB, b)
	}
	sort.Ints(rowA)
	sort.Ints(rowB)

	inData = append(inData, rowA)
	inData = append(inData, rowB)

	return inData
}

func calculateDistance(data[][] int) int {
	dist := 0

	for i := 0; i < len(data[0]); i++ {
		if data[0][i] > data[1][i] {
			dist = dist + (data[0][i] - data[1][i])
		} else {
			dist = dist + (data[1][i] - data[0][i])
		}
	}

	return dist
}

func main() {
	inData := getDataFields()
	dataFields := flipDataFields(inData)

	distance := calculateDistance(dataFields)

	fmt.Println(distance)

	//var dataSeq []string

	/*inDataSubs := strings.Split(string(inData), "\n")
	for _, row := range inDataSubs {
		fmt.Println(row)

		rowDat := strings.Split(row, "\t")
		dataSeq = append(dataSeq, rowDat)
	}

	fmt.Println(dataSeq)*/



}