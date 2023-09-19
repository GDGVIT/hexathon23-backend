package auth

// JWTSecret is the secret key used to sign the JWT
var JWTSecret string

// InitializeAuth initializes the auth package
func InitializeAuth(jwtSecret string) {
	JWTSecret = jwtSecret
}
