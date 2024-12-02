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

type salaryEmployee struct{
	category string
	hoursPerMonth	float64
	monthlyRevenue float64
}

type Employee interface {
	VerifySalary() error
}

func main()  {
	// Cria uma instância do salaryEmployee
	employee := &salaryEmployee{}

	err := employee.verifySalary() // Chama o método corretamente
	targetError := validationError

	if errors.Is(err, targetError) {
		logrus.Infoln(errors.Is(err, targetError), err)
	} else if err != nil {
		logrus.Error(err) // Trata outros erros
	}
}

func (s *salaryEmployee)verifySalary() error {
	fmt.Println("Enter salary category:")
	_, err := fmt.Scan(&s.category) // Modificado para passar o endereço
	if err != nil {
		return fmt.Errorf("Error: the category entered does not match our monetary method (%w)", inputError)
	}

	fmt.Println("Enter hours worked this month:")
	_, err = fmt.Scan(&s.hoursPerMonth) // Modificado para passar o endereço
	if err != nil {
		return fmt.Errorf("Error: the number of hours entered does not match our monetary method (%w)", inputError)
	}
	if s.hoursPerMonth < 80 {
		logrus.Infoln(errors.New("Error: the worker cannot have worked less than 80 hours per month"))
	}

	err = s.calculateSalary() // Chamando o método da instância
	if err != nil {
		return err
	}

	if s.monthlyRevenue < 150000 { // Usando s.monthlyRevenue
		return fmt.Errorf("Error: the salary entered does not reach the taxable minimum (%w)", validationError)
	}
	logrus.Infoln("Must pay taxes equals 1/10 of the revenue: U$", s.monthlyRevenue/10)
	return nil
}

func (s *salaryEmployee)calculateSalary () error{
	switch s.category {
	case "A":
		s.monthlyRevenue = float64(3000*s.hoursPerMonth) * 1.5
		return nil
	case "B":
		s.monthlyRevenue = float64(1500*s.hoursPerMonth) * 1.2
		return nil
	case "C":
		s.monthlyRevenue = float64(1000 * s.hoursPerMonth)
		return nil
	default:
		return errors.New("Error: category not found")
	}
}