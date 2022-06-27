package jwt

type JWT interface {
	GenerateToken(payload Payload) (*string, error)
}
