package main

import (
	"./src/Model"
)

func main() {
	var cfg email.EmailConfig
	cfg.Username = ""
	cfg.Password = "D3m0st3n3s"
	cfg.ServerHost = "smtp-mail.outlook.com"
	cfg.ServerPort = "587"
	cfg.SenderAddr = "test.miderp@hotmail.com"

}
