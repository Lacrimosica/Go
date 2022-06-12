package main

func main(){
	//slices have variable size, up to the whole array 
	arr := [...]string {"a", "b" , "c" , "d" , "e" , "f" , "g"}
	s1:= arr[1:3]
	s2 := arr[2:5]


	//capacity and length are different for slices

	//writing to a slice changes underlying array

	//alice literals

	sli := []int{1,2,3}	//we don't put anything in the brakcets, it makes an underlying array and makes the slice be the whole array


	//using make to make slices without using literals

	sli2 := make( []int, 10)
	sli3 := make ([]int, 10, 15) //length and capacity defined separately



	//append, adds until the capacity and then adds even more if we reach the boundary

	sli = append(sli, 100) //100 is the neew number you are adding to the int slice

}
