package populate

import (
	"apiposterr/src/database"
	"apiposterr/src/structs"

	"github.com/bxcodec/faker/v3"
)

func populateUsers() error {
	database.Connect()

	for i := 0; i < 30; i++ {
		user := structs.User{
			Username: faker.FirstName(),
		}

		database.DB.Create(&user)
	}
	return nil
}
