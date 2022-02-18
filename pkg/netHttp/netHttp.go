package netHttp

import (
	"fmt"
	"io/ioutil"
	"net/http" // https://gobyexample.com/http-clients
)

func GetBody(url string) string {
	fmt.Println("call URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		panic(err)
	}

	return string(body)
}
