package utils

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(to []string, otp string) error {

	var (
		fromGmail = "chodak158@gmail.com"
		password  = "avlh zuii dtss grxw"
	)

	smsHost := "smtp.gmail.com"
	port := "587"

	message := []byte("Your otp code is\n" + otp)

	byteM := []byte(message)

	auth := smtp.PlainAuth("", fromGmail, password, smsHost)

	err := smtp.SendMail(smsHost+":"+port, auth, fromGmail, to, byteM)

	if err != nil {
		log.Println("error", err)
		return err
	}

	fmt.Println("sms successfully")

	return nil
}
