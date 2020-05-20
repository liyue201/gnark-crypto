// Copyright 2020 ConsenSys AG
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

// Code generated by gurvy/internal/generators DO NOT EDIT

package bls381

import "github.com/consensys/gurvy/bls381/fp"

// e6 is a degree-three finite field extension of fp2:
// B0 + B1v + B2v^2 where v^3-1,1 is irrep in fp2

type e6 struct {
	B0, B1, B2 e2
}

// Equal returns true if z equals x, fasle otherwise
// TODO can this be deleted?  Should be able to use == operator instead
func (z *e6) Equal(x *e6) bool {
	return z.B0.Equal(&x.B0) && z.B1.Equal(&x.B1) && z.B2.Equal(&x.B2)
}

// SetString sets a e6 elmt from stringf
func (z *e6) SetString(s1, s2, s3, s4, s5, s6 string) *e6 {
	z.B0.SetString(s1, s2)
	z.B1.SetString(s3, s4)
	z.B2.SetString(s5, s6)
	return z
}

// Set Sets a e6 elmt form another e6 elmt
func (z *e6) Set(x *e6) *e6 {
	z.B0 = x.B0
	z.B1 = x.B1
	z.B2 = x.B2
	return z
}

// SetOne sets z to 1 in Montgomery form and returns z
func (z *e6) SetOne() *e6 {
	z.B0.A0.SetOne()
	z.B0.A1.SetZero()
	z.B1.A0.SetZero()
	z.B1.A1.SetZero()
	z.B2.A0.SetZero()
	z.B2.A1.SetZero()
	return z
}

// ToMont converts to Mont form
func (z *e6) ToMont() *e6 {
	z.B0.ToMont()
	z.B1.ToMont()
	z.B2.ToMont()
	return z
}

// FromMont converts from Mont form
func (z *e6) FromMont() *e6 {
	z.B0.FromMont()
	z.B1.FromMont()
	z.B2.FromMont()
	return z
}

// Add adds two elements of e6
func (z *e6) Add(x, y *e6) *e6 {
	z.B0.Add(&x.B0, &y.B0)
	z.B1.Add(&x.B1, &y.B1)
	z.B2.Add(&x.B2, &y.B2)
	return z
}

// Neg negates the e6 number
func (z *e6) Neg(x *e6) *e6 {
	z.B0.Neg(&z.B0)
	z.B1.Neg(&z.B1)
	z.B2.Neg(&z.B2)
	return z
}

// Sub two elements of e6
func (z *e6) Sub(x, y *e6) *e6 {
	z.B0.Sub(&x.B0, &y.B0)
	z.B1.Sub(&x.B1, &y.B1)
	z.B2.Sub(&x.B2, &y.B2)
	return z
}

// MulByGen Multiplies by v, root of X^3-1,1
// TODO deprecate in favor of inlined MulByNonResidue in fp12 package
func (z *e6) MulByGen(x *e6) *e6 {
	var result e6

	result.B1 = x.B0
	result.B2 = x.B1
	{ // begin: inline result.B0.MulByNonResidue(&x.B2)
		var buf e2
		buf.Set(&x.B2)
		result.B0.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(result.B0).A0, &buf.A1)
			(&(result.B0).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(result.B0).A0, &buf.A1)
		result.B0.A0.AddAssign(&buf.A0)
	} // end: inline result.B0.MulByNonResidue(&x.B2)

	z.Set(&result)
	return z
}

// Double doubles an element in e6
func (z *e6) Double(x *e6) *e6 {
	z.B0.Double(&x.B0)
	z.B1.Double(&x.B1)
	z.B2.Double(&x.B2)
	return z
}

// String puts e6 elmt in string form
func (z *e6) String() string {
	return (z.B0.String() + "+(" + z.B1.String() + ")*v+(" + z.B2.String() + ")*v**2")
}

