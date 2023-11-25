package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const SCREENWIDTH = 1800
const SCREENHEIGHT = 1100

const G = 6.6743e-11
const AU = 149.6e6 * 1000

const time float64 = 3600 * 24 // in seconds

const TPS int = 240

// creates and adds planets to the SolarSystem struct
func initPlanets(system *SolarSystem) {
	sun := Planet{"Sun", [3]float64{0, 0, 0}, 1.989e30, 696340, [3]float64{0, 0, 0}}
	mercury := Planet{"Mercury", [3]float64{0.3 * AU, 0, 0}, 3.301e23, 2439, [3]float64{0, 50000, 0}}
	venus := Planet{"Venus", [3]float64{0.7 * AU, 0, 0}, 4.867e24, 6051, [3]float64{0, 35000, 0}}
	earth := Planet{"Earth", [3]float64{AU, 0, 0}, 6.046e24, 6371, [3]float64{0, 30000, 0}}
	mars := Planet{"Mars", [3]float64{1.5 * AU, 0, 0}, 6.417e23, 3389, [3]float64{0, 25000, 0}}
	jupiter := Planet{"Jupiter", [3]float64{2.5 * AU, 0, 0}, 1.899e27, 69911, [3]float64{0, 20000, 0}}
	saturn := Planet{"Saturn", [3]float64{4 * AU, 0, 0}, 5.685e26, 58232, [3]float64{0, 15000, 0}}
	uranus := Planet{"Uranus", [3]float64{6 * AU, 0, 0}, 8.682e25, 25362, [3]float64{0, 12000, 0}}
	neptune := Planet{"Neptune", [3]float64{8 * AU, 0, 0}, 1.024e26, 24622, [3]float64{0, 11000, 0}}
	pluto := Planet{"Pluto", [3]float64{10 * AU, 0, 0}, 1.471e22, 1188, [3]float64{0, 10000, 0}}

	system.objects = append(system.objects, sun)
	system.objects = append(system.objects, mercury)
	system.objects = append(system.objects, venus)
	system.objects = append(system.objects, earth)
	system.objects = append(system.objects, mars)
	system.objects = append(system.objects, jupiter)
	system.objects = append(system.objects, saturn)
	system.objects = append(system.objects, uranus)
	system.objects = append(system.objects, neptune)
	system.objects = append(system.objects, pluto)

}

func main() {
	g := &Game{}

	system := SolarSystem{"Our solar system", []Planet{}}

	initPlanets(&system)
	g.system = &system

	ebiten.SetWindowSize(SCREENWIDTH, SCREENHEIGHT)
	ebiten.SetWindowTitle("Solar System")
	ebiten.SetTPS(TPS)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

// Extra output.csv from first test
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func testCSV(system *SolarSystem) {
	f, err := os.Create("./out.csv")
	check(err)
	defer f.Close()

	for i := 0; i < 400; i++ {
		system.SunOnly()
		//fmt.Print(system.objects[3].pos, "\n")

		f.WriteString(fmt.Sprintf("%f", system.objects[3].pos[0]))
		f.WriteString(",")

		f.WriteString(fmt.Sprintf("%f", system.objects[3].pos[1]))
		f.WriteString("\n")

	}
}
