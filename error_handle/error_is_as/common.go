package errorisas

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrorNotFound = errors.New("error not found")
var ErrorInvalid = errors.New("error invalid")

func testError() {
	e := ErrorNotFound
	if errors.Is(e, ErrorInvalid) {
		fmt.Println("e is ErrorNotFound")
		return
	} else {

		fmt.Println("e is not ErrorNotFound")
	}

	if errors.Is(e, ErrorNotFound) {
		fmt.Println("e is ErrorNotFound")
		return
	}
	// errors.As()
	//    errors.Is()
}
