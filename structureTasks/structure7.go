package structureTasks

import "fmt"

// 7. Вложенные структуры
// Создайте структуру Company, которая будет содержать поле Name (тип string) и вложенную структуру Address с полями
// City и Country (оба типа string). Создайте экземпляр структуры Company, заполните поля и выведите результат на экран.
type Address struct {
	City    string
	Country string
}
type Company struct {
	Name    string
	Address Address
}

func structure7() {

	comp := Company{Name: "DanyaAndOlejkaLaserPashkaRentalGroupIncorporated",
		Address: Address{City: "Almaty", Country: "KZ"}}

	fmt.Printf("Company name: %v\nCompany address: %v, %v", comp.Name, comp.Address.City, comp.Address.Country)
}
