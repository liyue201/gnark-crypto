package fptower

import "github.com/liyue201/gnark-crypto/ecc/bw6-761/fp"

func (z *E6) nSquare(n int) {
	for i := 0; i < n; i++ {
		z.CyclotomicSquare(z)
	}
}

func (z *E6) nSquareCompressed(n int) {
	for i := 0; i < n; i++ {
		z.CyclotomicSquareCompressed(z)
	}
}

// Expt set z to x^t in E6 and return z
func (z *E6) Expt(x *E6) *E6 {
	// const tAbsVal uint64 = 9586122913090633729
	// tAbsVal in binary: 1000010100001000110000000000000000000000000000000000000000000001
	// drop the low 46 bits (all 0 except the least significant bit): 100001010000100011 = 136227
	// Shortest addition chains can be found at https://wwwhomes.uni-bielefeld.de/achim/addition_chain.html

	var result, x33 E6

	// a shortest addition chain for 136227
	result.Set(x)
	result.nSquare(5)
	result.Mul(&result, x)
	x33.Set(&result)
	result.nSquare(7)
	result.Mul(&result, &x33)
	result.nSquare(4)
	result.Mul(&result, x)
	result.CyclotomicSquare(&result)
	result.Mul(&result, x)

	// the remaining 46 bits
	result.nSquareCompressed(46)
	result.DecompressKarabina(&result)
	result.Mul(&result, x)

	z.Set(&result)
	return z
}

// Expc2 set z to x^c2 in E6 and return z
// ht, hy = 13, 9
// c1 = ht+hy = 22 (10110)
func (z *E6) Expc2(x *E6) *E6 {

	var result E6

	result.CyclotomicSquare(x)
	result.CyclotomicSquare(&result)
	result.Mul(&result, x)
	result.CyclotomicSquare(&result)
	result.Mul(&result, x)
	result.CyclotomicSquare(&result)

	z.Set(&result)

	return z
}

// Expc1 set z to x^c1 in E6 and return z
// ht, hy = 13, 9
// c1 = ht**2+3*hy**2 = 412 (110011100)
func (z *E6) Expc1(x *E6) *E6 {

	var result E6

	result.CyclotomicSquare(x)
	result.Mul(&result, x)
	result.CyclotomicSquare(&result)
	result.CyclotomicSquare(&result)
	result.CyclotomicSquare(&result)
	result.Mul(&result, x)
	result.CyclotomicSquare(&result)
	result.Mul(&result, x)
	result.CyclotomicSquare(&result)
	result.Mul(&result, x)
	result.CyclotomicSquare(&result)
	result.CyclotomicSquare(&result)

	z.Set(&result)

	return z
}

// MulBy034 multiplication by sparse element (c0,0,0,c3,c4,0)
func (z *E6) MulBy034(c0, c3, c4 *fp.Element) *E6 {

	var a, b, d E3

	a.MulByElement(&z.B0, c0)

	b.Set(&z.B1)
	b.MulBy01(c3, c4)

	c0.Add(c0, c3)
	d.Add(&z.B0, &z.B1)
	d.MulBy01(c0, c4)

	z.B1.Add(&a, &b).Neg(&z.B1).Add(&z.B1, &d)
	z.B0.MulByNonResidue(&b).Add(&z.B0, &a)

	return z
}

// Mul034By034 multiplication of sparse element (c0,0,0,c3,c4,0) by sparse element (d0,0,0,d3,d4,0)
func (z *E6) Mul034By034(d0, d3, d4, c0, c3, c4 *fp.Element) *E6 {
	var tmp, x0, x3, x4, x04, x03, x34 fp.Element
	x0.Mul(c0, d0)
	x3.Mul(c3, d3)
	x4.Mul(c4, d4)
	tmp.Add(c0, c4)
	x04.Add(d0, d4).
		Mul(&x04, &tmp).
		Sub(&x04, &x0).
		Sub(&x04, &x4)
	tmp.Add(c0, c3)
	x03.Add(d0, d3).
		Mul(&x03, &tmp).
		Sub(&x03, &x0).
		Sub(&x03, &x3)
	tmp.Add(c3, c4)
	x34.Add(d3, d4).
		Mul(&x34, &tmp).
		Sub(&x34, &x3).
		Sub(&x34, &x4)

	z.B0.A0.MulByNonResidue(&x4).
		Add(&z.B0.A0, &x0)
	z.B0.A1.Set(&x3)
	z.B0.A2.Set(&x34)
	z.B1.A0.Set(&x03)
	z.B1.A1.Set(&x04)
	z.B1.A2.SetZero()

	return z
}
