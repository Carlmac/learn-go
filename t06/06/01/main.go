package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	// Gets an item from the os.Args string slice:
	//     os.Args[INDEX]
	// INDEX can be 0 or greater
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("3rd argument:", os.Args[3])

	// `len` function can find how many items
	// inside a slice value
	fmt.Println("Items inside os.Args:", len(os.Args))
}
