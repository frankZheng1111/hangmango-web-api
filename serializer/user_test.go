package serializer

import (
	"github.com/stretchr/testify/assert"
	db "hangmango-web-api/model"
	"testing"
)

func TestSerializeBaseUsers(t *testing.T) {
	user := new(db.User)
	user.Id = 1
	user.LoginName = "test"
	user.PasswordHash = "testPass"

	userResource := SerializeBaseUsers(1, []*db.User{user})
	assert.Equal(t, 1, userResource.TotalCount)
	assert.Equal(t, uint(1), userResource.Data[0].Id)
	assert.Equal(t, "test", userResource.Data[0].LoginName)
}
