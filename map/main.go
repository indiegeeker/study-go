package main

import "fmt"

type test_map struct {
	Lat, Long float64
}

var t_map map[string]test_map

func main() {
	t_map = make(map[string]test_map)
	t_map["Bell Labs"] = test_map{
		40.68433, -74.39967,
	}
	fmt.Println(t_map["Bell Labs"])
}