// Mul multiplies two numbers in e6
func (z *e6) Mul(x, y *e6) *e6 {
	// Algorithm 13 from https://eprint.iacr.org/2010/354.pdf
	var rb0, b0, b1, b2, b3, b4 e2
	b0.Mul(&x.B0, &y.B0) // step 1
	b1.Mul(&x.B1, &y.B1) // step 2
	b2.Mul(&x.B2, &y.B2) // step 3
	// step 4
	b3.Add(&x.B1, &x.B2)
	b4.Add(&y.B1, &y.B2)
	rb0.Mul(&b3, &b4).
		SubAssign(&b1).
		SubAssign(&b2)
	{ // begin: inline rb0.MulByNonResidue(&rb0)
		var buf e2
		buf.Set(&rb0)
		rb0.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(rb0).A0, &buf.A1)
			(&(rb0).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(rb0).A0, &buf.A1)
		rb0.A0.AddAssign(&buf.A0)
	} // end: inline rb0.MulByNonResidue(&rb0)
	rb0.AddAssign(&b0)
	// step 5
	b3.Add(&x.B0, &x.B1)
	b4.Add(&y.B0, &y.B1)
	z.B1.Mul(&b3, &b4).
		SubAssign(&b0).
		SubAssign(&b1)
	{ // begin: inline b3.MulByNonResidue(&b2)
		var buf e2
		buf.Set(&b2)
		b3.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(b3).A0, &buf.A1)
			(&(b3).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(b3).A0, &buf.A1)
		b3.A0.AddAssign(&buf.A0)
	} // end: inline b3.MulByNonResidue(&b2)
	z.B1.AddAssign(&b3)
	// step 6
	b3.Add(&x.B0, &x.B2)
	b4.Add(&y.B0, &y.B2)
	z.B2.Mul(&b3, &b4).
		SubAssign(&b0).
		SubAssign(&b2).
		AddAssign(&b1)
	z.B0 = rb0
	return z
}

// MulByE2 multiplies x by an elements of e2
func (z *e6) MulByE2(x *e6, y *e2) *e6 {
	var yCopy e2
	yCopy.Set(y)
	z.B0.Mul(&x.B0, &yCopy)
	z.B1.Mul(&x.B1, &yCopy)
	z.B2.Mul(&x.B2, &yCopy)
	return z
}

// MulByNotv2 multiplies x by y with &y.b2=0
func (z *e6) MulByNotv2(x, y *e6) *e6 {
	// Algorithm 15 from https://eprint.iacr.org/2010/354.pdf
	var rb0, b0, b1, b2, b3 e2
	b0.Mul(&x.B0, &y.B0) // step 1
	b1.Mul(&x.B1, &y.B1) // step 2
	// step 3
	b2.Add(&x.B1, &x.B2)
	rb0.Mul(&b2, &y.B1).
		SubAssign(&b1)
	{ // begin: inline rb0.MulByNonResidue(&rb0)
		var buf e2
		buf.Set(&rb0)
		rb0.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(rb0).A0, &buf.A1)
			(&(rb0).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(rb0).A0, &buf.A1)
		rb0.A0.AddAssign(&buf.A0)
	} // end: inline rb0.MulByNonResidue(&rb0)
	rb0.AddAssign(&b0)
	// step 4
	b2.Add(&x.B0, &x.B1)
	b3.Add(&y.B0, &y.B1)
	z.B1.Mul(&b2, &b3).
		SubAssign(&b0).
		SubAssign(&b1)
	// step 5
	z.B2.Mul(&x.B2, &y.B0).
		AddAssign(&b1)
	z.B0 = rb0
	return z
}

