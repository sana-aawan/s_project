package main

import "fmt"
type interest []string

type interestStud struct{
	i interest
}
type contactInfo struct{
	city string
	phone int
	email string
}
type Student struct {
	firstName string
	lastName string
	std int
	aggregate float64
	contact contactInfo
	interest
}

func (s Student) print(){
	fmt.Printf("Student = %+v",s)
}

func (s Student) getLastName(){
	if s.lastName != "" {
		fmt.Printf("Lastname = %+v ", s.lastName)
	} else {
		fmt.Printf("Lastname is empty")
	}
}

func (s *Student) updateName(fname string){
	if s.firstName != "" {
		s.firstName = fname
	}
}

func (s *Student) updateInterest(inter interest){
	s.interest = inter
}
