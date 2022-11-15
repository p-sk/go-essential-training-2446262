// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}

// contentType will return the value of Content-Type header returned by making an
// HTTP GET request to url
func contentType(url string) (string, error) {
	// FIXME
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		ctype := res.Header.Get("Content-Tye")
		if ctype == "" {
			return ctype, fmt.Errorf("content-type Not found")
		}
		return ctype, nil
	}

	return "", nil
}
