package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	fmt.Println("enter your name!")
	var name string
	fmt.Scan(&name)

	fmt.Println("enter your email!")
	var email string
	fmt.Scan(&email)

	m := make(map[string]string)
	
	m["name"]= name
	m["address"] = email
	
	mba, err := json.Marshal(m)
	

	if err !=nil {
		fmt.Println("error: " , err)
	}
	
	fmt.Println("-------Marshaled Byte Array---------")

	fmt.Println(mba)
	
	fmt.Println("------UnMarshaled Map-----------")
	var m2 map[string]string

	err2 := json.Unmarshal(mba, &m2)
	
	if err2 !=nil {
	fmt.Println("error: " , err)
	}

	fmt.Println(m2)
}

