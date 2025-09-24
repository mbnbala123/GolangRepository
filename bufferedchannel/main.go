package main

import (
	"policymanagement/models"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {

	claimchannel := make(chan models.Claim, 5)
	//claim array
	var claims []models.Claim

	go func() {
		for i := 0; i < 10; i++ {
			claim := models.Claim{
				ID:     uint(gofakeit.IntRange(1, 100)),
				Amount: gofakeit.IntRange(1000, 10000),
			}
			claims = append(claims, claim)

		}

	}()
	go func() {
		//claim raised by the client
		println("Raising claims...")
		count := 0
		for _, claim := range claims {
			claimchannel <- claim
			count++
			println("Claims raised:", count)
		}

	}()
	//claim processed by the server
	println("Processing claims...")
	for i := 1; i <= 10; i++ {
		claim := <-claimchannel
		println("Claim ID:", claim.ID, "Amount:", claim.Amount)
	}

}
