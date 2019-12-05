package main

//noinspection GoUnresolvedReference
import (
	"github.com/mozillazg/request"
	"net/http"
)


func main() {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get("http://httpbin.org/get")
	j, err := resp.Json()
	println(j,err)
}
