package src

import (
	"fmt"
	"io"
	"net/http"
)

func GetRequest() {
	req, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
