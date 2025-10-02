package main

import (
	"encoding/json"
	"log"
	"net/http"
	"policymanagement/policyapi/interfaces"
	"policymanagement/policyapi/models"
)

func handlePolicyRequest() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var claimRepo interfaces.IClaimRepo = nil
		claimInstance := models.Claim{}
		claimRepo = &claimInstance
		claims, err := claimRepo.GetAllClaim()
		if err != nil {
			log.Println("Error fetching claims:", err)
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		resp := map[string]any{
			"claims": claims,
		}

		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(resp)
	}
}

func main() {
	http.HandleFunc("/claims", handlePolicyRequest())
	log.Println("Server started on port 7081")
	http.ListenAndServe(":7081", nil)
}
