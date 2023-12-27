package fakedata

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/uniqueisnfjllpbb/blogging/model"
)

func FakeDataGenerator() {

	// 関数で生成
	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.Color())

	// Structのフィールドにセット
	var testuser model.User
	var testpost model.Post

	//gofakeit.Struct(&testuser) // ランダムな値でPerson構造体を生成

	for i := 0; i < 10; i++ {
		//テストユーザー作成
		//tpost := gofakeit.Struct(&testpost)
		gofakeit.Struct(&testuser) //テストユーザー作成

		fmt.Println("苗字:", testuser.LastName+string(rune(i)))
		fmt.Println("名前:", testuser.FirstName+string(rune(i)))
		fmt.Println("パスワード:", testuser.Password+string(rune(i)))
		fmt.Println("メールアドレス:", testuser.Email+string(rune(i)))

		//テスト記事作成

		//tuser := gofakeit.Struct(&testuser)
		gofakeit.Struct(&testpost) //テスト記事作成

		fmt.Println("苗字:", testpost.ID+string(rune(i)))
		fmt.Println("名前:", testpost.Title+string(rune(i)))
		fmt.Println("パスワード:", testpost.Body+string(rune(i)))
		fmt.Println("メールアドレス:", testpost.CreatedAt+string(rune(i)))
		i++
	}

}
