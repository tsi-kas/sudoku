package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	if !VerifyParams(arguments) {
		fmt.Println("Error")
		return
	}
	table := [9][9]int{}
	table = FillTable(table, arguments)

	if !(SolvesSudoku(table)) {
		fmt.Println("Error")
	}
}

func VerifyParams(args []string) bool { // this makes sure there are 9 string arguments in the terminal
	if len(args) != 9 {
		return false
	}
	for _, row := range args {
		if len(row) != 9 {
			return false
		}
		for _, value := range row {
			if value != '.' && (value < '1' || value > '9') {
				return false
			}
		}
	}
	return true
}

func FillTable(table [9][9]int, arguments []string) [9][9]int { // this fills an empty 2D array with the rune values from the arguments in the terminal
	for rowindex := range arguments {
		for columnindex, value := range arguments[rowindex] {
			if value == 46 {
				value = value + 2
			}
			table[rowindex][columnindex] = int(value - '0')
		}
	}
	return table
}

func FindEmptySlots(table [9][9]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if table[row][column] == 0 { // finds the empty slots
				return true
			}
		}
	}
	return false
}

func SolvesSudoku(table [9][9]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if table[row][column] == 0 { // finds the empty slots
				for value := 1; value <= 9; value++ {
					if IsValid(table, row, column, value) { // if the value does not already exist in the same subtable, row and column, then IsValid is true, and the value is placed into the table
						table[row][column] = value
						if NoEmptiesLeft(table) {
							for row := 0; row < 9; row++ {
								for column := 0; column < 9; column++ {
									if column < 8 {
										fmt.Print(table[row][column])
										z01.PrintRune(' ')
									}
									if column == 8 {
										fmt.Print(table[row][column])
									}
								}
								z01.PrintRune('\n')
							}
						}
						if SolvesSudoku(table) { // then the table with the new value is sent back into the SolvesSudoku function. This achieves the backtracking.
							return true
						}
					}
				}
				return false // sends back to the for loop above IsValid.
			}
		}
	}
	return true
}

func NoEmptiesLeft(table [9][9]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if table[row][column] == 0 {
				return false
			}
		}
	}
	return true
}

func IsValid(table [9][9]int, row, column, value int) bool {
	for index := 0; index < 9; index++ { // this is checking if the value already exists in the same row and column
		if table[row][index] == value || table[index][column] == value {
			return false
		}
	}
	subtablerow := (row / 3) * 3 // this is checking if the value already exists in the subtable
	subtablecolumn := (column / 3) * 3
	for a := subtablerow; a < subtablerow+3; a++ {
		for b := subtablecolumn; b < subtablecolumn+3; b++ {
			if table[a][b] == value {
				return false
			}
		}
	}
	return true

}
