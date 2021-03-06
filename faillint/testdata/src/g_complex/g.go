package g

import (
	"errors"
	"fmt"
)

func foo() error {
	// Multiple on single line.
	defer func() {
		_ = fmt.Errorf(fmt.Sprintf("%v", fmt.Errorf("err"))) // want `declaration "Errorf" from package "fmt" shouldn't be used, suggested: "github.com/pkg/errors.{Errorf}"` `declaration "Errorf" from package "fmt" shouldn't be used, suggested: "github.com/pkg/errors.{Errorf}"`
	}()

	// Multiple on different lines.
	defer func() {
		_ = fmt.Errorf( // want `declaration "Errorf" from package "fmt" shouldn't be used, suggested: "github.com/pkg/errors.{Errorf}"`
			fmt.Sprintf(
				"%v",
				fmt.Errorf("err"), // want `declaration "Errorf" from package "fmt" shouldn't be used, suggested: "github.com/pkg/errors.{Errorf}"`
			),
		)
	}()
	return errors.New("bar!")
}
