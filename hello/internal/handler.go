package internal

import (
	
	"fmt"
	"net.http"

)

type handler struct {

	return &handler{}
}

func (h *handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprint(w "Hello, %s!", r.URL.Path[1:])

}