//Package keybaord - reading data entered by user on keyboard
package readKeyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)
//GetFloat - reading float64 number form keyboard
//Returns read number and an error
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return number, nil
}
