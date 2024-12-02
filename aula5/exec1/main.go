package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {

	err := verifySalary()
	if err != nil {
		logrus.Infoln(err)
	}
}

type erroP struct {
	msg string
}

func (e *erroP) Error() string {
	return fmt.Sprintf("An error occurred: %s", e.msg)
}

func verifySalary() error {
	var salary int

	msgErro := &erroP{"Input invalid"}

	fmt.Println("Enter annual salary:")

	_, err := fmt.Scan(&salary)
	if err != nil {
		return msgErro
	}
	if salary < 150000 {
		return errors.New("Error: the salary entered does not reach the taxable minimum")
	}
	logrus.Infoln("Must pay taxes")
	return nil
}
