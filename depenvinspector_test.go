package depenvinspector


import (
	"testing"

	"fmt"
	"math/rand"
	"time"
)


func TestNew(t *testing.T) {

	inspector := New()
	if nil == inspector {
		t.Errorf("Expected new deployment environment inspector to not be nil: %v", inspector)
	}
}


func TestRegister(t *testing.T) {

	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )


	// Create some specific tests.
	tests := []struct{
		Data []string
		ExpectedSizes []int
	}{
		{
			Data:       []string{"apple"},
			ExpectedSizes: []int{1},
		},
		{
			Data:       []string{"apple", "banana"},
			ExpectedSizes: []int{1,       2},
		},
		{
			Data:       []string{"apple", "banana", "cherry"},
			ExpectedSizes: []int{1,       2,        3},
		},


		{
			Data:       []string{"apple", "apple", "apple"},
			ExpectedSizes: []int{1,       1,       1},
		},
		{
			Data:       []string{"banana", "banana", "banana"},
			ExpectedSizes: []int{1,       1,       1},
		},
		{
			Data:       []string{"cherry", "cherry", "cherry"},
			ExpectedSizes: []int{1,       1,       1},
		},


		{
			Data:       []string{"apple", "apple"},
			ExpectedSizes: []int{1,       1},
		},
		{
			Data:       []string{"apple", "apple", "banana"},
			ExpectedSizes: []int{1,       1,       2},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana"},
			ExpectedSizes: []int{1,       1,       2,        2},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana", "cherry"},
			ExpectedSizes: []int{1,       1,       2,        2,        3},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana", "cherry", "cherry"},
			ExpectedSizes: []int{1,       1,       2,        2,        3,        3},
		},


		{
			Data:       []string{"apple", "banana", "apple", "banana", "apple", "banana"},
			ExpectedSizes: []int{1,       2,        2,       2,        2,       2},
		},
		{
			Data:       []string{"apple", "banana", "cherry", "apple", "banana", "cherry", "apple", "banana"},
			ExpectedSizes: []int{1,       2,        3,        3,       3,        3,        3,       3},
		},
	}


	// Create some random tests.
	const NUM_RANDOM_TESTS = 20

	for i:=0; i<NUM_RANDOM_TESTS; i++ {
		testLength := 1 + randomness.Intn(19)

		test := struct{
			Data []string
			ExpectedSizes []int
		}{
			Data: make([]string, testLength),
			ExpectedSizes: make([]int, testLength),
		}

		stringSet := make(map[string]struct{})
		for ii:=0; ii<testLength; ii++ {
			str := fmt.Sprintf("num-%d", randomness.Intn(5))

			test.Data[ii] = str

			previousExpectedSize := 0
			if 0 != ii {
				previousExpectedSize = test.ExpectedSizes[ii-1]
			}

			if _, ok := stringSet[str]; !ok {
				test.ExpectedSizes[ii] = previousExpectedSize+1
			} else {
				test.ExpectedSizes[ii] = previousExpectedSize
			}

			stringSet[str] = struct{}{}
		}

		tests = append(tests, test)
	}


	// Do tests.
Loop:	for testNumber, test := range tests {
		inspector := New()

		if expected, actual := 0, len(inspector.(*internalInspector).registry); expected != actual {
			t.Errorf("For test %#d, expected initial registry size to be empty (i.e., %d), but actually was %d.", testNumber, expected, actual)
			continue
		}

		for datumNumber, datum := range test.Data {
			inspector2 := inspector.Register(datum)
			if inspector2 != inspector {
				t.Errorf("For test #%d and datum #%d, expected Register to return the same inspector but didn't.", testNumber, datumNumber)
				continue Loop
			}

			if expected, actual := test.ExpectedSizes[datumNumber], len(inspector.(*internalInspector).registry); expected != actual {
				t.Errorf("For test #%d and datum #%d, expected registry size to be %d, but actually was %d.", testNumber, datumNumber, expected, actual)
				continue Loop
			}

			if _, ok := inspector.(*internalInspector).registry[datum]; !ok {
				t.Errorf("For test #%d and datum #%d, expected %q to be in registry, but wasn't", testNumber, datumNumber, datum)
				continue Loop
			}
		}
	}
}


