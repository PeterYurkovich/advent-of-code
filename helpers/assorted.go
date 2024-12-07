package helpers

import (
	"os"
)

func GetFileString(fileName string) string {
	byteFile, err := os.ReadFile(fileName)
	AssertError(err)
	return string(byteFile)
}

func Remove(s []int, i int) []int {
	newArray := make([]int, len(s))
	copy(newArray, s)
	return append(newArray[:i], newArray[i+1:]...)
}
