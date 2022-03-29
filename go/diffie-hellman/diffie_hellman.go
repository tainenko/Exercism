package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

// PrivateKey create a private key which is greater than 1 and less than p
func PrivateKey(p *big.Int) *big.Int {
	key := new(big.Int)
	limit := new(big.Int).Sub(p, big.NewInt(2))
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return key.Rand(seed, limit).Add(key, big.NewInt(2))
}

// PublicKey calculates a public key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair creates a key pair
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

// SecretKey creates the secret key used for encryption.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
