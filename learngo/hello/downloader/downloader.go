package main

import (
	"fmt"
	"learngo/downloader/infra"
)

// func retriver(url string) string {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer resp.Body.Close()

// 	bytes, _ := ioutil.ReadAll(resp.Body)

// 	return string(bytes)
// }

// func getRetriver() infra.Retriver {
// 	return infra.Retriver{}
// }

// func getRetriver() testing.Retriver {
// 	return testing.Retriver{}
// }

func getRetriver() retriver {
	// return testing.Retriver{}//测试
	return infra.Retriver{}//真实数据
}

//是infra还是testing的Retriver呢？如何判断呢：通过接口

type retriver interface {
	Get(string) string
}

func main() {
	//var retriver infra.Retriver = getRetriver()
	// retriver := getRetriver()
	var r retriver = getRetriver()
	fmt.Printf("%s\n", r.Get("http://www.imooc.com"))

}
