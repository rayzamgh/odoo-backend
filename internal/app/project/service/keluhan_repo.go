package service

import (
	"context"
	"fmt"

	"github.com/odoo-assignment/internal/app/project"
	log "github.com/sirupsen/logrus"
)

func (r *MongoRepo) FetchStoreKeluhan(keluhan *project.Keluhan) (*project.Keluhan, error) {
	collection := r.Collection(CollectionKeluhan)

	insertResult, err := collection.InsertOne(context.TODO(), keluhan)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	fmt.Println("Inserted a single keluhan: ", insertResult.InsertedID)

	keluhan.ID = insertResult.InsertedID

	return keluhan, nil
}
