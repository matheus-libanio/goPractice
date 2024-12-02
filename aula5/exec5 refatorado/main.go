package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	inputError      = errors.New("unexpected input value")
	validationError  = errors.New("lower Value")
)

type Employee interface {
	VerifySalary() error
}

type salaryEmployee struct {
	category       string
	hoursPerMonth  float64
	monthlyRevenue float64
}

func main() {
	var emp Employee = &salaryEmployee{}

	err := emp.VerifySalary() // Call the method using the interface
	targetError := validationError

	if errors.Is(err, targetError) {
		logrus.Infoln(errors.Is(err, targetError), err)
	} else if err != nil {
		logrus.Error(err) // Handle other errors
	}
}

func (s *salaryEmployee) VerifySalary() error {
	fmt.Println("Enter salary category:")
	_, err := fmt.Scan(&s.category) // Pass the address
	if err != nil {
		return fmt.Errorf("Error: the category entered does not match our monetary method (%w)", inputError)
	}

	fmt.Println("Enter hours worked this month:")
	_, err = fmt.Scan(&s.hoursPerMonth) // Pass the address
	if err != nil {
		return fmt.Errorf("Error: the number of hours entered does not match our monetary method (%w)", inputError)
	}

	if s.hoursPerMonth < 80 {
		logrus.Infoln(errors.New("Error: the worker cannot have worked less than 80 hours per month"))
	}

	err = s.calculateSalary() // Calling the method of the instance
	if err != nil {
		return err
	}

	if s.monthlyRevenue < 150000 { // Use s.monthlyRevenue
		return fmt.Errorf("Error: the salary entered does not reach the taxable minimum (%w)", validationError)
	}
	logrus.Infoln("Must pay taxes equals 1/10 of the revenue: U$", s.monthlyRevenue/10)
	return nil
}

func (s *salaryEmployee) calculateSalary() error {
	switch s.category {
	case "A":
		s.monthlyRevenue = float64(3000*s.hoursPerMonth) * 1.5
	case "B":
		s.monthlyRevenue = float64(1500*s.hoursPerMonth) * 1.2
	case "C":
		s.monthlyRevenue = float64(1000 * s.hoursPerMonth)
	default:
		return errors.New("Error: category not found")
	}
	return nil
}