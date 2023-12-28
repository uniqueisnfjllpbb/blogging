package fakedata

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/uniqueisnfjllpbb/blogging/model"
)

func FakeDataGenerator() {

	var tu model.User
	gofakeit.Struct(&tu)

	fmt.Println(tu.ID)
	fmt.Println(tu.LastName)
	fmt.Println(tu.FirstName)
	fmt.Println(tu.Email)
	fmt.Println(tu.Password)
	fmt.Println(tu.IsActive)
	fmt.Println(tu.CreatedAt)
	fmt.Println(tu.UpdatedAt)

}