// Square squares a e6
func (z *e6) Square(x *e6) *e6 {
	// Algorithm 16 from https://eprint.iacr.org/2010/354.pdf
	var b0, b1, b2, b3, b4 e2
	b3.Mul(&x.B0, &x.B1).Double(&b3) // step 1
	b4.Square(&x.B2)                 // step 2

	// step 3
	{ // begin: inline b0.MulByNonResidue(&b4)
		var buf e2
		buf.Set(&b4)
		b0.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(b0).A0, &buf.A1)
			(&(b0).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(b0).A0, &buf.A1)
		b0.A0.AddAssign(&buf.A0)
	} // end: inline b0.MulByNonResidue(&b4)
	b0.AddAssign(&b3)
	b1.Sub(&b3, &b4)                                  // step 4
	b2.Square(&x.B0)                                  // step 5
	b3.Sub(&x.B0, &x.B1).AddAssign(&x.B2).Square(&b3) // steps 6 and 8
	b4.Mul(&x.B1, &x.B2).Double(&b4)                  // step 7
	// step 9
	{ // begin: inline z.B0.MulByNonResidue(&b4)
		var buf e2
		buf.Set(&b4)
		z.B0.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(z.B0).A0, &buf.A1)
			(&(z.B0).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(z.B0).A0, &buf.A1)
		z.B0.A0.AddAssign(&buf.A0)
	} // end: inline z.B0.MulByNonResidue(&b4)
	z.B0.AddAssign(&b2)

	// step 10
	z.B2.Add(&b1, &b3).
		AddAssign(&b4).
		SubAssign(&b2)
	z.B1 = b0
	return z
}

// Square2 squares a e6
func (z *e6) Square2(x *e6) *e6 {
	// Karatsuba from Section 4 of https://eprint.iacr.org/2006/471.pdf
	var v0, v1, v2, v01, v02, v12 e2
	v0.Square(&x.B0)
	v1.Square(&x.B1)
	v2.Square(&x.B2)
	v01.Add(&x.B0, &x.B1)
	v01.Square(&v01)
	v02.Add(&x.B0, &x.B2)
	v02.Square(&v02)
	v12.Add(&x.B1, &x.B2)
	v12.Square(&v12)
	z.B0.Sub(&v12, &v1).SubAssign(&v2)
	{ // begin: inline z.B0.MulByNonResidue(&z.B0)
		var buf e2
		buf.Set(&z.B0)
		z.B0.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(z.B0).A0, &buf.A1)
			(&(z.B0).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(z.B0).A0, &buf.A1)
		z.B0.A0.AddAssign(&buf.A0)
	} // end: inline z.B0.MulByNonResidue(&z.B0)
	z.B0.AddAssign(&v0)
	{ // begin: inline z.B1.MulByNonResidue(&v2)
		var buf e2
		buf.Set(&v2)
		z.B1.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(z.B1).A0, &buf.A1)
			(&(z.B1).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(z.B1).A0, &buf.A1)
		z.B1.A0.AddAssign(&buf.A0)
	} // end: inline z.B1.MulByNonResidue(&v2)
	z.B1.AddAssign(&v01).SubAssign(&v0).SubAssign(&v1)
	z.B2.Add(&v02, &v1).SubAssign(&v0).SubAssign(&v2)
	return z
}

// Square3 squares a e6
func (z *e6) Square3(x *e6) *e6 {
	// CH-SQR2 from from Section 4 of https://eprint.iacr.org/2006/471.pdf
	var s0, s1, s2, s3, s4 e2
	s0.Square(&x.B0)
	s1.Mul(&x.B0, &x.B1).Double(&s1)
	s2.Sub(&x.B0, &x.B1).AddAssign(&x.B2).Square(&s2)
	s3.Mul(&x.B1, &x.B2).Double(&s3)
	s4.Square(&x.B2)
	{ // begin: inline z.B0.MulByNonResidue(&s3)
		var buf e2
		buf.Set(&s3)
		z.B0.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(z.B0).A0, &buf.A1)
			(&(z.B0).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(z.B0).A0, &buf.A1)
		z.B0.A0.AddAssign(&buf.A0)
	} // end: inline z.B0.MulByNonResidue(&s3)
	z.B0.AddAssign(&s0)
	{ // begin: inline z.B1.MulByNonResidue(&s4)
		var buf e2
		buf.Set(&s4)
		z.B1.A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(z.B1).A0, &buf.A1)
			(&(z.B1).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(z.B1).A0, &buf.A1)
		z.B1.A0.AddAssign(&buf.A0)
	} // end: inline z.B1.MulByNonResidue(&s4)
	z.B1.AddAssign(&s1)
	z.B2.Add(&s1, &s2).AddAssign(&s3).SubAssign(&s0).SubAssign(&s4)
	return z
}

