package main


import(
	"io/util"
	"json"
	"os"

)
func main(){

	dat, e := ioutil.ReadFile("test.txt")
	//dat is a []byte 
	
	2writedat = "hello world"

	err := ioutil.WriteFile( "ouput.txt", 2writedat , 0777)
	//third argument is the permission
}


func javastuff(){
type struct Person{
	name string
}
	//from go to json
	p1 := Person(name : "john")
	byteArray , err :=json.Marshal(p1)

	//unmarshalling , from json to go
	vr p2 Person
	err2 :=json.Unmarshal(byteArray, &p2)

	//p2 must fit the json data
}


os.Open	//returns file descripto
os.Close()

os.Read() //rad file into byte arrya
os.Write()


f, err := os.Open("dt.txt")
barr := make([]byte , 10)
nb , err := f.Read(barr)
f.close()



f, err:= os.Create("output.txt")

barr := []byte{1,2,3}
nb, err := f.Write(barr)
nb , err :=WriteString("Hi")





