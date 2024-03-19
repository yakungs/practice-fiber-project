package handler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

func handlerError(w http.ResponseWriter, err error) {

	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprint(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, e)
	}

}
