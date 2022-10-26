package main

import (
	handlers "MC2/Handlers"
	"fmt"
)

func main() {
	profile := handlers.MakeProfile("Naruto")
	fmt.Println("Name :", profile.Name)
	fmt.Println("Health :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	PowerUp := handlers.PowerUp(profile, 2)
	fmt.Println("Name :", PowerUp.Name)
	fmt.Println("Health :", PowerUp.Health)
	fmt.Println("Power :", PowerUp.Power)
	fmt.Println("Exp :", PowerUp.Exp)
}
