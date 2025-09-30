package data

import (
	"fmt"
	"math/rand"
)

func Generate(dataType string) any {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_DATE:
		return generateDate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	}
	return ""
}

func generateName() string {
	nameLen := len(name)
	index := rand.Intn(nameLen)
	return name[index]
}

func generateDate() string {
	streetLen := len(address[SUBTYPE_STREET])
	cityLen := len(address[SUBTYPE_CITY])

	streetIndex := rand.Intn(streetLen)
	cityIndex := rand.Intn(cityLen)
	number := rand.Intn(100)

	return fmt.Sprintf("%s, nomor %d, %s", address[SUBTYPE_STREET][streetIndex], number, address[SUBTYPE_CITY][cityIndex])
}

func generateAddress() string {
	return ""
}

func generatePhone() string {
	return ""
}
