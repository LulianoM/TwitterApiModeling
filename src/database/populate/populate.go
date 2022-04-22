package populate

import "gorm.io/gorm"

var DB *gorm.DB

func Populate() {

	err := populateUsers()
	if err != nil {
		panic("Could not connect with the database!")
	}

}
