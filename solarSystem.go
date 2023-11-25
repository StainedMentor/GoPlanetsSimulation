package main

type SolarSystem struct {
	name string

	objects []Planet
}

// Steps with only the sun influencing the planets
func (system *SolarSystem) SunOnly() {

	for i := 1; i < len(system.objects); i++ {
		system.objects[i].Influence(system.objects[0])
		system.objects[i].TimeStep(time)
	}

}

// Step with each planet influencing every other planet
func (system *SolarSystem) UniverseStep() {
	for p := 0; p < len(system.objects); p++ {

		for i := 0; i < len(system.objects); i++ {
			if i != p {
				system.objects[p].Influence(system.objects[i])

			}
		}

	}
	for p := 0; p < len(system.objects); p++ {
		system.objects[p].TimeStep(time)
	}

}
