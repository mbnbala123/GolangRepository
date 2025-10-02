package models

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//self refrential structure

type Claim struct {
	ClaimID     uint   `json:"id" gorm:"primaryKey"`
	ClaimAmount int    `json:"amount"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

func CreateMongoDBConnection() (*mongo.Client, error) {

	_ = godotenv.Load(".env")
	uri := os.Getenv("uri")
	// Placeholder for MongoDB connection and insertion logic
	//timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB connection URI
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	client.Database("PolicyDB").Collection("claims")
	return client, nil
}

// implement the interface methods
func (c *Claim) Save() (int64, error) {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return 0, err
	}
	collection := client.Database("PolicyDB").Collection("claims")
	result, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		return 0, err
	}
	id := result.InsertedID.(int64)
	return id, nil
}

func (c *Claim) GetAllClaim() ([]*Claim, error) {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return nil, err
	}

	collection := client.Database("PolicyDB").Collection("claims")
	cursor, err := collection.Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		return nil, err

	}
	defer cursor.Close(context.TODO())
	var claims []*Claim
	for cursor.Next(context.TODO()) {
		var claim Claim
		if err := cursor.Decode(&claim); err != nil {
			return nil, err
		}
		claims = append(claims, &claim)
	}

	return claims, nil

}

func (c *Claim) GetByClaimID(id uint) (*Claim, error) {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return nil, err
	}
	collection := client.Database("PolicyDB").Collection("claims")
	var claim Claim
	err = collection.FindOne(context.TODO(), map[string]interface{}{"ClaimID": id}).Decode(&claim)
	if err != nil {
		return nil, err
	}
	return &claim, nil
}

func (c *Claim) UpdateClaim(claimID uint, claimAmount int) (*Claim, error) {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return nil, err
	}

	collection := client.Database("ClaimDB").Collection("claims")
	_, err = collection.UpdateOne(context.TODO(), map[string]interface{}{"claimid": claimID}, map[string]interface{}{"$set": map[string]interface{}{"claimamount": claimAmount}})
	if err != nil {
		return nil, err
	}

	return c, nil

}
