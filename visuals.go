package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const MAP_SCALE = 40
const PLANET_SCALE = 4

type Game struct {
	system      *SolarSystem
	pixels      []byte
	trajectory  []byte
	fadeCounter int
}

// Execute on every tick & runs 1 step
func (g *Game) Update() error {
	g.system.UniverseStep()

	// g.system.SunOnly()
	return nil
}

// Executes on every screen update
func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, SCREENWIDTH*SCREENHEIGHT*4)
	}
	if g.trajectory == nil {
		g.trajectory = make([]byte, SCREENWIDTH*SCREENHEIGHT*4)
	}

	g.DrawTrajectory()

	copy(g.pixels, g.trajectory)
	g.DrawPlanets()

	if g.fadeCounter%3 == 0 {
		g.FadeTrajectory()
	}
	g.fadeCounter++

	screen.WritePixels(g.pixels)
}

// Sets layout for full screen for ebitengine
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

// Adds position of every planet to trajectory byte map
func (g *Game) DrawTrajectory() {

	for i := 0; i < len(g.system.objects); i++ {
		x := int(g.system.objects[i].pos[0] * MAP_SCALE / AU)
		y := int(g.system.objects[i].pos[1] * MAP_SCALE / AU)

		pos := y*SCREENWIDTH + x + SCREENHEIGHT*SCREENWIDTH/2 + SCREENWIDTH/2
		if pos > SCREENHEIGHT*SCREENWIDTH {
			continue
		}
		if pos < 0 {
			continue
		}
		g.trajectory[pos*4] = 0xff

	}

}

// Draws planets on the pixel byte map
func (g *Game) DrawPlanets() {
	for i := range g.system.objects {
		x := int(g.system.objects[i].pos[0] * MAP_SCALE / AU)
		y := int(g.system.objects[i].pos[1] * MAP_SCALE / AU)

		for xOffset := 0; xOffset < PLANET_SCALE; xOffset++ {
			for yOffset := 0; yOffset < PLANET_SCALE; yOffset++ {

				pos := (y+yOffset-PLANET_SCALE/2)*SCREENWIDTH + x + xOffset - PLANET_SCALE/2 + SCREENHEIGHT*SCREENWIDTH/2 + SCREENWIDTH/2

				if pos > SCREENHEIGHT*SCREENWIDTH {
					continue
				}
				if pos < 0 {
					continue
				}
				g.pixels[pos*4+1] = 0xff

			}
		}
	}

}

// Dims trajectory
func (g *Game) FadeTrajectory() {
	for i := range g.trajectory {
		if g.trajectory[i] == 0 {
			continue
		}
		g.trajectory[i] = g.trajectory[i] - 1
	}
}
