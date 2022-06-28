package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

// PrivateKey generate the private key which is greate than 1 and less than p
func PrivateKey(p *big.Int) *big.Int {
	key := new(big.Int)
	limit := new(big.Int).Sub(p, big.NewInt(2))
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return key.Rand(seed, limit).Add(key, big.NewInt(2))
}

// PublicKey is calculated from g**a mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair create a key pair with prime p and g
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey creates the secret key used for encryption
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