func TestValidate(t *testing.T) {

	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )


	// Create some specific tests.
	tests := []struct{
		Data []string
	}{
		{
			Data:       []string{"apple"},
		},
		{
			Data:       []string{"apple", "banana"},
		},
		{
			Data:       []string{"apple", "banana", "cherry"},
		},


		{
			Data:       []string{"apple", "apple", "apple"},
		},
		{
			Data:       []string{"banana", "banana", "banana"},
		},
		{
			Data:       []string{"cherry", "cherry", "cherry"},
		},


		{
			Data:       []string{"apple", "apple"},
		},
		{
			Data:       []string{"apple", "apple", "banana"},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana"},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana", "cherry"},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana", "cherry", "cherry"},
		},


		{
			Data:       []string{"apple", "banana", "apple", "banana", "apple", "banana"},
		},
		{
			Data:       []string{"apple", "banana", "cherry", "apple", "banana", "cherry", "apple", "banana"},
		},
	}


	// Create some random tests.
	const NUM_RANDOM_TESTS = 20

	for i:=0; i<NUM_RANDOM_TESTS; i++ {
		testLength := 1 + randomness.Intn(19)

		test := struct{
			Data []string
		}{
			Data: make([]string, testLength),
		}

		for ii:=0; ii<testLength; ii++ {
			str := fmt.Sprintf("num-%d", randomness.Intn(5))

			test.Data[ii] = str
		}

		tests = append(tests, test)
	}


	// Do tests.
Loop:	for testNumber, test := range tests {
		inspector := New()

		for datumNumber, datum := range test.Data {
			if err := inspector.Validate(datum); nil == err {
				t.Errorf("For test #%d and datum #%d, expected nothing to validate with the inspector, but %q did.", testNumber, datumNumber, datum)
				continue Loop
			}
		}


		for datumNumber, datum := range test.Data {
			inspector.Register(datum)

			if err := inspector.Validate(datum); nil != err {
				t.Errorf("For test #%d and datum #%d, expected what was just registered (%q) to validate with the inspector but it didn't, because: %v.", testNumber, datumNumber, datum, err)
				continue Loop
			}
		}

		const DATUM_THAT_SHOULD_NOT_VALIDATE = "THIS DATUM SHOULD NOT VALIDATE"
		if err := inspector.Validate(DATUM_THAT_SHOULD_NOT_VALIDATE); nil == err {
			t.Errorf("For test #%d, %q should not have validated, with the inspector, but it did.", testNumber, DATUM_THAT_SHOULD_NOT_VALIDATE)
			continue Loop
		}
	}
}


func TestInspect(t *testing.T) {

	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )


	// Create some specific tests.
	tests := []struct{
		Data []string
	}{
		{
			Data:       []string{"apple"},
		},
		{
			Data:       []string{"apple", "banana"},
		},
		{
			Data:       []string{"apple", "banana", "cherry"},
		},


		{
			Data:       []string{"apple", "apple", "apple"},
		},
		{
			Data:       []string{"banana", "banana", "banana"},
		},
		{
			Data:       []string{"cherry", "cherry", "cherry"},
		},


		{
			Data:       []string{"apple", "apple"},
		},
		{
			Data:       []string{"apple", "apple", "banana"},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana"},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana", "cherry"},
		},
		{
			Data:       []string{"apple", "apple", "banana", "banana", "cherry", "cherry"},
		},


		{
			Data:       []string{"apple", "banana", "apple", "banana", "apple", "banana"},
		},
		{
			Data:       []string{"apple", "banana", "cherry", "apple", "banana", "cherry", "apple", "banana"},
		},
	}


	// Create some random tests.
	const NUM_RANDOM_TESTS = 20

	for i:=0; i<NUM_RANDOM_TESTS; i++ {
		testLength := 1 + randomness.Intn(19)

		test := struct{
			Data []string
		}{
			Data: make([]string, testLength),
		}

		for ii:=0; ii<testLength; ii++ {
			str := fmt.Sprintf("num-%d", randomness.Intn(5))

			test.Data[ii] = str
		}

		tests = append(tests, test)
	}


	// Do tests.
Loop:	for testNumber, test := range tests {
		inspector := New()

		for datumNumber, datum := range test.Data {
			if _, err := inspector.(*internalInspector).inspect(func()string{return datum}); nil == err {
				t.Errorf("For test #%d and datum #%d, expected nothing to validate with the inspector, but %q did.", testNumber, datumNumber, datum)
				continue Loop
			}
		}


		for datumNumber, datum := range test.Data {
			inspector.Register(datum)

			if actual, err := inspector.(*internalInspector).inspect(func()string{return datum}); nil != err {
				t.Errorf("For test #%d and datum #%d, expected what was just registered (%q) to validate with the inspector but it didn't, because: %v.", testNumber, datumNumber, datum, err)
				continue Loop
			} else if expected := datum; expected != actual {
				t.Errorf("For test #%d and datum #%d, expected what was just registered %q to to be returned by inspector but actually got %q.", testNumber, datumNumber, expected, actual)
				continue Loop
			}
		}

		const DATUM_THAT_SHOULD_NOT_VALIDATE = "THIS DATUM SHOULD NOT VALIDATE"
		if _, err := inspector.(*internalInspector).inspect(func()string{return DATUM_THAT_SHOULD_NOT_VALIDATE}); nil == err {
			t.Errorf("For test #%d, %q should not have validated, with the inspector, but it did.", testNumber, DATUM_THAT_SHOULD_NOT_VALIDATE)
			continue Loop
		}
	}
}
