package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("You've made it past the guards, now what?")
	} else {
		isHeistOn = false
		fmt.Println("Ya'll didnt fool the guards")
	}

	fmt.Println("heist is on:", isHeistOn)
	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Vaults open! Secure the bag!")
	} else if isHeistOn && openedVault < 70 {
		fmt.Println("The vault was so hard to open that when it did your soul was vaccuumed out and you died bro. Sorry.")
		isHeistOn = false
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely{
		case 0: 
			isHeistOn = false
			fmt.Println("The heist was not successful lil bruh.")
		case
		}
	}

}
