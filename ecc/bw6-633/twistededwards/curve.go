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

package twistededwards

import (
	"math/big"
	"sync"

	"github.com/liyue201/gnark-crypto/ecc/bw6-633/fr"
)

// CurveParams curve parameters: ax^2 + y^2 = 1 + d*x^2*y^2
type CurveParams struct {
	A, D     fr.Element
	Cofactor fr.Element
	Order    big.Int
	Base     PointAffine
}

// GetEdwardsCurve returns the twisted Edwards curve on bw6-633/Fr
func GetEdwardsCurve() CurveParams {
	initOnce.Do(initCurveParams)
	// copy to keep Order private
	var res CurveParams

	res.A.Set(&curveParams.A)
	res.D.Set(&curveParams.D)
	res.Cofactor.Set(&curveParams.Cofactor)
	res.Order.Set(&curveParams.Order)
	res.Base.Set(&curveParams.Base)

	return res
}

var (
	initOnce    sync.Once
	curveParams CurveParams
)

func initCurveParams() {
	curveParams.A.SetString("-1")
	curveParams.D.SetString("37248940285811842784899494310834635440994424264352085037441815381151934266434102922992043546621")
	curveParams.Cofactor.SetString("8")
	curveParams.Order.SetString("4963142838689179791878211236301121218116687802119716497817028544854034649070444389864454748079", 10)

	curveParams.Base.X.SetString("37635937024655419978837220647164498012335808680404874556501960268316961933409049243153117555100")
	curveParams.Base.Y.SetString("23823085625708063001015413934245381846960101450148849601038571303382730455875805408244170280142")
}

// mulByA multiplies fr.Element by curveParams.A
func mulByA(x *fr.Element) {
	x.Neg(x)
}
