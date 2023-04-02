package main

type Config struct {
	token         string
	usersJsonPath string
	port          string
}

var localConfig = Config{
	token:         "6107668057:AAHdg0mkTFYYQIqJoeC4t7hjV-D9XqMPGKI",
	usersJsonPath: "users.json",
	port:          "9095",
}

var dockerConfig = Config{
	token:         "6107668057:AAHdg0mkTFYYQIqJoeC4t7hjV-D9XqMPGKI",
	usersJsonPath: "/etc/tgbot/users.json",
	port:          "9095",
}

var config = dockerConfig
