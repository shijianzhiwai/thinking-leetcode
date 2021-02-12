package main

import "time"

func main() {
	var i = 1
	go func() {
		time.Sleep(1*time.Second)
		println(i)
	}()
	i = 2
	for  {
	}
}
