package comsoc

import (
	"errors"
	"fmt"
)

/*
Types of basic methods

Les alternatives seront représentées par des entiers.
Les profils de préférences sont telles que si profile est un profil, profile[12] représentera les préférences du votant 12. Les alternatives sont classée de la préférée à la moins préférée :  profile[12][0] represente l'alternative préférée du votant 12.
Enfin, les méthodes de vote renvoient un décompte sous forme d'une map qui associe à chaque alternative un entier : plus cet entier est élevé, plus l'alternative a de points et plus elle est préférée pour le groupe compte tenu de la méthode considérée.
*/
type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int // map associant à chaque alternative un entier

// e.g. of type
// profile[12][0] is the first choice of voter 12
// profile[12][1] is the second choice of voter 12
// Count is a map associating each alternative with an integer
// e.g. count[12] = 3 means that alternative 12 has 3 votes

// return the index of the alt in the prefs, or -1 if not found
func rank(alt Alternative, prefs []Alternative) int {
	for i, a := range prefs {
		if a == alt {
			return i
		}
	}
	return -1
}

// return true if alt1 is preferred to alt2 in the profile
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return rank(alt1, prefs) < rank(alt2, prefs)
}

// return the best alternatives for the given profile in order of preference
func maxCount(count Count) (bestAlts []Alternative) {
	idxBestAltMap := make(map[int][]Alternative)
	maxCount := 0
	for alt, vote := range count {
		if vote > maxCount {
			maxCount = vote
		}
		idxBestAltMap[vote] = append(idxBestAltMap[vote], alt)
	}

	for _, alt := range idxBestAltMap[maxCount] {
		bestAlts = append(bestAlts, Alternative(alt))
	}

	return bestAlts
}

// check if the given profile, e.g. that they are all complete and that each alt only appears once per pref
func checkProfile(prefs Profile) error {
	//the profil is not complete, return error
	if len(prefs) == 0 {
		return errors.New("empty profile")
	}

	numAlts := len(prefs[0])
	//check the length of the preference list
	for i, pref := range prefs {
		if len(pref) != numAlts {
			return fmt.Errorf("preference list %d is incomplete", i)
		}

		seen := make(map[Alternative]bool)
		for _, alt := range pref {
			if seen[alt] {
				return fmt.Errorf("alternative %d appears more than once in preference list %d", alt, i)
			}
			seen[alt] = true
		}

		for alt := 0; alt < numAlts; alt++ {
			if !seen[Alternative(alt)] {
				return fmt.Errorf("alternative %d is missing in preference list %d", alt, i)
			}
		}
	}

	return nil
}

// check if the given profile, e.g. that they are all complete and that each alt only appears once per pref
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	altMap := make(map[Alternative]bool)
	for _, alt := range alts {
		altMap[alt] = true
	}

	for i, pref := range prefs {
		// check the length of the preference list
		if len(pref) != len(alts) {
			return fmt.Errorf("preference list %d is incomplete", i)
		}

		seen := make(map[Alternative]bool)
		for _, alt := range pref {
			if seen[alt] {
				return fmt.Errorf("alternative %d appears more than once in preference list %d", alt, i)
			}
			if !altMap[alt] {
				return fmt.Errorf("alternative %d is not in the alternative list", alt)
			}
			seen[alt] = true
		}
	}

	return nil
}

/*
Process of the vote

We distinguish between Social Welfare Functions (SWF),
which return a count based on a profile, and Social Choice Functions (social choice function, SCF)
which return only the preferred alternatives.
*/
// return the count for the given profile
func SWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	count = make(Count)
	for _, pref := range p {
		for _, alt := range pref {
			count[alt]++
		}
	}
	return count, nil
}

// return the best alternatives for the given profile
func SCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := SWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
}

// The simple Majority method
func MajoritySWF(p Profile) (bestAlts []Alternative, err error) {
	count, err := SWF(p)
	if err != nil {
		return nil, err
	}

	// check if there is a majority
	bestAlts = maxCount(count)
	if len(bestAlts) == 1 && count[bestAlts[0]] > len(p)/2 {
		return bestAlts, nil
	}

	return nil, fmt.Errorf("there is no majority")
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	bestAlts, err = MajoritySWF(p)
	if err != nil {
		return nil, err
	}
	err = checkProfileAlternative(p, bestAlts)
	return bestAlts, err
}

// The Borda method
func BordaSWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	count = make(Count)
	for _, pref := range p {
		for i, alt := range pref {
			count[alt] += len(pref) - i - 1
		}
	}
	return count, nil
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
}

// Approval voting method
func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	count = make(Count)
	for _, pref := range p {
		for _, alt := range pref {
			for _, threshold := range thresholds {
				if rank(alt, pref) <= threshold {
					count[alt]++
				}
			}
		}
	}
	return count, nil
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
}
