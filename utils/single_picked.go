package utils

import (
	"math/rand"
)

func IsSPPref(pref []int) bool {
	var ag, ad int
	top := pref[0]

	for i, ac := range pref {
		if i == 0 {
			ag = top
			ad = top
			continue
		}

		if ac > top {
			if ac < ad {
				return false
			} else {
				ad = ac
			}
		} else {
			if ac > ag {
				return false
			} else {
				ag = ac
			}
		}
	}
	return true
}

func GenerateSPPref(m int) (pref []int) {
	var ag, ac, ad int

	pref = make([]int, m)
	top := rand.Intn(m)

	ag = top - 1
	ad = top + 1

	for i := range pref {
		if i == 0 {
			ac = top
			pref[i] = ac
			continue
		}

		if rand.Intn(2) == 0 { // g
			if ag >= 0 {
				ac = ag
				ag--
			} else {
				ac = ad
				ad++
			}
		} else { // d
			if ad < m {
				ac = ad
				ad++
			} else {
				ac = ag
				ag--
			}
		}
		pref[i] = ac
	}
	return
}

func GenerateSPPrefSlice(m, n int) (prefs [][]int) {
	prefs = make([][]int, n)

	for i := range prefs {
		prefs[i] = GenerateSPPref(m)
	}

	return
}
