package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/varuuntiwari/tg-notifier/notifier"
)

func main() {
	godotenv.Load()

	userID, err := strconv.ParseInt(os.Getenv("USER_ID"), 10, 64)
	if err != nil {
		fmt.Printf("%s is not a valid user id\n", os.Getenv("USER_ID"))
		return
	}

	notifService := notifier.NewTelegramNotifier(os.Getenv("BOT_TOKEN"), userID)

	// Keep message here
	if len(os.Args) == 2 {
		msg := os.Args[1]
		err = notifService.SendText(msg)
	} else {
		err = errors.New("invalid parameters passed")
	}

	if err != nil {
		fmt.Println("sorry, an unexpected error occurred:", err)
	} else {
		fmt.Printf("message sent successfully to %d\n", userID)
	}
}
