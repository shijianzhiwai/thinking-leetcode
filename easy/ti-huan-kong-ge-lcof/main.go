package main

import "fmt"

func main() {
	fmt.Println(replaceSpace("We are happy."))
}


func replaceSpace(s string) string {
	bs := []byte(s)
	ret := make([]byte, 0, len(s))
	for _, char := range bs {
		if char == ' ' {
			ret = append(ret, []byte("%20")...)
		} else {
			ret = append(ret, char)
		}
	}
	return string(ret)
}