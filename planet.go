package main

import (
	"math"
)

type Planet struct {
	name string

	// coordinates
	pos [3]float64

	// size
	m float64
	r float64

	// trajectory
	vel [3]float64
}

func (p *Planet) TimeStep(t float64) {

	for i := 0; i < 3; i++ {
		p.pos[i] += t * p.vel[i]
	}

}

func (p1 Planet) Distance(p2 Planet) float64 {
	return math.Sqrt(math.Pow(p1.pos[0]-p2.pos[0], 2) + math.Pow(p1.pos[1]-p2.pos[1], 2) + math.Pow(p1.pos[2]-p2.pos[2], 2))

}
func (p1 *Planet) Influence(p2 Planet) {
	for i := 0; i < 3; i++ {
		p1.vel[i] += G * p2.m / math.Pow(p2.pos[i]-p1.pos[i], 2)
	}

}
