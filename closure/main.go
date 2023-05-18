package main

import (
	"log"

	"github.com/pkg/errors"
)

func main() {
	errs := evalutate(56)

	log.Println(errs)
}

func evalutate(x int) []error {
	aggregateErrors := aggregateErrorClosure()

	if x > 10 {
		err := errors.Errorf("%d is over 10\n", x)
		aggregateErrors(&err)
	}

	if x > 42 {
		err := errors.Errorf("%d is over 42\n", x)
		aggregateErrors(&err)
	}

	return aggregateErrors(nil)
}

func aggregateErrorClosure() func(err *error) []error {
	var errs []error

	return func(err *error) []error {
		if err != nil {
			errs = append(errs, *err)
		}
		return errs
	}
}
