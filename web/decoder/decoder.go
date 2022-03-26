package decoder

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func IDInt64(r *http.Request) (int64, error) {
	v := mux.Vars(r)
	idStr := v["id"]
	if idStr == "" {
		return 0, errors.New("id is required")
	}

	ids, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || ids == 0 {
		return 0, errors.New("id is invalid")
	}

	return ids, nil
}
