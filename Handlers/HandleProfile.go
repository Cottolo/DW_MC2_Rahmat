package handlers

import (
	"MC2/models"
)

func MakeProfile(Name string) models.Profile {

	var Profile models.Profile

	if Name == "Sasuke" {
		Profile.Name = "Sasuke"
	}
	if Name == "Sasuke" {
		Profile.Health = 200
	}
	if Name == "Sasuke" {
		Profile.Power = 100
	}
	if Name == "Sasuke" {
		Profile.Exp = 0
	}

	if Name == "Goku" {
		Profile.Name = "Goku"
	}
	if Name == "Goku" {
		Profile.Health = 400
	}
	if Name == "Goku" {
		Profile.Power = 300
	}
	if Name == "Goku" {
		Profile.Exp = 100
	}

	if Name == "Naruto" {
		Profile.Name = "Naruto"
	}
	if Name == "Naruto" {
		Profile.Health = 150
	}
	if Name == "Naruto" {
		Profile.Power = 200
	}
	if Name == "Naruto" {
		Profile.Exp = 50
	}

	return Profile
}