// Inverse an element in e6
func (z *e6) Inverse(x *e6) *e6 {
	// Algorithm 17 from https://eprint.iacr.org/2010/354.pdf
	// step 9 is wrong in the paper!
	// memalloc
	var t [7]e2
	var c [3]e2
	var buf e2
	t[0].Square(&x.B0)     // step 1
	t[1].Square(&x.B1)     // step 2
	t[2].Square(&x.B2)     // step 3
	t[3].Mul(&x.B0, &x.B1) // step 4
	t[4].Mul(&x.B0, &x.B2) // step 5
	t[5].Mul(&x.B1, &x.B2) // step 6
	// step 7
	{ // begin: inline c[0].MulByNonResidue(&t[5])
		var buf e2
		buf.Set(&t[5])
		c[0].A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(c[0]).A0, &buf.A1)
			(&(c[0]).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(c[0]).A0, &buf.A1)
		c[0].A0.AddAssign(&buf.A0)
	} // end: inline c[0].MulByNonResidue(&t[5])
	c[0].Neg(&c[0]).AddAssign(&t[0])
	// step 8
	{ // begin: inline c[1].MulByNonResidue(&t[2])
		var buf e2
		buf.Set(&t[2])
		c[1].A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(c[1]).A0, &buf.A1)
			(&(c[1]).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(c[1]).A0, &buf.A1)
		c[1].A0.AddAssign(&buf.A0)
	} // end: inline c[1].MulByNonResidue(&t[2])
	c[1].SubAssign(&t[3])
	c[2].Sub(&t[1], &t[4]) // step 9 is wrong in 2010/354!
	// steps 10, 11, 12
	t[6].Mul(&x.B2, &c[1])
	buf.Mul(&x.B1, &c[2])
	t[6].AddAssign(&buf)
	{ // begin: inline t[6].MulByNonResidue(&t[6])
		var buf e2
		buf.Set(&t[6])
		t[6].A1.Add(&buf.A0, &buf.A1)
		{ // begin: inline MulByNonResidue(&(t[6]).A0, &buf.A1)
			(&(t[6]).A0).Neg(&buf.A1)
		} // end: inline MulByNonResidue(&(t[6]).A0, &buf.A1)
		t[6].A0.AddAssign(&buf.A0)
	} // end: inline t[6].MulByNonResidue(&t[6])
	buf.Mul(&x.B0, &c[0])
	t[6].AddAssign(&buf)

	t[6].Inverse(&t[6])    // step 13
	z.B0.Mul(&c[0], &t[6]) // step 14
	z.B1.Mul(&c[1], &t[6]) // step 15
	z.B2.Mul(&c[2], &t[6]) // step 16
	return z
}

// MulByNonResidue multiplies a e2 by (1,1)
func (z *e2) MulByNonResidue(x *e2) *e2 {
	var buf e2
	buf.Set(x)
	z.A1.Add(&buf.A0, &buf.A1)
	{ // begin: inline MulByNonResidue(&(z).A0, &buf.A1)
		(&(z).A0).Neg(&buf.A1)
	} // end: inline MulByNonResidue(&(z).A0, &buf.A1)
	z.A0.AddAssign(&buf.A0)
	return z
}

// MulByNonResidueInv multiplies a e2 by (1,1)^{-1}
func (z *e2) MulByNonResidueInv(x *e2) *e2 {
	// (z).A0 = ((x).A0 + (x).A1)/2
	// (z).A1 = ((x).A1 - (x).A0)/2
	buf := *(x)
	(z).A0.Add(&buf.A0, &buf.A1)
	(z).A1.Sub(&buf.A1, &buf.A0)
	twoInv := fp.Element{
		1730508156817200468,
		9606178027640717313,
		7150789853162776431,
		7936136305760253186,
		15245073033536294050,
		1728177566264616342,
	}
	(z).A0.MulAssign(&twoInv)
	(z).A1.MulAssign(&twoInv)
	return z
}
