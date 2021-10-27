package repository

import (
	"context"
	"fmt"
	"log"
	"rest-poc/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
}

//NewMongoRepository
func NewMongoRepository() ProductRpository {
	return &mongoRepo{}
}

var (
	mongoDb *mongo.Database
)

func init() {
	clientOptions := options.Client().
		ApplyURI("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	mongoDb = client.Database("Order-services")
}

func (ps *mongoRepo) List() ([]*entity.Product, error) {
	var products []*entity.Product
	productCollection := mongoDb.Collection("Products")

	ctx := context.Background()
	cursor, err := productCollection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		return products, err
	}

	for cursor.Next(ctx) {
		product := &entity.Product{}
		err := cursor.Decode(product)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (ps *mongoRepo) Add(p *entity.Product) error {
	productCollection := mongoDb.Collection("Products")
	p.ID = primitive.NewObjectID()
	result, err := productCollection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	fmt.Printf("%v", result.InsertedID)
	return nil
}

func (ps *mongoRepo) Update(p *entity.Product) error {
	productCollection := mongoDb.Collection("Products")
	update := bson.M{
		"$set": bson.M{
			"name":     p.Name,
			"cost":     p.Cost,
			"quantity": p.Quantity,
		},
	}
	result, err := productCollection.UpdateByID(context.Background(), p.ID, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%v", result.ModifiedCount)
	return nil
}

func (ps *mongoRepo) Delete(id string) error {
	productCollection := mongoDb.Collection("Products")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	delete := bson.M{
		"_id": objID,
	}
	result, err := productCollection.DeleteOne(context.Background(), delete)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%v", result.DeletedCount)
	return nil
}
