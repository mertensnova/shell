package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	
	//  Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}	

	// Store the .env values in a variable
	var ACCOUND_SID string = os.Getenv("ACCOUND_SID")
	var ACCOUND_KEY string = os.Getenv("ACCOUND_KEY")
	var FROM string = os.Getenv("FORM")

	// Initialize twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:  ACCOUND_SID,
		Password: ACCOUND_KEY,
	})

	// Put all the number in an array
	sent_to := []string{"+966558581107"} 


	// Loop throuh the array
	for _,num := range sent_to{
		params := &openapi.CreateMessageParams{}
		// Set the numbers
		params.SetTo(num)
		params.SetFrom(FROM)
		// Set the body
		params.SetBody("Hello World")
		_, err := client.Api.CreateMessage(params)
		// Catch error
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Message sent successfully!")
		}
	}
}
