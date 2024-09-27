package notifier

type TelegramNotifier struct {
	BotToken string
	UserID   int64
}

func NewTelegramNotifier(BotToken string, UserID int64) *TelegramNotifier {
	return &TelegramNotifier{BotToken: BotToken, UserID: UserID}
}

func (config TelegramNotifier) SendText(msg string) (status error) {
	status = sendTelegramText(config, msg)
	return
}
