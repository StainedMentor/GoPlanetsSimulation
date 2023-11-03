package main

const G = 6.6743e-11
const AU = 149.6e6

const time float64 = 3600

func initPlanets(system *SolarSystem) {
	sun := Planet{"Sun", [3]float64{0, 0, 0}, 1.989e30, 696340, [3]float64{0, 0, 0}}
	mercury := Planet{"Mercury", [3]float64{0.3 * AU, 0, 0}, 3.301e23, 2439, [3]float64{0, 50000, 0}}
	venus := Planet{"Venus", [3]float64{0.7 * AU, 0, 0}, 4.867e24, 6051, [3]float64{0, 40000, 0}}
	earth := Planet{"Earth", [3]float64{AU, 0, 0}, 6.046e24, 6371, [3]float64{0, 0, 30000}}
	mars := Planet{"Mars", [3]float64{1.5 * AU, 0, 0}, 6.417e23, 3389, [3]float64{0, 25000, 0}}
	jupiter := Planet{"Jupiter", [3]float64{2.5 * AU, 0, 0}, 1.899e27, 69911, [3]float64{0, 20000, 0}}
	saturn := Planet{"Saturn", [3]float64{4 * AU, 0, 0}, 5.685e26, 58232, [3]float64{0, 15000, 0}}
	uranus := Planet{"Uranus", [3]float64{6 * AU, 0, 0}, 8.682e25, 25362, [3]float64{0, 10000, 0}}
	neptune := Planet{"Neptune", [3]float64{8 * AU, 0, 0}, 1.024e26, 24622, [3]float64{0, 5000, 0}}
	pluto := Planet{"Pluto", [3]float64{10 * AU, 0, 0}, 1.471e22, 1188, [3]float64{0, 1000, 0}}

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

	system := SolarSystem{"Our solar system", []Planet{}}

	initPlanets(&system)

}
