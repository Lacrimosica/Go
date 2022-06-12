package main


import( 
	"fmt"
	"sort"
	"strconv"
	"strings"
	)

func main(){
	s:= make([]int, 0 , 3)
	
	var input string 

	for {
		fmt.Scan(&input)
		if strings.ToLower(input)[0] == 'x' {
			break
	}
	
	num, _ := strconv.Atoi(input)

	s = append(s, num)
	sort.Ints(s)
	fmt.Println(s)
	}
}
