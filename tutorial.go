package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	game := "it home iron man"
	fmt.Println(game[8:])

	gender := "female"
	switch gender {
	case "female":
		fmt.Println("you are a girl")
	case "male":
		fmt.Println("your are a boy")
	default:
		fmt.Println("wtf")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// key value pairs
	kvs := map[string]string{
		"name":    "it home",
		"website": "https://ithelp.ithome.com.tw",
	}

	for key, value := range kvs {
		fmt.Println(key, value)
	}

}
