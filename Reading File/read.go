package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println("Error Found : ",err)
		return
	}
	defer file.Close() // It will run at the end at any cost. Like Finally in java

	fileinfo, err := file.Stat() // its check the status of file .. Like it is readable or not .. data is exist or not
	if err != nil {
		fmt.Println("Status is wroung : ",err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("bytes read: ", bytesread)
fmt.Println("bytestream to string: ", string(buffer))
}
