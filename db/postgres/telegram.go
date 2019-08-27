package postgres

import "fmt"

func SetChatID(id int64) error {
	_, err := DBX.Exec("INSERT INTO telegram_chats (id) VALUES ($1);", id)
	if err != nil {
		fmt.Println("SetChatID", err)
		return err
	}
	return nil
}

func GetChatID() ([]int64, error) {
	ids := make([]int64, 0)
	err := DBX.Select(&ids, "SELECT id FROM telegram_chats;")
	if err != nil {
		fmt.Println("GetChatID", err)
		return nil, err
	}
	return ids, nil
}
