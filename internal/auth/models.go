package auth

type AuthConfig struct {
	SecretKey []byte
	Issuer    string
}

type User struct {
	Id          string
	Email       string
	Roles       []string
	Permissions []string
	Org_id      int
}
