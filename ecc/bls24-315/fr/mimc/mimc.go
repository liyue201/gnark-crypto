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

package mimc

import (
	"hash"

	"github.com/liyue201/gnark-crypto/ecc/bls24-315/fr"
	"golang.org/x/crypto/sha3"
	"math/big"
	"sync"
)

const (
	mimcNbRounds = 91
	seed         = "seed"   // seed to derive the constants
	BlockSize    = fr.Bytes // BlockSize size that mimc consumes
)

// Params constants for the mimc hash function
var (
	mimcConstants [mimcNbRounds]fr.Element
	once          sync.Once
)

// digest represents the partial evaluation of the checksum
// along with the params of the mimc function
type digest struct {
	h    fr.Element
	data []byte // data to hash
}

// GetConstants exposed to be used in gnark
func GetConstants() []big.Int {
	once.Do(initConstants) // init constants
	res := make([]big.Int, mimcNbRounds)
	for i := 0; i < mimcNbRounds; i++ {
		mimcConstants[i].ToBigIntRegular(&res[i])
	}
	return res
}

// NewMiMC returns a MiMCImpl object, pure-go reference implementation
func NewMiMC() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

// Reset resets the Hash to its initial state.
func (d *digest) Reset() {
	d.data = nil
	d.h = fr.Element{0, 0, 0, 0}
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (d *digest) Sum(b []byte) []byte {
	buffer := d.checksum()
	d.data = nil // flush the data already hashed
	hash := buffer.Bytes()
	b = append(b, hash[:]...)
	return b
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (d *digest) Size() int {
	return BlockSize
}

// BlockSize returns the number of bytes Sum will return.
func (d *digest) BlockSize() int {
	return BlockSize
}

// Write (via the embedded io.Writer interface) adds more data to the running hash.
// It never returns an error.
func (d *digest) Write(p []byte) (n int, err error) {
	n = len(p)
	d.data = append(d.data, p...)
	return
}

// Hash hash using Miyaguchi–Preneel:
// https://en.wikipedia.org/wiki/One-way_compression_function
// The XOR operation is replaced by field addition, data is in Montgomery form
func (d *digest) checksum() fr.Element {

	var buffer [BlockSize]byte
	var x fr.Element

	// if data size is not multiple of BlockSizes we padd:
	// .. || 0xaf8 -> .. || 0x0000...0af8
	if len(d.data)%BlockSize != 0 {
		q := len(d.data) / BlockSize
		r := len(d.data) % BlockSize
		sliceq := make([]byte, q*BlockSize)
		copy(sliceq, d.data)
		slicer := make([]byte, r)
		copy(slicer, d.data[q*BlockSize:])
		sliceremainder := make([]byte, BlockSize-r)
		d.data = append(sliceq, sliceremainder...)
		d.data = append(d.data, slicer...)
	}

	if len(d.data) == 0 {
		d.data = make([]byte, 32)
	}

	nbChunks := len(d.data) / BlockSize

	for i := 0; i < nbChunks; i++ {
		copy(buffer[:], d.data[i*BlockSize:(i+1)*BlockSize])
		x.SetBytes(buffer[:])
		r := d.encrypt(x)
		d.h.Add(&r, &d.h).Add(&d.h, &x)
	}

	return d.h
}

// plain execution of a mimc run
// m: message
// k: encryption key
func (d *digest) encrypt(m fr.Element) fr.Element {
	once.Do(initConstants) // init constants

	for i := 0; i < mimcNbRounds; i++ {
		// m = (m+k+c)^5
		var tmp fr.Element
		tmp.Add(&m, &d.h).Add(&tmp, &mimcConstants[i])
		m.Square(&tmp).
			Square(&m).
			Mul(&m, &tmp)
	}
	m.Add(&m, &d.h)
	return m
}

// Sum computes the mimc hash of msg from seed
func Sum(msg []byte) ([]byte, error) {
	var d digest
	if _, err := d.Write(msg); err != nil {
		return nil, err
	}
	h := d.checksum()
	bytes := h.Bytes()
	return bytes[:], nil
}

func initConstants() {
	bseed := ([]byte)(seed)

	hash := sha3.NewLegacyKeccak256()
	_, _ = hash.Write(bseed)
	rnd := hash.Sum(nil) // pre hash before use
	hash.Reset()
	_, _ = hash.Write(rnd)

	for i := 0; i < mimcNbRounds; i++ {
		rnd = hash.Sum(nil)
		mimcConstants[i].SetBytes(rnd)
		hash.Reset()
		_, _ = hash.Write(rnd)
	}
}
