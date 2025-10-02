package main

import (
	"log"
	"net/http"
	_ "policymanagement/claimapi/docs"
	"policymanagement/claimapi/store"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Claim API
// @version 1.0
// @description This is api service for managing Claims
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email mbnbala@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7082
// @BasePath /
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /claims/v1.0", store.GetClaims)
	mux.HandleFunc("POST /claims/v1.0", store.SaveClaim)
	mux.HandleFunc("GET /claims/v1.0/{claimId}", store.GetClaimByID)
	mux.HandleFunc("PUT /claims/v1.0/{id}", store.UpdateClaim)
	mux.HandleFunc("DELETE /claims/v1.0/{id}", store.DeleteClaim)

	// Swagger UI served at /swagger/
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Your own handlers
	// mux.HandleFunc("/claims", claimsHandler)

	log.Println("Server running at http://localhost:7082")
	log.Fatal(http.ListenAndServe(":7082", mux))

}
