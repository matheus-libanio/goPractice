package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	inputError error =  errors.New("unexpected input value")
	validationError error = errors.New("lower Value")
)

func main()  {

	err:= verifySalary()
	targetError := inputError

	if errors.Is(err, targetError) {
		logrus.Infoln(errors.Is(err, targetError), err)
	}
}

func verifySalary() error {
	var salary int

	fmt.Println("Enter annual salary:")

	_, err :=fmt.Scan(&salary)
	if err != nil {
		return fmt.Errorf("Error: the salary entered does not match our monetary method (%w)",inputError)
	}
	if salary < 150000{
		return fmt.Errorf("Error: the salary entered does not reach the taxable minimum (%w)",validationError)
	}
	logrus.Infoln("Must pay taxes")
	return nil
}