package comsoc

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	//we need to consider the situation where alt1 or alt2 is not in the preference list
	rank1 := rank(alt1, prefs)
	rank2 := rank(alt2, prefs)
	//if alt1 is not in the preference list, it is not preferred to alt2
	//if alt2 is not in the preference list, it is preferred to alt1
	//if both are not in the preference list, return false
	if (rank1 == -1) && (rank2 == -1) {
		return false
	}
	if rank1 == -1 {
		return false
	}
	if rank2 == -1 {
		return true
	}
	return rank1 < rank2
}

// return the best alternatives for the given profile in order of preference
func maxCount(count Count) (bestAlts []Alternative) {
	max := 0
	for alt, cnt := range count {
		if cnt > max {
			max = cnt
			bestAlts = []Alternative{alt}
			// TODO: maxCount函数应该根据test函数应返回一个map，但是题目要求却返回一个数组
		} else if cnt == max {
			bestAlts = append(bestAlts, alt)
		}
	}
	return bestAlts

}

// check if the given profile, e.g. that they are all complete and that each alt only appears once per pref
func checkProfile(prefs Profile) error {
	// Check if the profile is complete
	if len(prefs) == 0 {
		return errors.New("empty profile")
	}

	numAlts := len(prefs[0])
	// Check the length of the preference list
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

		// Ensure that each alternative is presented in the preference list
		for _, pref := range prefs {
			for _, alt := range pref {
				if !seen[alt] {
					return fmt.Errorf("alternative %d is missing in preference list %d", alt, i)
				}
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
func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	count = make(Count)
	for _, pref := range p {
		count[pref[0]]++
	}
	return count, nil
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
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

// 这个阈值是每个选民可以选择的最大的候选人数
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
	//the thresholds is the maximum number of alternatives that each voter can choose
	//if the voter chooses more than the threshold, the alternatives beyond the threshold will be ignored
	//if the voter chooses less than the threshold, the alternatives he chooses will be counted

	//check if the length of the thresholds is equal to the length of the profile
	if len(thresholds) != len(p) {
		return nil, fmt.Errorf("the length of the thresholds is not equal to the length of the profile")
	}

	count = make(Count)
	for i, pref := range p {
		for j := 0; j < thresholds[i]; j++ {
			count[pref[j]]++
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

// The tie-breaking method, which returns the best alternative among the given alternatives
// we use the random method to break the tie
func TieBreak(alts []Alternative) (Alternative, error) {
	if len(alts) == 0 {
		return 0, errors.New("empty alternatives list")
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Randomly select an alternative from the list
	selected := alts[rand.Intn(len(alts))]

	return selected, nil
}

//这里要我们创建多种TieBreak函数并实现工厂模式，我省略掉了

// SWFFactory is a function that returns a SWF
func SWFFactory(SWF func(p Profile) (Count, error), TieBreak func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {
	return func(p Profile) ([]Alternative, error) {
		count, err := SWF(p)
		if err != nil {
			return nil, err
		}

		bestAlts := maxCount(count)
		if len(bestAlts) == 1 {
			return bestAlts, nil
		}

		if len(bestAlts) > 1 {
			bestAlt, err := TieBreak(bestAlts)
			if err != nil {
				return nil, err
			}
			return []Alternative{bestAlt}, nil
		}

		//if there is no best alternative, return error
		return nil, fmt.Errorf("there is no best alternative")
	}
}

// SCFFactory is a function that returns a SCF

func SCFFactory(scf func(p Profile) ([]Alternative, error), TieBreak func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {
	return func(p Profile) (Alternative, error) {
		bestAlts, err := scf(p)
		if err != nil {
			return 0, err
		}
		if len(bestAlts) == 1 {
			return bestAlts[0], nil
		}
		if len(bestAlts) > 1 {
			bestAlt, err := TieBreak(bestAlts)
			if err != nil {
				return 0, err
			}
			return bestAlt, nil
		}
		return 0, fmt.Errorf("there is no best alternative")
	}
}

// To find the Condorcet winner, we need to compare each pair of alternatives.
// the return value is void or the Condorcet winner
func CondorcetWinner(p Profile) (bestAlt []Alternative, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	numAlts := len(p[0])
	//the number of wins of each alternative
	wins := make([]int, numAlts)
	for i := 0; i < numAlts; i++ {
		for j := 0; j < numAlts; j++ {
			if i != j {
				winsForI := 0
				for _, voterPref := range p {
					if isPref(Alternative(i+1), Alternative(j+1), voterPref) {
						winsForI++
					}
				}
				if winsForI > len(p)/2 {
					wins[i]++
				}
			}
		}
	}
	fmt.Println(wins)
	//find the Condorcet winner
	for i, win := range wins {
		if win == numAlts-1 {
			bestAlt = append(bestAlt, Alternative(i+1))
		}
	}

	return bestAlt, nil
}

// The Copeland method
// win +1, lose -1, tie 0
func CopelandSWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return nil, err
	}

	numAlts := len(p[0])
	count = make(Count)
	for i := 0; i < numAlts; i++ {
		for j := 0; j < numAlts; j++ {
			if i != j {
				winsForI := 0
				for _, voterPref := range p {
					if isPref(Alternative(i), Alternative(j), voterPref) {
						winsForI++
					}
				}
				if winsForI > len(p)/2 {
					count[Alternative(i)]++
				} else if winsForI < len(p)/2 {
					count[Alternative(i)]--
				}
			}
		}
	}

	return count, nil
}

func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
}
