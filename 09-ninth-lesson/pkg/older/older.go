package older

// тип для работников.
type Employee struct {
	Age int
}

// тип для заказчика.
type Customer struct {
	Age int
}

// для возвращения объекта со самым большим возрастом
func Eldest(arg ...interface{}) interface{} {
	var old, age int
	var res interface{}
	for _, v := range arg {
		if employee, ok := v.(Employee); ok {
			age = employee.Age
		}
		if customer, ok := v.(Customer); ok {
			age = customer.Age
		}
		if old < age {
			old, res = age, v
		}
	}
	return res
}
