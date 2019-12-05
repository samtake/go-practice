package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriver interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriver) string {
	return r.Get(url)
}

func poster(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":     "sam",
			"language": "golang",
		})
}

type RetriverPoster interface {
	Retriver
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another facked imooc.com",
	})
	return s.Get(url)
}

func main() {

	var r Retriver

	retriver := mock.Retriver{"this is a fake imooc.com"}
	r = &retriver

	fmt.Printf("类型=%T 值=%v\n", r, r)
	r = &real.Retriver{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("类型=%T 值=%v\n", r, r)
	fmt.Println()
	inspect(r)

	//tye assertion

	// //real
	// realRetriver := r.(*real.Retriver)
	// fmt.Println(realRetriver.TimeOut)

	// //出错写法
	// realRetriver := r.(real.Retriver)
	// fmt.Println(realRetriver.TimeOut)

	//防止出错
	if mockRetriver, ok := r.(*mock.Retriver); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not a mock retriver")
	}
	//fmt.Println(download(r))

	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&retriver))
}

func inspect(r Retriver) {
	fmt.Printf("类型=%T 值=%v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriver:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
