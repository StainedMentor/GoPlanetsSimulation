package main

const G float64 = 1

func main() {

	sun := Planet{"Sun", [3]float64{0, 0, 0}, 1, 1, [3]float64{0, 0, 0}}

	system := SolarSystem{"Our solar system", []Planet{}}

	system.objects = append(system.objects, sun)

}
