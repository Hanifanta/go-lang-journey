package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	Golang_Database_MySQL "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(Golang_Database_MySQL.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Testing Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(Golang_Database_MySQL.GetConnection())

	comment, err := commentRepository.FindbyId(context.Background(), 34)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindALL(t *testing.T) {
	commentRepository := NewCommentRepository(Golang_Database_MySQL.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
