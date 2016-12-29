package main

import gomail "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "sender@gmail.com")
	m.SetHeader("To", "receiver@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello there!")

	d := gomail.NewDialer("smtp.example.com", 587, "UserName", "PWD")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
