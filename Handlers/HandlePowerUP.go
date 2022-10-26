package handlers

import "MC2/models"

func PowerUp(Profile models.Profile, nilai int) models.Profile {

	var PowerUp models.Profile
	PowerUp.Name = Profile.Name
	PowerUp.Health = Profile.Health + (Profile.Health * nilai)
	PowerUp.Power = Profile.Power + (Profile.Power * nilai)
	PowerUp.Exp = Profile.Exp + (Profile.Exp * nilai)

	return PowerUp
}
