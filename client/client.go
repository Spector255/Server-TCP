package main

import (
	"fmt"
	"net"
	"io/ioutil"
)

func command(filename string){
	loc, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	return loc
}

func main () {
	conn,err := net.Dial("tcp","127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	var s string
	
	fmt.Print("Enter action - file location (json/create.json): ")
	fmt.Scan(&s)
		
	var loc = command(s) 
	
	bytes := []byte(loc)
	
	
	conn.Write([]byte(loc))
	
	buf := make([]byte, 2000)
	
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(buf[:n]))
}
