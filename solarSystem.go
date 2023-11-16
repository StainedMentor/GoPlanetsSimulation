package main

type SolarSystem struct {
	name string

	objects []Planet
}

func (system *SolarSystem) SunOnly() {

	for i := 1; i < len(system.objects); i++ {
		system.objects[i].Influence(system.objects[0])
		system.objects[i].TimeStep(time)
	}

}

func (system *SolarSystem) UniverseStep() {
	for p := 0; p < len(system.objects); p++ {

		for i := 0; i < len(system.objects); i++ {
			if i != p {
				system.objects[p].Influence(system.objects[i])

			}
		}
		system.objects[p].TimeStep(time)

	}

}
