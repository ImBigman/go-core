package ages

// Ager - абстрактный интерфейсный тип данных.
// Объявляет контракт на получение возраста.
type Ager interface {
	Age() int // контракт интерфейса
}

// тип для работников.
type Employee struct {
	age int
}

// тип для заказчика.
type Customer struct {
	age int
}

// для получения возраста работника.
func (e Employee) Age() int {
	return e.age
}

// для получения возраста клиента.
func (c Customer) Age() int {
	return c.age
}

// конструкторы для структур
func NewEmployee(n int) Employee {
	return Employee{age: n}
}
func NewCustomer(n int) Customer {
	return Customer{age: n}
}

// для возвращения самого большого возраста
func Oldest(arg ...Ager) int {
	var res int
	for _, value := range arg {
		age := value.Age()
		if res < age {
			res = age
		}
	}
	return res
}
