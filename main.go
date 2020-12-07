package main

import (
	"context"
	"fmt"
	"github.com/Nilsas/passwordstate-sdk-go/sdk/secret"
	"os"
)

func main() {
	client := secret.NewClient(os.Getenv("PASS_URL"), os.Getenv("PASS_KEY"), secret.DefaultHTTPClient)
	//password := secret.Secret{
	//	PasswordListID: 2,
	//	Title:          "Made with Go SDK",
	//	UserName:       "welp",
	//	Password:       "my-password",
	//	APIKey:         client.Key,
	//}
	//
	//err := client.AddPassword(context.TODO(), &password)
	//if err != nil {
	//	fmt.Println(err)
	//}

	pass, err := client.GetPassword(context.TODO(), 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
}
