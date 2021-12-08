package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url, appcat string, content []byte) []byte {
	r, err := http.Post(url, appcat, bytes.NewReader(content))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return b
}
