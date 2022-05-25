package unimes
// Unimes
type UniversalMessage struct {
	ID        int64       //message ID
	Sender    Destination //Создатель сообщения
	Locale    Destination //Место где было получено сообщение
	Recipient Destination //Кем получено сообщение
	Date      int32       //Дата получения
	Content               //Содержание сообщения
}

type Content struct {
	Type string
	Data string
}

type Destination struct {
	ID        int64
	Type      string //тип адрессата(пользователь, канал, группа)
	Service   string //телеграм, фейсбук, файловая система
	FirstName string
	Lastname  string
	Username  string
}

func New() *UniversalMessage {
	return &UniversalMessage{}
}
