package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

//Tutorial: https://medium.com/wireless-registry-engineering/gonum-tutorial-linear-algebra-in-go-21ef136fc2d7
func main() {
	// u := mat.NewVecDense(3, []float64{1, 2, 3})
	// v := mat.NewVecDense(3, []float64{4, 5, 6})
	// fmt.Println(u)
	// // Find 1st element of u
	// a := u.AtVec(0)
	// fmt.Println(a)
	// // Equivalent getter method for vectors
	// // The At method may be used on anything that satisfies the Matrix // interface.
	// a = u.At(1, 0)
	// fmt.Println(a)
	// // Overwrite 1st element of u with 33.2
	// u.SetVec(1, 33.2)
	// fmt.Println(u)

	// w := mat.NewVecDense(3, nil)
	// w.AddVec(u, v)
	// matPrint(w)
	// start := time.Now()
	// v := []float64{.1, .2, .3, .4, .5, .6, .7, .8, .9, .10, .11, .12}
	// A := mat.NewDense(3, 4, v)
	// matPrint(A)
	// fmt.Println(A.At(0, 2))
	// fmt.Println(A.At(1, 3))
	// fmt.Println(A.At(2, 1))

	// B := mat.NewDense(3, 4, nil)
	// B.Add(A, A)
	// println("B:")
	// matPrint(B)

	// C := mat.NewDense(3, 4, nil)
	// C.Sub(A, B)
	// println("A - B:")
	// matPrint(C)

	// C.Scale(3.5, B)
	// println("3.5 * B:")
	// matPrint(C)

	// D := mat.NewDense(3, 3, nil)
	// D.Product(A, B.T())
	// println("A * B'")
	// matPrint(D)

	// D.Product(D, A, B.T(), D)
	// println("D * A * B' * D")
	// matPrint(D)
	// end := time.Now()
	// elapsed := end.Sub(start)
	// println(elapsed)

	// fmt.Println("Raw", D.RawMatrix())

	a := mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})
	b := mat.NewDense(1, 3, []float64{.5, 1, 0})
	var c mat.Dense
	c.Mul(a, b)
	matPrint(&c)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
