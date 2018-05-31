package serrializer

type BaseUser struct {
	ResourceTimestamp
	Id        uint   `json:"id"`
	LoginName string `json:"login_name"`
}

type BaseUserResource struct {
	BaseResource
	Data []BaseUser `json:"data"`
}
