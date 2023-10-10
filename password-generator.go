package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Password:", password())
}

func password() string {
	lettersLowerCase := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	lettersUpperCase := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	passwordArr := []rune{}
	for i := 0; i < 16; i++ {
		temp, err := rand.Int(rand.Reader, big.NewInt(int64(len(lettersLowerCase))))
		if err != nil {
			return "Something went wrong"
		}
		passwordArr = append(passwordArr, lettersLowerCase[temp.Int64()])
	}
	temp, err := rand.Int(rand.Reader, big.NewInt(int64(len(lettersUpperCase))))
	if err != nil {
		return "Something went wrong"
	}
	passwordArr = append(passwordArr, lettersUpperCase[temp.Int64()])
	temp, err = rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
	if err != nil {
		return "Something went wrong"
	}
	passwordArr = append(passwordArr, numbers[temp.Int64()])

	password := ""
	for len(passwordArr) != 0 {
		temp, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordArr))))
		if err != nil {
			return "Something went wrong"
		}
		password = password + string(passwordArr[temp.Int64()])
		passwordArr = deleteElement(passwordArr, passwordArr[temp.Int64()])
		if len(passwordArr)%6 == 0 {
			password = password + "-"
		}
	}
	return password[:len(password)-1]
}
func deleteElement(slice []rune, character rune) []rune {
	var index int
	for i, _ := range slice {
		if slice[i] == character {
			index = i
			break
		}
	}
	return append(slice[:index], slice[index+1:]...)
}