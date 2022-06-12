pakcage main

func main(){
	//maps are the implementation of the hash table

	//using make
	vat idMap map[string]int
	idMap = make(map[string]int)


	//map literal

	idMap := map[string]int {
		"joe" : 123
	}

	//two value assignment test

	id, p : idMap["joe"]
	//id is value p is presence of the key

	len(idMap)


	//iteration

	for key,val : = range idMap {
		fmt.Println(key,val)
	}



}
