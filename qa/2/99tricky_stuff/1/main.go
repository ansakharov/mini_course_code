package main

import (
	"fmt"
)

//Пример для разбора
// 1. // Что выведет код и почему?

func setLinkHome(link *string) {
	*link = "http://home"

	/*	str := "http://home"
		link = &str*/

}

func main() {

	link := "http://other"
	fmt.Println(link)

	setLinkHome(&link)
	fmt.Println(link)
}
