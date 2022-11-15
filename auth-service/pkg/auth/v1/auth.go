package v1

import (
	"fmt"
	"log"
	"net/http"
)

type HttpPassthroughService struct{}

const ServerPort = 9001

func (h *HttpPassthroughService) StartServer() {
	fmt.Printf("Listening on port %d for auth requests\n", ServerPort)
	handler := func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("received request with url: %s, with headers %+v\n", r.URL.String(), r.Header)
		fmt.Println("exchanging API Key for access token ...")
		rw.Header().Set("Authorization", "Bearer abc123xyz890")
	}
	address := fmt.Sprintf(":%d", ServerPort)
	err := http.ListenAndServe(address, http.HandlerFunc(handler))
	if err != nil {
		log.Fatal(err.Error())
	}
}
