package main

import (
	"fmt"
	"go-core/09-ninth-lesson/pkg/ages"
	"go-core/09-ninth-lesson/pkg/older"
	"go-core/09-ninth-lesson/pkg/strwriter"
	"log"
	"os"
)

func main() {
	emp, emp_1 := ages.NewEmployee(52), ages.NewEmployee(13)
	customer, customer_1 := ages.NewCustomer(30), ages.NewCustomer(53)

	emp_2, emp_3 := older.Employee{Age: 52}, older.Employee{Age: 13}
	customer_2, customer_3 := older.Customer{Age: 30}, older.Customer{Age: 53}

	file, err := os.Create("./output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	str := "hello"

	strwriter.WriteStr(file, emp_2, str)
	result := ages.Oldest(customer_1, emp, emp_1, customer)
	result_1 := older.Eldest(emp_2, emp_3, customer_2, customer_3)
	fmt.Printf("Наибольший возраст: %v, %v \n", result, result_1)
}
