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

	var taccounts []*model.Accounts
	for i := 0; i < 10; i++ {
		var taccount *model.Accounts

		gofakeit.Struct(&taccount)
		taccounts = append(taccounts, taccount)

		fmt.Println(taccounts)

		//fmt.Printf("テストユーザー%v番目スタート\n", i)
		//
		//tu = append(tu.Lastname)
		//users = append(tu.Firstname)
		//users = append(tu.Email)
		//users = append(tu.Password)
		//users = append(tu.Isactive)
		//users = append(tu.CreatedAt)
		//users = append(tu.UpdatedAt)

		//tu := model.Accounts{
		//	ID:        tu.Firstname,
		//	Lastname:  tu.Lastname,
		//	Firstname: tu.Firstname,
		//	Email:     tu.Email,
		//	Isactive:  tu.Isactive,
		//	CreatedAt: tu.CreatedAt,
		//	UpdatedAt: tu.UpdatedAt,
		//}
		//fmt.Println(users)

		//fmt.Println(tu)

	}

	fmt.Println("テストユーザー終了")

	var tp model.Post
	fmt.Println("テストポストスタート")

	for i := 0; i < 10; i++ {
		gofakeit.Struct(&tp)
		fmt.Printf("テストポスト%v番目スタート\n", i)
		fmt.Printf("ID:%v\n", tp.ID)
		//fmt.Printf("User:%v\n", tp.Account)
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
