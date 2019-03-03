package maxRotate

import (
	"sort"
	"strconv"
	"strings"
)

func rotateLeft(startIndex int, number []string) []string {
	numberLength := len(number)
	newNumbers := make([]string, numberLength)
	copy(newNumbers, number)

	for i := startIndex; i < (numberLength - 1); i++ {
		newNumbers[i], newNumbers[(i+1)] = newNumbers[(i+1)], newNumbers[i]
	}

	return newNumbers
}

func MaxRot(n int64) int64 {
	nString := strconv.FormatInt(n, 10)
	nLength := len(nString)
	nRotating := strings.Split(nString, "")
	rotations := make([]string, 0)
	rotations = append(rotations, nString)

	for i := 0; i < nLength; i++ {
		nRotating = rotateLeft(i, nRotating)
		rotations = append(rotations, strings.Join(nRotating, ""))
	}

	sort.Strings(rotations)
	max, _ := strconv.ParseInt(rotations[nLength], 10, 64)

	return max
}
