package book

import (
	"context"
	"log"

	"graphql/infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBookByName(ctx context.Context, name string) (result interface{}) {
	var book Book
	data := infrastructure.Mongodb.Collection("booklist").FindOne(ctx, bson.M{"name": name})
	data.Decode(&book)
	return book
}

func GetBookList(ctx context.Context, limit int) (result interface{}) {
	var book Book
	var books []Book

	option := options.Find().SetLimit(int64(limit))

	cur, err := infrastructure.Mongodb.Collection("booklist").Find(ctx, bson.M{}, option)
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	for cur.Next(ctx) {
		cur.Decode(&book)
		books = append(books, book)
	}
	return books
}

func InsertBook(ctx context.Context, book Book) error {
	_, err := infrastructure.Mongodb.Collection("booklist").InsertOne(ctx, book)
	return err
}

func UpdateBook(ctx context.Context, book Book) error {
	filter := bson.M{"name": book.Name}
	update := bson.M{"$set": book}
	upsertBool := true
	updateOption := options.UpdateOptions{
		Upsert: &upsertBool,
	}
	_, err := infrastructure.Mongodb.Collection("booklist").UpdateOne(ctx, filter, update, &updateOption)
	return err
}

func DeleteBook(ctx context.Context, name string) error {
	_, err := infrastructure.Mongodb.Collection("booklist").DeleteOne(ctx, bson.M{"name": name})
	return err
}
