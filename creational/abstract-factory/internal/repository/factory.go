package repository

import "go.mongodb.org/mongo-driver/mongo"

// RepoFactory : Abstract Factory
type RepoFactory interface {
	UserRepo() UserRepo
	ItemRepo() ItemRepo
}

// mongoFactory - ConcreteFactory
type mongoFactory struct {
	db *mongo.Database
}

func NewRepoFactory(db *mongo.Database) RepoFactory {
	return &mongoFactory{db}
}

func (f *mongoFactory) UserRepo() UserRepo {
	return NewUserRepo(f.db.Collection("users"))
}

func (f *mongoFactory) ItemRepo() ItemRepo {
	return NewItemRepo(f.db.Collection("items"))
}
