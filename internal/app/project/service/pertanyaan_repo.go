package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	log "github.com/sirupsen/logrus"

	"github.com/odoo-backend/internal/app/project"
)

func (r *MongoRepo) FetchShowPertanyaanJawaban(pertanyaan string) (string, error) {

	var pertanyaanjawaban *project.PertanyaanJawaban

	search := ".*" + pertanyaan + ".*"
	searchFilter := []bson.D{
		bson.D{{"pertanyaan", primitive.Regex{Pattern: search, Options: "i"}}},
	}

	collection := r.Collection(CollectionQNA)
	filter := bson.M{"$or": searchFilter}

	err := collection.FindOne(context.TODO(), filter).Decode(&pertanyaanjawaban)
	if err != nil {
		return "", errors.New("ppg")
	}

	fmt.Printf("Found a single document: %+v\n", pertanyaanjawaban.ID)

	return pertanyaanjawaban.Jawaban, nil
}

func (r *MongoRepo) FetchStorePertanyaanJawaban(data *project.PertanyaanJawaban) (*project.PertanyaanJawaban, error) {
	collection := r.Collection(CollectionQNA)

	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	data.ID = insertResult.InsertedID

	return data, nil
}
