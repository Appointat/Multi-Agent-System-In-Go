package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Fill : Remplit un slice avec des nombres aléatoires.
func Fill(sl []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range sl {
		sl[i] = rand.Intn(1000) // pour des nombres aléatoires entre 0 et 999
	}
}

// Moyenne : Calcule la moyenne des éléments d'un slice.
func Moyenne(sl []int) float64 {
	var sum int
	for _, v := range sl {
		sum += v
	}
	return float64(sum) / float64(len(sl))
}

// ValeursCentrales : Retourne la ou les valeurs centrales d'un slice.
func ValeursCentrales(sl []int) []int {
	l := len(sl)
    sl_copy := make([]int, l) // copie de sl
    copy(sl_copy, sl)
    mid := l / 2 // indice du milieu
    // sl n'est pas trié, donc on doit le trier
    for i := 0; i < l-1; i++ {
        for j := i + 1; j < l; j++ {
            if sl_copy[i] > sl_copy[j] {
                sl_copy[i], sl_copy[j] = sl_copy[j], sl_copy[i]
            }
        }
    }
    fmt.Println("Slice trié:", sl_copy)
    if l%2 == 0 { // si la taille est paire
        return []int{sl[mid-1], sl[mid]}
    } else {
        return []int{sl[mid]}
    }
}

// Plus1 : Ajoute 1 à chaque élément d'un slice.
func Plus1(sl []int) {
	for i := range sl {
		sl[i]++
	}
}

// TestFunctions : Pour tester les fonctions précédentes avec Plus1.
func TestFunctions(sl []int) {
	Fill(sl)
	Plus1(sl)

	fmt.Println("Slice:", sl)
	fmt.Println("Moyenne:", Moyenne(sl))
	fmt.Println("Valeurs centrales:", ValeursCentrales(sl))
}

// Compte : Affiche les nombres contenus dans un tableau.
func Compte(tab []int) {
	for _, v := range tab {
		fmt.Println(v)
	}
}

// func main() {
// 	s := make([]int, 10) // Crée un slice de taille 10
// 	TestFunctions(s)
// 	Compte(s)
// }
