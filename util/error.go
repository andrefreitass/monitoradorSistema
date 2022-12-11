package util

import "fmt"

func ValidaError(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
}
