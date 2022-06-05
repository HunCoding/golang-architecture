package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type UserRepository interface {
	InsertUser(user_id string, name string) *User
	GetUserByID(user_id string) *User
}

type mysqlRepository struct{}
type mongodbRepository struct{}

func (m *mongodbRepository) InsertUser(user_id string, name string) *User {
	panic("not implemented") // TODO: Implement
}

func (m *mongodbRepository) GetUserByID(user_id string) *User {
	panic("not implemented") // TODO: Implement
}

type userServices struct {
	repo UserRepository
}

func (m *mysqlRepository) InsertUser(user_id string, name string) *User {
	user := User{user_id, name}
	inMemoUsers[user_id] = user
	return &user
}

func (m *mysqlRepository) GetUserByID(user_id string) *User {
	user, ok := inMemoUsers[user_id]
	if !ok {
		return nil
	}

	return &user
}

var (
	inMemoUsers map[string]User
)

func init() {
	inMemoUsers = make(map[string]User)
}

type User struct {
	id   string
	name string
}

func main() {
	mongodbRepo := &mongodbRepository{}

	services := &userServices{mongodbRepo}

	name := "HunCoding"
	user_inserted := services.InsertUserServices(name)
	fmt.Printf("User inserted: %#v \n", user_inserted)

	user_returned := services.GetUserByIDServices(user_inserted.id)
	fmt.Printf("User returned from get: %#v \n", user_returned)
}

// Alto nivel
func (s *userServices) GetUserByIDServices(user_id string) *User {
	return s.repo.GetUserByID(user_id)
}

// Alto nivel
func (s *userServices) InsertUserServices(name string) *User {
	user_id := strconv.Itoa(rand.Intn(100000))

	return s.repo.InsertUser(user_id, name)
}
