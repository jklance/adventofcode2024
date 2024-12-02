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

func calculateSimilarity(data[][] int) int {
	sim := 0

	for i := 0; i < len(data[0]); i++ {
		count := 0
		for j := 0; j < len(data[1]); j++ {
			if data[0][i] == data[1][j] {
				count++
			}
		}
		sim = sim + (data[0][i] * count)

	}

	return sim
}

func main() {
	inData := getDataFields()
	dataFields := flipDataFields(inData)

	similarity := calculateSimilarity(dataFields)

	fmt.Println(similarity)

}