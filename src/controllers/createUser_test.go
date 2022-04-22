package controllers_test

import (
	"apiposterr/src/controllers"
	"apiposterr/src/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserNames(t *testing.T) {
	testCases := []struct {
		NameCase    string
		User        structs.User
		ExpectError bool
	}{
		{
			"only alphanumerics e less than 14 characters",
			structs.User{
				Username: "luci",
			},
			false,
		},
		{
			"just 14 characters and only alphanumerics",
			structs.User{
				Username: "uuuuuuuuuuuuuu",
			},
			false,
		},
		{
			"not only alphanumerics",
			structs.User{
				Username: "luci@",
			},
			true,
		},
		{
			"more than 14 characters",
			structs.User{
				Username: "luuuuuuuuuuuuuuci",
			},
			true,
		},
		{
			"more than 14 characters and not only alphanumerics",
			structs.User{
				Username: "luuuuuuuuuuuuuuci@@@@@",
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.NameCase, func(t *testing.T) {
			err := controllers.ValidadeUsername(tc.User)
			if tc.ExpectError {
				assert.NotEmpty(t, err, "should not be empty")
			} else {
				assert.Empty(t, err, "should be empty")
			}
		})
	}
}
