package comsoc

import "fmt"

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
	// check that each alternative appears only once per voter\
	for _, pref := range prefs {
		seen := make(map[Alternative]bool)
		for _, alt := range pref {
			if seen[alt] {
				return fmt.Errorf("alternative %v appears more than once in the profile", alt)
			}
			seen[alt] = true
		}
	}

	return nil
}

// check if the given profile, e.g. that they are all complete and that each alt only appears once per pref
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	if len(prefs[0]) != len(alts) {
		return fmt.Errorf("the number of alternatives in the profile (%v) is different from the number of alternatives (%v)", len(prefs[0]), len(alts))
	}
	altSet := make(map[Alternative]bool)
	for _, pref := range prefs {
		for _, alt := range pref {
			altSet[alt] = true
		}
	}
	for _, alt := range alts {
		if !altSet[alt] {
			return fmt.Errorf("alternative %v is not in the profile", alt)
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

// The Borda method
func BordaSWF(p Profile) (count Count, err error) {
	
