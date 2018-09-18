package utils

import (
	"encoding/hex"
	"log"

	"github.com/satori/go.uuid"
)

// GetUniqueID generate and retrn uinique id
func GetUniqueID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	return hex.EncodeToString(uuid.Bytes())
}
