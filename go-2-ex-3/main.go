package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Datenmodell ",
		117: "Datenbank",
		346: "Cloud",
	}
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])
	delete(modules, 104)
	modules[104] = "DM V2" // replace the removed value
	// now replace a value with antother
	modules[117] = "DB V2" //?

	// TODO: delete one
	// TODO: add one
	// TODO: replace one
	fmt.Println(modules)
}
