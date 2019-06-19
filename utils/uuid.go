package utils

import (
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

// GetUniqueID generate and retrn uinique id
func GetUniqueID() string {
	uuid := uuid.NewV4()
	return hex.EncodeToString(uuid.Bytes())
}
