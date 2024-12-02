package main

import (
	"errors"
	"fmt"
)

// we define a type struct
type myError struct {
	msg string
	x string
 }
 //we make our type struct implement the Error() method
 func (e *myError) Error() string {
	return fmt.Sprintf("An error occurred: %s, %s", e.msg, e.x)
 }
 
 func main() {
	e := &myError{"new error", "404"}
	
	var err *myError
 
	isMyError := errors.As(e, &err)// compares errors
 
	fmt.Println(isMyError)//prints true because the errors match
 }
 