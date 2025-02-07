// in Go we don't use exceptions
// if != err
package main

import (
	"errors"
	"fmt"
)

// err.New()
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}

	return a / b, nil
}

// fmt.Errorf() --> allows formatting error messages
func underAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %v provided is below threshold", age)
	}
	return nil
}

// wrap existing error with %w
func readfile() error {
	return errors.New("file not found")
}
func openfile() error {
	err := readfile()

	if err != nil {
		return fmt.Errorf("Error: %w occurred", err)
	}
	return nil
}

//errors.Is() --> compares errors to check for a specific type
var ErrFileNotFound = errors.New("File not found")

func readfile2() error{
	return ErrFileNotFound
}

func main() {
	res, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Printf("Result : %v \n", res)

	err1 := underAge(16)

	if err1 != nil {
		fmt.Println("Error :", err1)
	} else {
		fmt.Println("Registration successful")
	}

	err2 := openfile()

	if err2 != nil {
		fmt.Println("Error2 :", err2)
	}

	//unwrapping error
	err3 := openfile()

	if err3 != nil {
		fmt.Println(err3)

		//unwrapping
		unwrapped := errors.Unwrap(err3)
		fmt.Println("Unwrapped error: ", unwrapped)
	}

	err4 := readfile2()

	//using errors.Is()
	if errors.Is(err4, ErrFileNotFound) {
		fmt.Println("Error is a file not found error")
	} else {
		fmt.Println("Some other error occurred")
	}
}
