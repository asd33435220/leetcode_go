package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	res := 0
	var idList []int
	idList = append(idList, id)
	for len(idList) > 0 {
		thisId := idList[len(idList)-1]
		idList = idList[:len(idList)-1]
		for _, e := range employees {
			if e.Id == thisId {
				res += e.Importance
				idList = append(idList, e.Subordinates...)
			}
		}
	}
	return res
}
func main() {
	e1 := &Employee{
		Id:           1,
		Importance:   5,
		Subordinates: []int{2, 3},
	}
	e2 := &Employee{
		Id:           2,
		Importance:   3,
		Subordinates: []int{},
	}
	e3 := &Employee{
		Id:           3,
		Importance:   3,
		Subordinates: []int{},
	}
	emploeeList := []*Employee{e1, e2, e3}
	res := getImportance(emploeeList, 1)
	fmt.Println("res", res)
}
