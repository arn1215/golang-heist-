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

	if eludedGuards >= 30 {
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

	leftSafely := rand.Intn(10)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("The heist was not successful lil bruh. You forgot to have a get-away")
		case 1:
			isHeistOn = false
			fmt.Println("The police came to serve and protect the money. You died. Sad.")
		case 2:
			isHeistOn = false
			fmt.Println("You got stuck under a pile of money. You died. Sad.")
		case 3:
			isHeistOn = false
			fmt.Println("You got stuck in the vault. There was no food and you died. Oops.")
		default:
			fmt.Println("YOU DID IT! You succesfully secured Jeff Bezos' bag! W.")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(100000000)
		fmt.Printf("You got away with %d racks. NIIICE.", amtStolen)
	}

}
