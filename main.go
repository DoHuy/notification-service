package main

import (
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
	//"path/filepath"
	firebase "firebase.google.com/go"
	//"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

//package main
//
//import (
//"context"
//firebase "firebase.google.com/go"
//"firebase.google.com/go/messaging"
//"fmt"
//"log"
//)


type AdapterMessaging struct  {
	registraionTokens []string
	app *firebase.App
}

func NewAdapter(registrationToken []string, app *firebase.App) *AdapterMessaging {
	return &AdapterMessaging{
		registraionTokens: registrationToken,
		app: app,
	}
}

func (adapter *AdapterMessaging) SendMessageToSpecifyDevice() error{
	ctx := context.Background()
	client, err := adapter.app.Messaging(ctx)
	if err != nil {
		log.Fatal("error getting Messaging Client %v \n", err)
		return err
	}
	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Test tin nh",
			Body:  "Thịnh cái này là test Trang",
		},
		Token: adapter.registraionTokens[0],
	}
	resp, err := client.Send(ctx, msg)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	fmt.Println("Success push message to client", resp)
	return nil
}




func initializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account_golang]

	opt := option.WithCredentialsFile("/home/huydv/golang/src/notification-service/service-account-credential.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_service_account_golang]

	return app
}




const DEVICE_TOKEN = "e0sfg-i3IlY:APA91bH6quFlCxoW367vltrIf6f0MgGFkz6JBCF9wg7lNzNBbdifZzFW4MJ6BWVmACOMONzKYmipurt4-y_cZFzuKw55lyxgesp3dl1ifd-5a4T4t5SGpj5nuu0fCXI4Q-5leSWPvPuL"
func main () {
	app := initializeAppWithServiceAccount()
	var b = []string{DEVICE_TOKEN}
	adapter := NewAdapter(b, app)

	var error = adapter.SendMessageToSpecifyDevice()
	if error != nil {
		fmt.Errorf("%v", error)
	}


}