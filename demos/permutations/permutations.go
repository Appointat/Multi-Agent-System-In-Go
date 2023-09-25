package main

import (
	"fmt"

	"gitlab.utc.fr/lagruesy/ia04/utils"
)

func main() {
	const n = 4

	// création et affichage de la première permutation
	perm := utils.FirstPermutation(n)
	fmt.Println(perm)

	// compteur des permutations
	count := 1

	// itération et affichage de la permutation suivante
	perm, ok := utils.NextPermutation(perm)
	for ok {
		count++
		fmt.Println(perm)
		perm, ok = utils.NextPermutation(perm)
	}

	// on affiche la valeur de la factorielle...
	fmt.Printf("\n%d!=%d\n", n, count)
}
