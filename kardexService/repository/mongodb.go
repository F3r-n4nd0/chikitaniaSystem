package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kardexService/kardex"
	"kardexService/models"
	"log"
	"time"
)

type mongoDb struct {
	collection     *mongo.Collection
}

func NewMongoDbClient(url string) kardex.Repository {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("chikitania-db").Collection("kardex")

	return &mongoDb{
		collection: collection,
	}
}

func (m mongoDb) GetKardexById(ctx context.Context, kardexId string) (models.Kardex, error) {
	filter := bson.M{"kardex_id": kardexId}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var result struct {
		KardexId string `json:"kardex_id"`
		ProductId    string `json:"product_id"`
	}
	err := m.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Printf("error while get kardex: %v", err)
	}
	return models.Kardex{
		KardexId:     result.KardexId,
		ProductId:    result.ProductId,
		CreateOn:     0,
		ProductPrice: 0,
	}, nil
}

