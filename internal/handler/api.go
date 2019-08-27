package handler

import (
	"fmt"
	"github.com/sanity-io/litter"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, litter.Sdump(r))
}
