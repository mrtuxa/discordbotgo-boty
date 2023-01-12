package webinterface

import (
	"log"
	"net/http"
)

func Start(port string) {
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
