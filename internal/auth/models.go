package auth

type AuthConfig struct {
	SecretKey []byte
	Issuer    string
}

type UserLogin struct {
	Username string
	Password string
}

type UserDBEntry struct {
	Id           string
	Username     string
	PasswordHash string
	Role         []string
	CreatedAt    string
}
