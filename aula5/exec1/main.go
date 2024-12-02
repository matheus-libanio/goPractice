package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main()  {

	err:= verifySalary()
	if err != nil {
		logrus.Infoln(err)
	}
}

func verifySalary() error {
	var salary int

	fmt.Println("Enter annual salary:")

	_, err :=fmt.Scan(&salary)
	if err != nil {
		return errors.New("Error: the salary entered does not match our monetary method")
	}
	if salary < 150000{
		return errors.New("Error: the salary entered does not reach the taxable minimum")
	}
	logrus.Infoln("Must pay taxes")
	return nil
}