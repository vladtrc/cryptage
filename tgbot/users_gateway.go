package main

import (
	"encoding/json"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ChatId int64 `json:"chat_id"`
}

func GetChatIds() (res []int64, err error) {
	users, err := GetUsers()
	if err != nil {
		return
	}
	for _, user := range users.Users {
		res = append(res, user.ChatId)
	}
	return
}
func GetUsers() (res Users, err error) {
	fileBytes, err := os.ReadFile(config.usersJsonPath)
	if os.IsNotExist(err) {
		err = SaveUsers(Users{})
		return
	}
	if err != nil {
		return
	}
	err = json.Unmarshal(fileBytes, &res)
	return
}
func SaveUsers(res Users) (err error) {
	bytes, _ := json.Marshal(&res)
	err = os.WriteFile(config.usersJsonPath, bytes, os.ModePerm)
	return
}

func AddUserId(userid int64) (err error) {
	users, err := GetUsers()
	if err != nil {
		return
	}
	users.Users = append(users.Users, User{ChatId: userid})
	err = SaveUsers(users)
	return
}
