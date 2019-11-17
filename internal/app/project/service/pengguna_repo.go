package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/odoo-backend/internal/app/project"
	log "github.com/sirupsen/logrus"
)

func (r *MongoRepo) ShowAllPengguna() ([]*project.Pengguna, error) {

	collection := r.Collection(CollectionPengguna)
	filter := bson.M{}

	ppg, err := collection.Find(context.TODO(), filter, nil)
	if err != nil {
		return nil, errors.New("ppg")
	}

	penggunas := make([]*project.Pengguna, 0)
	for ppg.Next(context.TODO()) {
		var elem project.Pengguna
		err := ppg.Decode(&elem)

		if err != nil {
			log.Print(err)
			return nil, errors.New("500")
		}

		penggunas = append(penggunas, &elem)
	}

	fmt.Printf("Found a all document: ")

	return penggunas, nil
}

func (r *MongoRepo) FetchStorePengguna(pengguna *project.Pengguna) (*project.Pengguna, error) {
	collection := r.Collection(CollectionPengguna)

	insertResult, err := collection.InsertOne(context.TODO(), pengguna)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	pengguna.ID = insertResult.InsertedID

	return pengguna, nil
}
