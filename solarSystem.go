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
