package entity

import (
	"net/http"
	"strconv"
)

var (
	BadRequestError = strconv.Itoa(http.StatusBadRequest)
	NotFoundError   = strconv.Itoa(http.StatusNotFound)
	InternalError   = strconv.Itoa(http.StatusInternalServerError)
)
