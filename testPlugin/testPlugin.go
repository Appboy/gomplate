package main

import (
  "net/http"
  "net/url"
)

func Scheme() (string, error) {
  return "testname", nil
}

func Get(url *url.URL, headers http.Header, args []string) ([]byte, error) {
  a := []byte("testing byte array...")
  return a, nil
}

func GetMediaType() (string) {
	return "text/plain"
}

func main() {
	//do nothing
}
