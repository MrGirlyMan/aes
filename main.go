package main

import (
	"fmt"
//	"reflect"
	"errors"
)

func PrintHexStateMatrix(m [][]byte) {
	printStateMatrix(m, 0)
}

func PrintASCIIStateMatrix(m [][]byte) {
	printStateMatrix(m, 1)
}

func printStateMatrix(m [][]byte, mode int) {
	var format string

	switch mode {
	case 0:  // Hex
		format = "%3X"
	case 1:  // ASCII
		format = "%3c"
	default:
		format = "%3X"
	}

	for _, row := range m {
		for _, value := range row {
			fmt.Printf(format, value)
		}
		fmt.Println()
	}

	fmt.Println()
}

func isMatrix() error {
	err := errors.New("Value is not matrix")
	return err
}

func createStateMatrix(size int) [][]byte {
	stateMatrix := make([][]byte, size)

	for i := range stateMatrix {
		stateMatrix[i] = make([]byte, size)
	}

	return stateMatrix
}

func initializeStateMatrix(m [][]byte, secret string) {
	for i, row := range m {
		for j, _ := range row {
			m[j][i] = secret[i * len(row) + j]
		} 
	}
}

func main() {
	m := createStateMatrix(4)

	initializeStateMatrix(m, "Secret, secret!!")

	PrintHexStateMatrix(m)

	PrintASCIIStateMatrix(m)
}

