package fakedata

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	//"github.com/oklog/ulid"
	"github.com/uniqueisnfjllpbb/blogging/model"
	//"math/rand"
	//"time"
)

func FakeDataGenerator() {

	//t := time.Now()
	//entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	//id := ulid.MustNew(ulid.Timestamp(t), entropy)

	var tu model.User
	for i := 0; i < 10; i++ {

		gofakeit.Struct(&tu)
		fmt.Printf("テストユーザー%v番目スタート\n", i)
		fmt.Printf("ID:%v\n", tu.ID)
		fmt.Printf("LastName:%v\n", tu.LastName)
		fmt.Printf("FirstName:%v\n", tu.FirstName)
		fmt.Printf("Email:%v\n", tu.Email)
		fmt.Printf("Password:%v\n", tu.Password)
		fmt.Printf("IsActive:%v\n", tu.IsActive)
		fmt.Printf("CreatedAt:%v\n", tu.CreatedAt)
		fmt.Printf("UpdatedAt:%v\n", tu.UpdatedAt)

	}

	fmt.Println("テストユーザー終了")

	var tp model.Post
	fmt.Println("テストポストスタート")

	for i := 0; i < 10; i++ {
		gofakeit.Struct(&tp)
		fmt.Printf("テストポスト%v番目スタート\n", i)
		fmt.Printf("ID:%v\n", tp.ID)
		fmt.Printf("User:%v\n", tp.User)
		fmt.Printf("Title:%v\n", tp.Title)
		fmt.Printf("Body:%v\n", tp.Body)
		fmt.Printf("CreatedAt:%v\n", tp.CreatedAt)
		fmt.Printf("UpdatedAt:%v\n", tp.UpdatedAt)

	}

	fmt.Println("テストポスト終了")

}

//tu = append(tu, model.User{
//ID:        id,
//LastName:  gofakeit.LastName(),
//FirstName: gofakeit.FirstName(),
//Email:     gofakeit.Email(),
//Password:  gofakeit.Password(),
//IsActive:  gofakeit.Bool(),
//CreatedAt: gofakeit.Date(),
//UpdatedAt: gofakeit.Date(),
