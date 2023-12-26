package shipment

import "fmt"

type Error struct {
	message     string
	insideError error
}

func (e *Error) Error() string {
	var res string
	if e.insideError != nil {
		res = fmt.Sprintf("%s: %s", e.insideError.Error(), e.message)
	} else {
		res = fmt.Sprintf("%s", e.message)
	}
	return res
}
