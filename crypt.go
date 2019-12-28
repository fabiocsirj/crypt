package crypt

import (
	"math/big"
)

func x(d, k1, k2 *big.Int) *big.Int {
	r := big.NewInt(1)
	if k1.Cmp(big.NewInt(0)) == 0 && k2.Cmp(big.NewInt(0)) == 0 {
		return r
	}
	k1d := big.NewInt(0)
	k1d.Mul(k1, d)
	return r.Sub(k2, k1d)
}

//GetPrivKey ...
func GetPrivKey(a, b big.Int, uv ...*big.Int) *big.Int {
	var d, r big.Int
	d.Div(&a, &b)
	r.Mod(&a, &b)
	var u1, u2, v1, v2 *big.Int
	if len(uv) > 0 {
		u1, u2, v1, v2 = uv[0], uv[1], uv[2], uv[3]
	} else {
		u1, u2, v1, v2 = big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)
	}

	if r.Cmp(big.NewInt(0)) == 0 {
		return v1
	}
	u := x(&d, u1, u2)
	v := x(&d, v1, v2)
	if v.Cmp(big.NewInt(1)) == 0 {
		v.Neg(&d)
		v1.Set(big.NewInt(1))
	}

	return GetPrivKey(b, r, u, u1, v, v1)
}
