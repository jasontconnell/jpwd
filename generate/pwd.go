package generate

import (
	"math/rand/v2"
	"time"
)

type PasswordGenerator struct {
	rand *rand.Rand
}

const (
	LOWERCASE string = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS   string = "0123456789"
)

func NewPasswordGenerator() *PasswordGenerator {
	seeddt := time.Now()
	r := rand.New(rand.NewPCG(0, uint64(seeddt.UnixMicro())))
	pg := new(PasswordGenerator{rand: r})
	return pg
}

func (p *PasswordGenerator) GeneratePassword(length int, lowercase, uppercase, numbers bool, symbols string) string {
	pool := ""
	if lowercase {
		pool += LOWERCASE
	}
	if uppercase {
		pool += UPPERCASE
	}
	if numbers {
		pool += NUMBERS
	}
	if len(symbols) > 0 {
		pool += symbols
	}

	if len(pool) == 0 {
		return ""
	}

	pwd := ""
	for i := 0; i < length; i++ {
		idx := p.rand.Int() % len(pool)
		next := pool[idx]
		pwd += string(next)
	}
	return pwd
}
