package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func get() {
	res, _ := http.Get("http://goinpracticebook.com")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", b)
}

func delete() {
	req, _ := http.NewRequest("DELETE", "http://example.com/foo/bar", nil)
	res, _ := http.DefaultClient.Do(req)
	fmt.Printf("%s\n", res.Status)
}

func simpleUserClient() {
	cc := &http.Client{Timeout: time.Second}
	res, err := cc.Get("http://goinpracticebook.com")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", b)

}

func main() {
	simpleUserClient()
}
