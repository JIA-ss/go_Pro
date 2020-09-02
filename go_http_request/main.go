package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type ddd struct {
	Date string `json:"date"`
	Version string `json:"version"`
	Name string `json:"name"`
	Info string `json:"info"`
}
type sss struct {
	Errno int `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data  ddd `json:"data"`
}

func test() (rsp_ sss){
	url := "http://10.240.85.140:8080/ckv/version/query"
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("key", "value")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	read_, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(read_))
	json.Unmarshal(read_, &rsp_)
	fmt.Println(rsp_)
	return
}

func main(){
	test_ := test()
	fmt.Println(test_)
}