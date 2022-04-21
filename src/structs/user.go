package structs

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Username   string    `json:"username"`
	DataJoined time.Time `json:"data_joined"`
	Followers  []string  `json:"followers"`
	Following  []string  `json:"following"`
	IdAccount  uuid.UUID
}
