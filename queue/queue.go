/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 11-07-2018
 * |
 * | File Name:     queue.go
 * +===============================================
 */

package queue

// Queue models steady state queue mean performance parameters
type Queue interface {
	N() float64
	Nq() float64

	W() float64
	Wq() float64
}

// MM1 M/M/1 Queue
type MM1 struct {
	mu     float64
	lambda float64
	rho    float64
}

// NewMM1 creates M/M/1 Queue
func NewMM1(mu, lambda float64) MM1 {
	return MM1{
		mu:     mu,
		lambda: lambda,
		rho:    lambda / mu,
	}
}

// N number of customers in system
func (m MM1) N() float64 {
	return m.rho / (1 - m.rho)
}

// Nq number of customers in queue
func (m MM1) Nq() float64 {
	return m.rho/(1-m.rho) - m.rho
}

// W system waiting time
func (m MM1) W() float64 {
	return 1 / (m.mu - m.lambda)
}

// Wq queue waiting time
func (m MM1) Wq() float64 {
	return m.rho / (m.mu - m.lambda)
}
