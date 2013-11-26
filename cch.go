package main

import (
	"fmt"
	"os"
	"flag"
	"net/http"
	"io/ioutil"
	"bytes"
)

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func Post(url string, data *string) {
	var b bytes.Buffer;
	b.Write([]byte(*data))
	resp, err := http.Post(url, "application/json", &b)
	if err != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func Put(url string, data *string) {
	client := &http.Client{}
	var b bytes.Buffer;
	b.Write([]byte(*data))
	req, err := http.NewRequest("PUT", url, &b)
	resp, err := client.Do(req)
	if err != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	cmd := ""
	url := ""
	var data *string
	var flags *flag.FlagSet
	if (len(os.Args) <= 2) {
		fmt.Println(len(os.Args))
		cmd = "default"
	} else {
		cmd = os.Args[1]
		url = os.Args[2]
		flags = flag.NewFlagSet(cmd, flag.ExitOnError)
		data = flags.String("d", "{}", "Request body in the form of JSON")
		flags.Parse(os.Args[2:])
	}

	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		if (flags == nil) {
			fmt.Println("[command] [url] [arguments]") 
			fmt.Println("Commands: [get|post|put|delete]")
		} else {
			flags.PrintDefaults()
		}
	}


	switch cmd {
	case "get":
		Get(url)
	case "post":
		Post(url, data)
	case "put":
		Put(url, data)
	default:
		Usage()
	}
}