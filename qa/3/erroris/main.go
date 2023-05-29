package main

import (
	"errors"
	"fmt"
)

var ErrSpec = errors.New("err test")
var ErrSpec2 = errors.New("err test2")

func main() {
	//	fmt.Println(RetErr())
	//	fmt.Println(RetErrSpec())

	err := RetErrSpec()
	if errors.Is(err, ErrSpec) {
		fmt.Println("err spec")
	} else {
		fmt.Println("not err spec")
	}

	err = RetErrSpec()
	if err == ErrSpec {
		fmt.Println("err spec")
	} else {
		fmt.Println("not err spec")
	}

	// errors.As(err, ErrSpec)

	err = RetWrap()
	if errors.Is(err, ErrSpec) {
		fmt.Println("err spec")
	} else {
		fmt.Println("not err spec")
	}

	if errors.Is(err, ErrSpec2) {
		fmt.Println("err spec2")
	} else {
		fmt.Println("not err spec2")
	}
	fmt.Println("|", err)
}

func RetErr() error {
	return errors.New("123")
}

func RetWrap() error {
	// f2: f1: err1

	return fmt.Errorf("RetErrSpec2: %w", RetErrSpec2())
}
func RetErrSpec() error {
	return ErrSpec
}

func RetErrSpec2() error {
	err := RetErrSpec()
	if err != nil {
		return errors.Join(err, ErrSpec2)
	}

	return nil
}
