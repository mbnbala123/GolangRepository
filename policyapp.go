package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"policymanagement/gocharts/store"
	"policymanagement/gorm/DBStore"
	"policymanagement/gorm/orminterfaces"
	"policymanagement/interfaces"
	"policymanagement/models"
	"policymanagement/utility"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {

	//VehicleUpdate()
	//	UnmarshalJsonFile()
	//ConvertJsons()
	//ReadMemberDetailsfromCSV()
	//PolicyUpdate()
	//writeCSVFile()
	//Recursive()
	//CreateConnection()
	//SaveMember()
	PrintAllMembers()
}
func CreateChart() {
	models := []string{"BMW", "AUDI", "HYUNDAI", "KIA", "FORD"}
	claims := []store.Claim{}

	//generate 50 claims with random amounts
	for i := 0; i < 25; i++ {
		claim := store.Claim{
			ID:        uint(i + 1),
			ModelName: models[gofakeit.Number(0, len(models)-1)], // Random model
			Amount:    gofakeit.IntRange(100000000, 1000000000),  // Random amount
		}
		claims = append(claims, claim)
	}

	store.GenerateBarGraphClaims(claims, models)
}
func CreateConnection() {
	DBStore.CreateConnection()

}
func RecoverFromPanic() {
	if r := recover(); r != nil {
		log.Println("Recovered in RecoverFromPanic:", r)
	}
}

func PrintAllMembers() {
	var memberRepo orminterfaces.MemberRepo = nil
	memberInstance := DBStore.Member{}
	memberRepo = &memberInstance
	defer RecoverFromPanic()

	// Fetch all members
	members, err := memberRepo.GetAllMembers()
	if err != nil {
		panic(err)
	}
	for _, m := range members {
		println("Member:", m.FirstName, m.LastName, m.Email)
	}
}
func SaveMember() {

	var memberRepo orminterfaces.MemberRepo = nil
	memberInstance := DBStore.Member{
		FirstName: "Balasubramani",
		LastName:  "Muthusamy",
		Email:     "Balasubramani.Muthusamy@example.com",
		Phone:     "123-456-7890",
		Address:   "123 Main St, Anytown, USA",
		StartDate: "2020-01-01",
		EndDate:   "2020-12-31",
	}
	memberRepo = &memberInstance
	success, err := memberRepo.SaveMember()
	if err != nil {
		panic(err)
	}
	if success {
		println("Member saved successfully")
	}

}

