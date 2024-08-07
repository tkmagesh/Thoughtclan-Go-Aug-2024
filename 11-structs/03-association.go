/*
Create the types to represent Employee & Organization
Employee is associated with One Organization
Organization can have multiple employees
*/

package main

import "fmt"

type Organization struct {
	Id        int
	Name      string
	Employees []*Employee
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	google := &Organization{Id: 900, Name: "google"}
	magesh := &Employee{Id: 100, Name: "magesh", Org: google}
	suresh := &Employee{Id: 101, Name: "suresh", Org: google}
	ganesh := &Employee{Id: 102, Name: "ganesh", Org: google}

	google.Employees = append(google.Employees, magesh, suresh, ganesh)

	microsoft := &Organization{Id: 901, Name: "microsoft"}
	ramesh := &Employee{Id: 103, Name: "ramesh", Org: microsoft}
	rajesh := &Employee{Id: 104, Name: "rajesh", Org: microsoft}
	microsoft.Employees = append(microsoft.Employees, ramesh, rajesh)

	// fmt.Println((*magesh).Name)
	fmt.Println(magesh.Name)

	/*
		fmt.Println(magesh.Org.Name)
		google.Name = "Google Inc."
		fmt.Println(magesh.Org.Name)

		fmt.Println(google.Employees[0].Name)
		magesh.Name = "Magesh Kuppan"
		fmt.Println(google.Employees[0].Name)
	*/
}
