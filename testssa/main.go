package main

import "time"

func main() {
	var i = 4444
	go func() {
		time.Sleep(1*time.Second)
		println(i)
	}()
	i = 3333
	for  {

	}
}
