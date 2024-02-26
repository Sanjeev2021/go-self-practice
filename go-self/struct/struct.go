package main

import (
	"time"
	"fmt"
)

func main() {

	lastPurchasedTime := time.Date(2022, time.February, 25, 15, 30, 0, 0, time.UTC)
	
	sanjeev := Customer {
		Name : "sanjeev",
		Age :  23,
		PhoneNumber:  1234546789,
		Lastpurchased:  lastPurchasedTime,
		TotalPurchased : 7,
		Status: true,
		social : SocialMedia{
			Facebook : "Sanjeev",
			Instagram : "sanjev__10",
		},
	}

	fmt.Println("the customer sanjeev", sanjeev)
	fmt.Println("sanjeev ka fb", sanjeev.social.Facebook)
	fmt.Println("sanjeev ka insta", sanjeev.social.Instagram)

	

}

type SocialMedia struct {
	Facebook string
	Instagram string
}


type Customer struct {
	Name     string
	Age      int
	PhoneNumber int
	Lastpurchased time.Time
	Email string
	TotalPurchased int
	Status bool
	social SocialMedia
}