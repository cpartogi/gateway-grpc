package constant

import "fmt"

var (
	ErrNotFound   = fmt.Errorf("data not found")
	ErrConflict   = fmt.Errorf("conflict, data already exist")
	ErrBadRequest = fmt.Errorf("bad request")
	ErrForbidden  = fmt.Errorf("forbidden")
)
