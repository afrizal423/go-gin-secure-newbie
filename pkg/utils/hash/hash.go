package hash

import (
	"log"

	"github.com/alexedwards/argon2id"
)

type Hashing struct {
}

// helper untuk hashpassword menggunakan argon2id
func (h *Hashing) HashPassword(pw string) string {
	hash, err := argon2id.CreateHash(pw, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

// helper untuk verifikasi password menggunakan argon2id
func (h *Hashing) VerifikasiPassword(pw string, hashedPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(pw, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return match, err
}
