package main

import "time"

func Retrieve(start string) []string {
	datas := map[string][]string{
		"snapp.ir/":     {"snapp.ir/blog", "snapp.ir/about", "snapp.ir/contact"},
		"snapp.ir/blog": {"snapp.ir/blog/post1", "snapp.ir/blog/post2"},
	}
	time.Sleep(time.Second)
	result, ok := datas[start]
	if !ok {
		return []string{}
	} else {
		return result
	}
}
