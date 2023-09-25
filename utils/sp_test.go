package utils

import "testing"

func TestIsSPPref(t *testing.T) {
	prefs := [][]int{ //{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {3, 2, 1, 4, 5},
		{3, 4, 5, 2, 1}, {3, 4, 2, 5, 1}}

	for _, pref := range prefs {
		if !IsSPPref(pref) {
			t.Errorf("%v should be single picked", pref)
		}
	}
}

func TestGenerateSPPref(t *testing.T) {
	m := 10
	for i := 0; i < 10; i++ {
		pref := GenerateSPPref(m)
		if !IsSPPref(pref) {
			t.Errorf("%v is not single picked", pref)
		}
		for _, v := range pref {
			if v < 0 || v >= m {
				t.Errorf("bad value %v in %v", v, pref)
			}
		}
	}
}

func TestGenerateSPPrefSlice(t *testing.T) {
	m := 5
	n := 20
	prefs := GenerateSPPrefSlice(m, n)
	t.Errorf("%v", prefs)

	if len(prefs) != n {
		t.Errorf("bad size of %v, should be %v", prefs, n)
	}

	if len(prefs[0]) != m {
		t.Errorf("bad size of %v, should be %v", prefs[0], m)
	}
}
