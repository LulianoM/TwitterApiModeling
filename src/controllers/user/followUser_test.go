package user_test

import (
	"apiposterr/src/controllers/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateFollowSameUser(t *testing.T) {

	uuid_1 := uuid.New()
	uuid_2 := uuid.New()

	testCases := []struct {
		NameCase    string
		User        uuid.UUID
		UserFollow  uuid.UUID
		ExpectError bool
	}{
		{
			"differents uuids",
			uuid_1,
			uuid_2,
			false,
		},
		{
			"same uuids",
			uuid_1,
			uuid_1,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.NameCase, func(t *testing.T) {
			err := user.ValidadeSameUser(tc.User, tc.UserFollow)
			if tc.ExpectError {
				assert.NotEmpty(t, err, "should not be empty")
			} else {
				assert.Empty(t, err, "should be empty")
			}
		})
	}
}
