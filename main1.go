package main

import "fmt"

type Human struct {
	FirstName string
	Phone     string
	Email     string
}

func (h *Human) SetName(name string) {
	h.FirstName = name
}

func (h *Human) SetPhone(phone string) {
	h.Phone = phone
}

type Action struct {
	FirstName  string
	SecondName string
	H          Human
	Human
}

func (a *Action) SetName(name string) {
	a.FirstName = name
}

func main() {
	var action Action = Action{
		FirstName:  "Taslav",
		SecondName: "Konrad",
		//Можно заполнять переменную, как через идентификаторы структуры, так и без них.
		//Только если без идентификаторов нужно будет писать значения строго, как заданы в объявлении структуры, а с идентификаторами в каком угодно порядке
		Human: Human{
			"Kate",
			"+7999999999",
			"kate@ts.com",
		},
		H: Human{
			FirstName: "Artem",
			Phone:     "+111111111111",
			Email:     "artem@nur.com",
		},
	}

	// при наличии одикаковых полей, приоритет у поля более верхнего уровня
	fmt.Println(action.FirstName) // Taslav

	// чтобы обратиться к полю встроенной структуре, нужно указать ее тип, затем поле
	fmt.Println(action.Human.FirstName) // Kate

	// если поле уникальное, то можно напрямую обращаться к нему из родительской структуры
	fmt.Println(action.Email) // kate@ts.com

	// чтобы обратиться к полю типа структуры, нужно указать имя поля, затем необходимое поле вложенной структуры
	fmt.Println(action.H.FirstName) // Artem

	// методы встроенной структуры наследуются родительской
	action.SetPhone("+77777777")
	fmt.Println(action.Phone) // +77777777

	// при наличии одинаковых методов, приоритет у метода структуры более верхнего уровня
	action.SetName("The newest name")
	fmt.Println(action.FirstName)       // The newest name
	fmt.Println(action.Human.FirstName) // Kate

	// чтобы вызвать этот метод у встроенной структуры, необходимо обратиться к ней через тип
	action.Human.SetName("Volodya")
	fmt.Println(action.Human.FirstName) // Volodya

	//две одинаковых безпеременных структур создать нельзя в род. структуре!
}
