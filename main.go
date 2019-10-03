package main

import (
	"fmt"
)

func main2(){
	s := Student{firstName: "sana",
		lastName:"aawan",
		contact: contactInfo{
			city:  "Pune",
			phone: 9876532133,
			email: "golang@gmail.com",
		},
	interest:interest{
		"Dancing",
		"Singing",
	},
		}
	//fmt.Printf("Student = %+v" , s)
	//s.print()
	s.getLastName()
	s.updateName("Annie")
	s.print()
}

func main(){
	colors := map[string]string{
		"orange": "#FFA500",
		"white": "#ffffff",
		"red": "#FF0000",
	}


	fmt.Println("/n Before Update /n")
	fmt.Println("Colours" , colors)
	correctrgbvalue(colors)
	fmt.Println("/n After Correction /n")
	fmt.Println("Colours", colors)
}

func (c colors) correctrgbvalue(strings map[string]string) {
	addr := colors["orange"]
	colors[""] ="ddss"
	colors["orange"] = addr
	for n, a := range colors {
		fmt.Printf("%v: %v\n", n, a)
	}
}
func main1() {
	fmt.Println("--testing --")
	writeIntToFile("gdeck")

	//readFile("gdeck")
}

