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
		fmt.Println()
	}
}
