package models

import (
	"errors"
	"log"
	"time"

	"sample-go-app/database"
	"sample-go-app/forms"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string `json:"ID,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
	db := database.GetClient()
	id := uuid.NewV4()
	user := User{
		ID:        id.String(),
		Name:      userPayload.Name,
		BirthDay:  userPayload.BirthDay,
		Gender:    userPayload.Gender,
		PhotoURL:  userPayload.PhotoURL,
		Time:      time.Now().UnixNano(),
		Active:    true,
		UpdatedAt: time.Now().UnixNano(),
	}
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		errors.New("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Users"),
	}
	if _, err := db.PutItem(params); err != nil {
		return nil, errors.New("Error when try to save data to database")
	}
	return &user, nil
}

func (h User) GetByID(id string) (*User, error) {
	db := database.GetClient()
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
		TableName:      aws.String("Users"),
		ConsistentRead: aws.Bool(true),
	}
	resp, err := db.GetItem(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var user *User
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (h User) GetAll() ([]User, error) {
	db := database.GetClient()
	params := &dynamodb.ScanInput{
		TableName: aws.String("Users"),
	}
	resp, err := db.Scan(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var users []User
	if err := dynamodbattribute.UnmarshalListOfMaps(resp.Items, &users); err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}
