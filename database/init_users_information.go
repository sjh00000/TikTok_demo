package database

import (
	"fmt"
	"tiktok/pjdata"
)

func MapDefault() (map[string]bool, map[string]pjdata.Author) {
	var usersLoginInfo map[string]pjdata.Author
	var usersRegister map[string]bool
	var author []Author
	db.Find(&author)
	fmt.Printf("%v\n", author)
	for i := 0; i <= len(author)-1; i++ {
		authorNow := author[i]
		usersRegister[authorNow.Name] = true
		usersLoginInfo[authorNow.Token] = pjdata.Author(authorNow)
	}
	return usersRegister, usersLoginInfo
}

func AddIdNum() int64 {
	var author Author
	db.Last(&author)
	return author.Id + 1
}
