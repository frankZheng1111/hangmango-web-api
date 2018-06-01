package serializer

import (
	db "hangmango-web-api/model"
	"reflect"
	"time"
)

type BaseUser struct {
	Id        uint      `json:"id"`
	LoginName string    `json:"login_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseUserResource struct {
	BaseResource
	Data []*BaseUser `json:"data"`
}

func SerializeBaseUsers(count int, users []*db.User) *BaseUserResource {
	baseUserResource := new(BaseUserResource)
	baseUserResource.TotalCount = count

	for _, user := range users {
		baseUser := new(BaseUser)
		baseUserType := reflect.TypeOf(baseUser).Elem()
		for i := 0; i < baseUserType.NumField(); i++ {
			field := reflect.ValueOf(baseUser).Elem().Field(i)
			fieldValue := reflect.ValueOf(user).Elem().FieldByName(baseUserType.Field(i).Name)
			field.Set(fieldValue)
		}
		baseUserResource.Data = append(baseUserResource.Data, baseUser)
	}
	return baseUserResource
}
