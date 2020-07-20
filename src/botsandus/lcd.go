package botsandus

import "strconv"
import "fmt"

var lcdNumbers = []byte{
	// Golang only supports 0bxxxxxxx literals from 1.13
	0x00, //0b00000000, //0
	0x42, //0b01000010, //1
	0xb6, //0b10110110, //2
	0xd6, //0b11010110, //3
	0xc3, //0b11000011, //4
	0xd5, //0b11010101, //5
	0xf5, //0b11110101, //6
	0x46, //0b01000110, //7
	0xf7, //0b11110111, //8
	0xd7, //0b11010111 //9
}

func getLCDStringForNumber(input string) (buf string) {
	for _, num := range input {
		numValue, err := strconv.Atoi(string(num))
		if err == nil {
			buf += fmt.Sprintf("%08b", lcdNumbers[numValue])
		}
	}
	return
}
