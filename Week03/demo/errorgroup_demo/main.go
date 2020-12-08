package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main(){
	g:= new(errgroup.Group)
	var urls = []string{
		"http://www.sina.com",
		"http://www.cisco.com",
		"http://www.baidu.com",
	}

	for _,url := range urls{
		url := url
		g.Go(func() error{
			resp,err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}

	// wait for all http fetches to complete
	// 第一个报错，就会返回
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all Urls.")
	}
}