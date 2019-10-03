package main
import "fmt"

type person struct {
	n string
	age int
}
type intf interface {
	display()
}

func (p *person) display(){

	fmt.Println(p.n, p.age)
}

func main(){
	per := person{"sana", 24}
	per.display()
}

//func main(){
//	str := "training"
//	type mystring string
//	type mystr mystring
//	mystr = str
//	fmt.Println(str, mystr)
//}