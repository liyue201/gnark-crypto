// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package polynomial

import (
	"github.com/liyue201/gnark-crypto/ecc/bls24-317/fr"
	"github.com/liyue201/gnark-crypto/utils"
	"strconv"

	"math/big"
	"strings"
)

// Polynomial represented by coefficients bn254 fr field.
type Polynomial []fr.Element

// Degree returns the degree of the polynomial, which is the length of Data.
func (p *Polynomial) Degree() uint64 {
	return uint64(len(*p) - 1)
}

// Eval evaluates p at v
// returns a fr.Element
func (p *Polynomial) Eval(v *fr.Element) fr.Element {

	res := (*p)[len(*p)-1]
	for i := len(*p) - 2; i >= 0; i-- {
		res.Mul(&res, v)
		res.Add(&res, &(*p)[i])
	}

	return res
}

// Clone returns a copy of the polynomial
func (p *Polynomial) Clone() Polynomial {
	_p := make(Polynomial, len(*p))
	copy(_p, *p)
	return _p
}

// Set to another polynomial
func (p *Polynomial) Set(p1 Polynomial) {
	if len(*p) != len(p1) {
		*p = p1.Clone()
		return
	}

	for i := 0; i < len(p1); i++ {
		(*p)[i].Set(&p1[i])
	}
}

// AddConstantInPlace adds a constant to the polynomial, modifying p
func (p *Polynomial) AddConstantInPlace(c *fr.Element) {
	for i := 0; i < len(*p); i++ {
		(*p)[i].Add(&(*p)[i], c)
	}
}

// SubConstantInPlace subs a constant to the polynomial, modifying p
func (p *Polynomial) SubConstantInPlace(c *fr.Element) {
	for i := 0; i < len(*p); i++ {
		(*p)[i].Sub(&(*p)[i], c)
	}
}

// ScaleInPlace multiplies p by v, modifying p
func (p *Polynomial) ScaleInPlace(c *fr.Element) {
	for i := 0; i < len(*p); i++ {
		(*p)[i].Mul(&(*p)[i], c)
	}
}

// Scale multiplies p0 by v, storing the result in p
func (p *Polynomial) Scale(c *fr.Element, p0 Polynomial) {
	if len(*p) != len(p0) {
		*p = make(Polynomial, len(p0))
	}
	for i := 0; i < len(p0); i++ {
		(*p)[i].Mul(c, &p0[i])
	}
}

// Add adds p1 to p2
// This function allocates a new slice unless p == p1 or p == p2
func (p *Polynomial) Add(p1, p2 Polynomial) *Polynomial {

	bigger := p1
	smaller := p2
	if len(bigger) < len(smaller) {
		bigger, smaller = smaller, bigger
	}

	if len(*p) == len(bigger) && (&(*p)[0] == &bigger[0]) {
		for i := 0; i < len(smaller); i++ {
			(*p)[i].Add(&(*p)[i], &smaller[i])
		}
		return p
	}

	if len(*p) == len(smaller) && (&(*p)[0] == &smaller[0]) {
		for i := 0; i < len(smaller); i++ {
			(*p)[i].Add(&(*p)[i], &bigger[i])
		}
		*p = append(*p, bigger[len(smaller):]...)
		return p
	}

	res := make(Polynomial, len(bigger))
	copy(res, bigger)
	for i := 0; i < len(smaller); i++ {
		res[i].Add(&res[i], &smaller[i])
	}
	*p = res
	return p
}

// Equal checks equality between two polynomials
func (p *Polynomial) Equal(p1 Polynomial) bool {
	if (*p == nil) != (p1 == nil) {
		return false
	}

	if len(*p) != len(p1) {
		return false
	}

	for i := range p1 {
		if !(*p)[i].Equal(&p1[i]) {
			return false
		}
	}

	return true
}

func signedBigInt(v *fr.Element) big.Int {
	var i big.Int
	v.ToBigIntRegular(&i)
	var iDouble big.Int
	iDouble.Lsh(&i, 1)
	if iDouble.Cmp(fr.Modulus()) > 0 {
		i.Sub(fr.Modulus(), &i)
		i.Neg(&i)
	}
	return i
}

func (p Polynomial) Text(base int) string {

	var builder strings.Builder

	first := true
	for d := len(p) - 1; d >= 0; d-- {
		if p[d].IsZero() {
			continue
		}

		i := signedBigInt(&p[d])

		initialLen := builder.Len()

		if i.Sign() < 1 {
			i.Neg(&i)
			if first {
				builder.WriteString("-")
			} else {
				builder.WriteString(" - ")
			}
		} else if !first {
			builder.WriteString(" + ")
		}

		first = false

		asInt64 := int64(0)
		if i.IsInt64() {
			asInt64 = i.Int64()
		}

		if asInt64 != 1 || d == 0 {
			builder.WriteString(i.Text(base))
		}

		if builder.Len()-initialLen > 10 {
			builder.WriteString("×")
		}

		if d != 0 {
			builder.WriteString("X")
		}
		if d > 1 {
			builder.WriteString(
				utils.ToSuperscript(strconv.Itoa(d)),
			)
		}

	}

	if first {
		return "0"
	}

	return builder.String()
}
