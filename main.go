package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
)


func main()  {
	bufferReader := bufio.NewReader(os.Stdin)
	from := os.Getenv("FROM")
	passwd := os.Getenv("PASSWD")

	to := [][]string{}

  // smtp server configuration.
  smtpHost := "smtp.gmail.com"

  // Message.
  fmt.Printf("Enter your message:\t")
  msg, err := bufferReader.ReadString('\n')
  
  // Authentication.
  auth := smtp.PlainAuth("", from, passwd, smtpHost)

   // Sending email.
   for _, v := range to {
	   err1 := smtp.SendMail("smtp.gmail.com:587", auth, from, v, []byte(msg))
	   if err1 != nil {
		fmt.Println(err)
		return
	  }
   }


   fmt.Println("Email Sent Successfully!")

}