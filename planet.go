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

// Moves plane by velocity*time
func (p *Planet) TimeStep(t float64) {

	for i := 0; i < 3; i++ {
		p.pos[i] += t * p.vel[i]
	}

}

// Distance between 2 planets
func (p1 Planet) Distance3D(p2 Planet) float64 {
	return math.Sqrt(math.Pow(p1.pos[0]-p2.pos[0], 2) + math.Pow(p1.pos[1]-p2.pos[1], 2) + math.Pow(p1.pos[2]-p2.pos[2], 2))

}

// Modifies the velocity of a planet by the gravitational pull of another planet.
func (p1 *Planet) Influence(p2 Planet) {

	distance := p1.Distance3D(p2)

	force := G * p2.m / math.Pow(distance, 2)

	for i := 0; i < 3; i++ {
		f := force * (p2.pos[i] - p1.pos[i]) / distance
		p1.vel[i] += f * time
	}

}
