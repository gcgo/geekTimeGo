package main

import "fmt"

func main() {
	map1 := make(map[string]string, 5)
	map1["张三"] = "开心1"
	map1["张三"] = "开心2"
	map1["张五"] = "开心3"
	map1["张六"] = "开心4"

	for key, val := range map1 {
		fmt.Println(key, ": ", val)
	}

}
