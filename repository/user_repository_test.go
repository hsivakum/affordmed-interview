package repository

import (
	"affordmed/models"
	"affordmed/repository/test_helper"
	"testing"
)

func TestInsertUser_WhenNewUser(t *testing.T) {
	db := test_helper.GetDB()
	repository := NewUserRepository(db)
	err := repository.InsertUser(models.User{
		FirstName: "Dummy",
		LastName:  "User",
		Email:     "dummyuser@gmail.com",
		Password:  "randomPassword",
	})
	if err != nil {
		t.Fatalf("expectation not met %v", err)
		return
	}
	t.Log("expectation met")
	test_helper.ClearDB(db, t)
}
