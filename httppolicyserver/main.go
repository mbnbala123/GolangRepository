package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"policymanagement/gorm/DBStore"
	"policymanagement/gorm/orminterfaces"
)

func handlePolicyRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer RecoverFromPanic()
		var memberRepo orminterfaces.MemberRepo = nil
		memberInstance := DBStore.Member{}
		memberRepo = &memberInstance
		members, err := memberRepo.GetAllMembers()

		if err != nil {
			panic(err)
		}
		resp := map[string]any{
			"members": members,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func RecoverFromPanic() {
	if r := recover(); r != nil {
		log.Println("Recovered in RecoverFromPanic:", r)
	}
}
func main() {
	http.HandleFunc("/members", handlePolicyRequest())
	fmt.Println("Server started on port 7080")
	http.ListenAndServe(":7080", nil)
}
