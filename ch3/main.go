package main

import (
    "fmt"
    "github.com/asmcos/requests"
)

func main() {
    req := requests.Requests()
    resp, _ := req.Get("https://baidu.com")
    fmt.Println(resp.R.StatusCode)
    fmt.Println(resp.R.Header["Content-Type"])
}