func VehicleUpdate() {

	var vehicleRepo interfaces.IVehicleRepo = nil

	vehicle := &models.Vehicle{
		LicensePlateNo:     "ZH8B4P",
		Maker:              "Hyundai",
		Model:              "Santafe",
		DateOfRegistration: time.Now(),
		VIN:                "3434vdfgdfg44534",
		FuelType:           models.Petrol,
		EngineNo:           "32434534345345",
		Color:              "blue",
	}

	jsonString, _ := json.Marshal(vehicle)
	fmt.Printf("JSON string %s\n", string(jsonString))

	vehicleRepo = vehicle

	result, _ := vehicleRepo.Save()
	println(result)

	data, error := vehicleRepo.GetByID(vehicle.LicensePlateNo)
	if error != nil {
		println(error.Error())
	}

	for key, value := range utility.StructToMapVehicle(data) {

		fmt.Printf("%s : %v\n", key, value)
	}

	vehicles, _ := vehicleRepo.GetAll()
	fmt.Println("All vehicles:")
	for _, v := range vehicles {
		for key, value := range utility.StructToMapVehicle(v) {

			fmt.Printf("%s : %v\n", key, value)
		}
	}

}
func PolicyUpdate() {

	var policyRepo interfaces.IPolicyHolderRepo = nil

	policyHolder := &models.PolicyHolder{
		PolicyNumber: "4534534534",
		FirsName:     "Balasubramani",
		LastName:     "Muthusamy",
		DOB:          time.Now(),
		AddressDetails: models.Address{
			AddressLine1: "1351 Arbor Bluff Cir",
			City:         "Ballwin",
			State:        "Missouri",
			Zipcode:      "63021",
		},
		Gender: models.Male,
		Phone:  "636-779-6779",
		Email:  "mbnbala@gmail.com",
	}

	jsonPolicy, _ := json.Marshal(policyHolder)

	fmt.Println("Json Policy Details %n", string(jsonPolicy))
	policyRepo = policyHolder
	result, _ := policyRepo.AddPolicyDetails()
	print(result)
}
func UnmarshalJsonFile() {
	//read json file
	file, err := os.Open("location.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	//when the jso.n is an array of objects and we want to unmarshal it into a slice of maps
	var result []map[string]interface{}
	json.Unmarshal(byteValue, &result)

	for i, obj := range result {
		fmt.Printf("Object %d:\n", i+1)
		for key, value := range obj {
			fmt.Printf("%s : %v\n", key, value)
		}
	}
}

func ConvertJsonToModel(jsonString string) (*models.Location, error) {
	var location models.Location
	err := json.Unmarshal([]byte(jsonString), &location)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func ConvertJsons() {

	//unmarshalling json to struct
	jsonLocation := `{"AddressLine1":"1351 Arbor Bluff Cir","City":"Ballwin","State":"Missouri"}`
	location, err := ConvertJsonToModel(jsonLocation)
	if err != nil {
		println(err.Error())
	}
	for key, value := range utility.StructToMapLocation(location) {
		fmt.Printf("%s : %v\n", key, value)
	}
}

func ReadMemberDetailsfromCSV() {

	file, err := os.Open("team.csv")

	if err != nil {
		fmt.Println("Error Opening CSV File", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	for i, record := range records {
		fmt.Printf("Record %d:\n", i+1)
		for j, field := range record {
			fmt.Printf("Field %d: %s\n", j+1, field)

		}

	}

}

func writeCSVFile() {

	var vehicleRepo interfaces.IVehicleRepo = nil

	//create 10 vehicle instances and save them to csv file
	for i := 1; i <= 10; i++ {
		vehicle := &models.Vehicle{
			LicensePlateNo:     "KA-01-123" + fmt.Sprintf("%02d", i),
			Maker:              "Toyota",
			DateOfRegistration: time.Now(),
			VIN:                "CH123456789",
			FuelType:           models.Petrol,
			EngineNo:           "EN123456789",
			Color:              "Red",
		}

		vehicleRepo = vehicle
		_, err := vehicleRepo.Save()
		if err != nil {
			log.Fatalf("Error saving vehicle: %v", err)
		}

	}

	log.Println("10 vehicles saved successfully.")
	vehicles, _ := vehicleRepo.GetAll()
	fmt.Println("All vehicles:")
	headers := []string{"LicensePlateNo", "Maker", "DateOfRegistration", "VIN", "FuelType", "EngineNo", "Color"}
	fileName := "vehicles.csv"
	for _, v := range vehicles {
		fmt.Println("Vehicle:", v)
		for key, value := range utility.StructToMapVehicle(v) {

			fmt.Printf("%s : %v\n", key, value)
		}
	}

	success, err := vehicles[0].SaveToFile(fileName, headers, vehicles)
	if err != nil {
		log.Fatalf("Error saving vehicles to file: %v", err)
	}
	if success {
		log.Println("Vehicles saved to file successfully.")
	}

}

func Recursive() {

	rootClaim := &models.Claim{
		ID:     1,
		Amount: rand.Intn(200) + 100,
	}

	for i := 0; i < 5; i++ {
		claim := &models.Claim{
			ID:     uint(i + 2),
			Amount: rand.Intn(200) + 100,
		}
		rootClaim.Claims = append(rootClaim.Claims, claim)
	}

	total := rootClaim.TotalClaimsAmount()
	println("Total Claims Amount:", total) // Output: Total Claims Amount: 1000
}
