package main



//the error interface
type error interface{
	Error() string
}

//corret operation error == nil
//if not, the error message gerts printed



f, err := os.Open("path")

if err !=nil{
	fmt.Println(err)
	return 
}



