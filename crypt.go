package crypt

import (
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

func genPrimo(t *int) *big.Int {
	primo := big.NewInt(0)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var n int
	var ns string
	for i := 0; i < *t; i++ {
		n = r.Intn(10)
		if i == *t-1 && n%2 == 0 {
			n++
		}
		ns += strconv.Itoa(n)
	}

	primo.SetString(ns, 0)
	for !primo.ProbablyPrime(20) {
		primo.Add(primo, big.NewInt(2))
	}

	return primo
}

// GetKeys return 3 keys (*big.Int):
// 1 - Mod
// 2 - Public Key
// 3 - Private Key
func GetKeys() (mod, pub, priv *big.Int) {
	t := 155
	p := genPrimo(&t)
	q := genPrimo(&t)
	t = 78
	pub = genPrimo(&t)

	mod = big.NewInt(0)
	mod.Mul(p, q)

	var m big.Int
	p.Sub(p, big.NewInt(1))
	q.Sub(q, big.NewInt(1))
	m.Mul(p, q)

	priv = big.NewInt(0)
	priv.ModInverse(pub, &m)

	return
}
