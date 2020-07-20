package botsandus

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func calculateChecksum(input string) (checksum int) {
	for i, num := range input {
		numValue, err := strconv.Atoi(string(num))
		if err == nil {
			checksum += numValue * int(math.Pow10(i))
		}
	}
	checksum = checksum % 97
	return
}

func EncodeNumber(input string) (lcdString string, checkSummed string, err error) {
	if len(input) != 4 {
		err = errors.New("Incorrect number length. Must be 4.")
		return
	}
	checksum := calculateChecksum(input)
	checkSummed = fmt.Sprintf("%02v%v", checksum, input)
	lcdString = getLCDStringForNumber(checkSummed)
	return
}
