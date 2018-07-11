/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-07-2018
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"flag"
	"fmt"

	".."
)

func main() {
	var mu = flag.Float64("mu", 1.0, "service rate")
	var lambda = flag.Float64("lambda", 1.0, "arrival rate")

	flag.Parse()

	fmt.Printf("lambda: %g\nmu: %g\n", *lambda, *mu)

	q := queue.NewMM1(*mu, *lambda)
	fmt.Printf("N: %g\n", q.N())
	fmt.Printf("Nq: %g\n", q.Nq())
	fmt.Printf("W: %g\n", q.W())
	fmt.Printf("Wq: %g\n", q.Wq())
}
