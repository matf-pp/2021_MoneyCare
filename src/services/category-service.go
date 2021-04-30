package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/src/db"
	"main/src/models"
)

type CategoryService struct {
	ctx                *context.Context
	categoryCollection *mongo.Collection
}

func NewCategoryService(connection *db.Connection, collectionName string) *CategoryService {
	collection := connection.NewCollection(collectionName)
	categoryService := CategoryService{ctx: &connection.Ctx, categoryCollection: collection}
	return &categoryService
}

func (categoryService *CategoryService) InsertOne(name string) (*mongo.InsertOneResult, error) {
	category := models.Category{Name: name}
	res, err := categoryService.categoryCollection.InsertOne(*categoryService.ctx, category)
	return res, err
}

func (categoryService *CategoryService) FindOne(Cname string) (models.Category, error) {
	var result models.Category
	err := categoryService.categoryCollection.FindOne(*categoryService.ctx, bson.M{"name": Cname}).Decode(&result)
	return result, err
}
