package example

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	if res != 3 {
		t.Logf("Got: %d but expect: %d", res, 3)
		t.Fail()
	}
}

func TestAddTestTable(t *testing.T) {
	testTable := []struct {
		a, b            int
		expectedOutcome int
	}{
		{
			a:               1,
			b:               2,
			expectedOutcome: 3,
		},
		{
			a:               3,
			b:               4,
			expectedOutcome: 7,
		},
	}

	for _, test := range testTable {
		res := Add(test.a, test.b)
		if res != test.expectedOutcome {
			t.Logf("Got: %d but expect %d", res, test.expectedOutcome)
			t.Fail()
		}
	}
}
