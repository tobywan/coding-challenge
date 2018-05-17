package roman

import (
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	in   string
	want int
}
type errCase struct {
	in   string
	want string
}

func TestSimpleChars(t *testing.T) {
	tcs := []testCase{
		{"", 0},
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
	}

	for _, tc := range tcs {
		got, err := From(tc.in)
		if err != nil {
			t.Errorf("TestSimpleChars(%q) -> %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("TestSimpleChars(%q)=%d, want %d", tc.in, got, tc.want)
		}
	}
}
func TestSubtractors(t *testing.T) {
	tcs := []testCase{
		{"IV", 4},
		{"XIV", 14},
		{"MCDXLIV", 1444},
		{"CMXC", 990},
	}

	for _, tc := range tcs {
		got, err := From(tc.in)
		if err != nil {
			t.Errorf("TestSubtractors(%q) -> %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("TestSubtractors(%q)=%d, want %d", tc.in, got, tc.want)
		}
	}
}
func TestSimpleChords(t *testing.T) {
	tcs := []testCase{
		{"III", 3},
		{"XVI", 16},
		{"XXXVIII", 38},
		{"LI", 51},
		{"CCCXV", 315},
		{"DV", 505},
		{"MM", 2000},
		{"MDCCCLXXXVIII", 1888},
	}

	for _, tc := range tcs {
		got, err := From(tc.in)
		if err != nil {
			t.Errorf("TestSimpleChords(%q) -> %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("TestSimpleChords(%q)=%d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestRandomSample(t *testing.T) {
	tcs := []testCase{
		{"LX", 60},
		{"LXV", 65},
		{"CXXXI", 131},
		{"CL", 150},
		{"CLXI", 161},
		{"CCXCI", 291},
		{"CCCLXXVI", 376},
		{"CCCXCIX", 399},
		{"CDVII", 407},
		{"CDXIV", 414},
		{"CDLVII", 457},
		{"CDLVIII", 458},
		{"CDLXXXIX", 489},
		{"DXLV", 545},
		{"DLII", 552},
		{"DLXXII", 572},
		{"DLXXXII", 582},
		{"DLXXXVII", 587},
		{"DCII", 602},
		{"DCLVIII", 658},
		{"DCXCIX", 699},
		{"DCCLXXIV", 774},
		{"DCCCXXV", 825},
		{"DCCCXLIV", 844},
		{"DCCCXLVI", 846},
		{"DCCCXCIII", 893},
		{"CMXIII", 913},
		{"CMXXIV", 924},
		{"CMLVI", 956},
		{"CMLXVI", 966},
		{"CMXCIII", 993},
		{"CMXCIII", 993},
		{"MIX", 1009},
		{"MIX", 1009},
		{"MXXVII", 1027},
		{"MXXVIII", 1028},
		{"MXLIX", 1049},
		{"MLXIV", 1064},
		{"MLXVIII", 1068},
		{"MCLXVIII", 1168},
		{"MCXCI", 1191},
		{"MCCXXIX", 1229},
		{"MCCXLV", 1245},
		{"MCCXLVII", 1247},
		{"MCCLXV", 1265},
		{"MCCLXVI", 1266},
		{"MCCLXXXVI", 1286},
		{"MCCCIII", 1303},
		{"MCCCXXVIII", 1328},
		{"MCCCXXXIX", 1339},
		{"MCDLV", 1455},
		{"MDLXXVIII", 1578},
		{"MDXCII", 1592},
		{"MDCXXXII", 1632},
		{"MDCLVI", 1656},
		{"MDCLIX", 1659},
		{"MDCCXLIX", 1749},
		{"MDCCLXXV", 1775},
		{"MDCCLXXIX", 1779},
		{"MDCCCXVII", 1817},
		{"MDCCCXXI", 1821},
		{"MDCCCLXXXV", 1885},
		{"MCMXVII", 1917},
		{"MCMXXVI", 1926},
		{"MCMXXIX", 1929},
		{"MCMXLIX", 1949},
		{"MCMLVIII", 1958},
		{"MMVIII", 2008},
		{"MMXX", 2020},
		{"MMCV", 2105},
		{"MMCXL", 2140},
		{"MMCLXXX", 2180},
		{"MMCLXXXIII", 2183},
		{"MMCXCIX", 2199},
		{"MMCCCV", 2305},
		{"MMCCCXIV", 2314},
		{"MMCCCXXIII", 2323},
		{"MMCCCXXVI", 2326},
		{"MMCCCXXVIII", 2328},
		{"MMCCCLXVII", 2367},
		{"MMCDXXVIII", 2428},
		{"MMCDXL", 2440},
		{"MMDXXI", 2521},
		{"MMDXXIV", 2524},
		{"MMDLVIII", 2558},
		{"MMDLXVIII", 2568},
		{"MMDXCIV", 2594},
		{"MMDCXV", 2615},
		{"MMDCLXX", 2670},
		{"MMDCCLVIII", 2758},
		{"MMDCCLXI", 2761},
		{"MMDCCLXXXIII", 2783},
		{"MMDCCCIII", 2803},
		{"MMDCCCXI", 2811},
		{"MMDCCCXVII", 2817},
		{"MMDCCCLXV", 2865},
		{"MMCMXXIX", 2929},
		{"MMCMXXXIX", 2939},
		{"MMCMXCVII", 2997},
		{"MMM", 3000},
	}

	for _, tc := range tcs {
		got, err := From(tc.in)
		if err != nil {
			t.Errorf("TestSimpleChords(%q) -> %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("TestSimpleChords(%q)=%d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestCharacterErrors(t *testing.T) {
	format := "invalid character %q in %q"

	tcs := []errCase{
		{"MMMMMm", fmt.Sprintf(format, 'm', "MMMMMm")},
		{".", fmt.Sprintf(format, '.', ".")},
		{"X1234-_X", fmt.Sprintf(format, '1', "X1234-_X")},
	}

	for _, tc := range tcs {
		_, err := From(tc.in)

		if err.Error() != tc.want {
			t.Errorf("TestErrors(%q) -> %v, want %v", tc.in, err, tc.want)
		}
	}
}

func TestSequenceErrors(t *testing.T) {

	tcs := []errCase{
		{"IIII", "maximum occurences exceeded"},
		{"IL", "invalid order"},
		{"VMI", "invalid order"},
		{"IIV", "multiple subtractors"},
		{"MCMC", "reuse as subractor"},
		{"CMC", "reuse as subractor"},
		{"MCMCXXXVVIV", "reuse as subractor"},
		{"XCC", "subtracting from multiple occurrence"},
		{"XCCC", "subtracting from multiple occurrence"},
		{"CCXCC", "subtracting from multiple occurrence"},
	}

	for _, tc := range tcs {
		_, err := From(tc.in)

		if err == nil {
			t.Errorf("expected error for %s", tc.in)
			continue
		}
		if !strings.Contains(err.Error(), tc.want) {
			t.Errorf("TestErrors(%q) -> %v, want %v", tc.in, err, tc.want)
		}
	}
}
