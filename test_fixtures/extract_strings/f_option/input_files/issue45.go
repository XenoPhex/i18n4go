package app

import (
	"errors"
	"fmt"
)

type HelpCategory struct {
	CategoryName string
	CommandList  [][]string
}

func main() {
	translatedString := "translated"

	err := test()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = errors.New(T("This is another string which has been {{.Translated}}.", map[string]interface{}{"Translated": translatedString}))
	if err != nil {
		fmt.Println(err.Error())
	}

	data := HelpCategory{
		CategoryName: "ORG ADMIN:",
		CommandList: [][]string{
			{"quotas", "quota", "set-quota"},
			{"create-quota", "delete-quota", "update-quota"},
			{"share-private-domain", "unshare-private-domain"},
		},
	}

	T(data.CategoryName)
}

func test49() {
	if true {
		return errors.New(T("This is a string which has been {{.Translated}}.", map[string]interface{}{"Translated": translatedString}))
	}
}
