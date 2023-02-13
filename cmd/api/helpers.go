// FileName: cmd/api/healpers.go

package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParams(w http.ResponseWriter, r *http.Request) (int, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {

		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}