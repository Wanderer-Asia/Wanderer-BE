package encrypt

import "golang.org/x/crypto/bcrypt"

type BcryptHash interface {
	Compare(hashed string, input string) error
	Hash(password string) (string, error)
}

func NewBcrypt(cost int) BcryptHash {
	return &bcryptHash{cost: cost}
}

type bcryptHash struct {
	cost int
}

func (enc *bcryptHash) Compare(hashed string, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
}

func (enc *bcryptHash) Hash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), enc.cost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}
