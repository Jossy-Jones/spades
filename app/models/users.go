package models

import (
	"app/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

var client = database.CreateClient().Collection("users")
var ctx = context.Background()

//Users payload Format
type Users struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//Creating a new User
func CreateUser(us *Users) bool {
	//Creating Random ID for each new user
	ID := "usd"
	val := fmt.Sprintln(rand.Intn(200))
	user_id := ID + val
	//Creating new user and adding to firestore
	_, err := client.Doc(us.Id).Create(ctx, map[string]interface{}{
		"email":    us.Email,
		"username": us.Username,
		"id":       user_id,
		"password": us.Password,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

//Getting a single user by ID
func ReadUser(key string) []byte {
	data, err := client.Doc(key).Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	m := data.Data()

	//Storing Data Stored Respectively
	id := m["id"].(string)
	email := m["email"].(string)
	username := m["username"].(string)
	password := m["password"].(string)

	payload := &Users{
		Id:       id,
		Email:    email,
		Username: username,
		Password: password,
	}
	bs, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Error converting to JSON")
	}
	return bs
}

//Updating a particular User by ID
func UpdateUser(key string) bool {
	_, err := client.Doc(key).Set(ctx, map[string]interface{}{
		"id":        "Dillidading",
		"firstname": "Carlus",
		"username":  "qwerty54",
		"birthdate": "24/04/1997",
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}

//Deleting User Document by ID
func DeleteUser(key string) bool {
	_, err := client.Doc("Falcone").Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return false
	}
	return true
}
