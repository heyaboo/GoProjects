// This mini project is a stimulation of a Bank Heist, probability is determined using randomization variables

/* Algorithm Explanation:
   
    - There are 3 stages to a successful bank heist mission: Getting past the guards, opening the vault and escape plan.
    - Probability from first step to the next becomes lower.
    
*/

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

  //First step is getting past the guards
	if eludedGuards > 51 {
		fmt.Println("I wonder if those really received training, they were so easily fooled! Alright, onto the next task!")
	} else {
		isHeistOn = false
		fmt.Println("Darn it, these banks shouldn't be employing so many guards!")
	}

  //Second step is opening the vault
	openedVault := rand.Intn(100)
	fmt.Println("\nIs the heist still on?", isHeistOn)

	if openedVault > 69 && isHeistOn {
		fmt.Println("\nI will finally be able to buy myself some underwears and deodarant! Alright, let's get outta here!")
	} else if openedVault < 30 && isHeistOn {
		isHeistOn = false
		fmt.Println("\nNo no no~! I am going to continue smelling like the trash can!")
	}

  //3rd step is escaping the bank
	leftSafely := rand.Intn(6)

	if !isHeistOn {
		fmt.Println("\nMission failed, I will now have to think about how to get out of prison alive!")
	} else {

		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("\nMission failed, I will now have to think about how to get out of prison alive!")
		case 1:
			isHeistOn = false
			fmt.Println("\nShould have figured this hall has no blind spot!")
		case 2:
			isHeistOn = false
			fmt.Println("\nCrap! I think I just activated their security alarm!")
		case 3:
			isHeistOn = false
			fmt.Println("\nWait, what? As if guards aren't enough! Now we need to deal with cops too?")
		case 4:
			isHeistOn = false
			fmt.Println("\nMore guards!! I should have not let my guard down, it's over!")
		default:
			fmt.Println("\nYes yes, we made it! Phone our driver now, we are officially rich!")
		}

    //Should we succeed in every step, we get to count how much money we manage to rob!
		if isHeistOn {
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Println("\nWe are $", amtStolen, "richer! GOOD JOB, everybody!")
		}
	}
}
