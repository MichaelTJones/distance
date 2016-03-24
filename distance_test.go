package distance

import (
	"math"
	"testing"
)

const variance = 0.000000001

var jaroTests = []struct {
	a, b string  // strings
	d    float64 // distamce
}{
	// from Approximate String Comparison and its Effect on an Advanced Record Linkage System
	// by Edward H. Porter and William E. Winkler, U.S. Bureau of the Census
	// https://www.census.gov/srd/papers/pdf/rr97-2.pdf
	{"SHACKLEFORD", "SHACKELFORD", 32.0 / 33.0},   // 0.9696969696969697
	{"CUNNIGHAM", "DUNNINGHAM", 121.0 / 135.0},    // 0.8962962962962963
	{"NICHLESON", "NICHULSON", 25.0 / 27.0},       // 0.9259259259259259
	{"JONES", "JOHNSON", 83.0 / 105.0},            // 0.7904761904761904
	{"MASSEY", "MASSIE", 8.0 / 9.0},               // 0.8888888888888888
	{"ABROMS", "ABRAMS", 8.0 / 9.0},               // 0.8888888888888888
	{"HARDIN", "MARTINEZ", 13.0 / 18.0},           // 0.7222222222222222
	{"ITMAN", "SMITH", 7.0 / 15.0},                // 0.4666666666666667
	{"JERALDINE", "GERALDINE", 25.0 / 27.0},       // 0.9259259259259259
	{"MARHTA", "MARTHA", 17.0 / 18.0},             // 0.9444444444444444
	{"MICHAEL", "MICHELLE", 73.0 / 84.0},          // 0.8690476190476191
	{"JULIES", "JULIUS", 8.0 / 9.0},               // 0.8888888888888888
	{"TANYA", "TONYA", 13.0 / 15.0},               // 0.8666666666666667
	{"DUANE", "DWAYNE", 37.0 / 45.0},              // 0.8222222222222222
	{"SEAN", "SUSAN", 47.0 / 60.0},                // 0.7833333333333333
	{"JON", "JOHN", 11.0 / 12.0},                  // 0.9166666666666666
	{"JON", "JAN", 7.0 / 9.0},                     // 0.7777777777777778
	{"BROOKHAVEN", "BRROKHAVEN", 14.0 / 15.0},     // 0.9333333333333333
	{"BROOK HLLW", "BROOK HALLOW", 17.0 / 18.0},   // 0.9444444444444444
	{"DECATUR", "DECATIR", 19.0 / 21.0},           // 0.9047619047619048
	{"FITZRUREITER", "FITZENREITER", 77.0 / 90.0}, // 0.8555555555555555
	{"HIGBEE", "HIGHEE", 8.0 / 9.0},               // 0.8888888888888888
	{"HIGBEE", "HIGVEE", 8.0 / 9.0},               // 0.8888888888888888
	{"LACURA", "LOCURA", 8.0 / 9.0},               // 0.8888888888888888
	{"IOWA", "IONA", 5.0 / 6.0},                   // 0.8333333333333334
	{"1ST", "IST", 7.0 / 9.0},                     // 0.7777777777777778
}

func TestJaro(t *testing.T) {
	for i, a := range jaroTests {
		d := Jaro(a.a, a.b)
		if math.Abs(d-a.d) > variance {
			t.Errorf("#%d, Jaro(%q,%q) value is %v; want %v", i, a.a, a.b, d, a.d)
		}
	}
}

func BenchmarkJaro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := &jaroTests[i%len(jaroTests)]
		_ = Jaro(j.a, j.b)
	}
}

var winklerTests = []struct {
	a, b string  // strings
	d    float64 // distamce
}{
	// from the Wikipedia article on Jaro–Winkler distance
	// en.wikipedia.org/wiki/Jaro–Winkler_distance
	{"CRATE", "TRACE", 11.0 / 15.0},     // 0.7333333333333333
	{"martha", "marhta", 173.0 / 180.0}, // 0.9611111111111111
	{"dwayne", "duane", 21.0 / 25.0},    // 0.84
	{"DIXON", "DICKSONX", 61.0 / 75.0},  // 0.8133333333333334

	// from Data Matching: Concepts and Techniques for Record Linkage, Entity Resolution, and Duplicate Detection
	// by Peter Christen
	// page 109
	// http://www.springer.com/us/book/9783642311635
	{"shackleford", "shackelford", 54.0 / 55.0}, // 0.9818181818181818
	{"nichleson", "nichulson", 43.0 / 45.0},     // 0.9555555555555556
	{"jones", "johnson", 437.0 / 525.0},         // 0.8323809523809523
	{"massey", "massie", 14.0 / 15.0},           // 0.9333333333333333
	{"jeraldine", "geraldine", 25.0 / 27.0},     // 0.9259259259259259
	{"michelle", "michael", 129.0 / 140.0},      // 0.9214285714285714

	// from Matching and Record Linkage
	// by William E. Winkler, U.S. Bureau of the Census
	// https://www.census.gov/srd/papers/pdf/rr93-8.pdf
	// Table 5 (but not "range adjusted" so values slightly different)
	{"billy", "billy", 1.0},               // 1
	{"billy", "bill", 24.0 / 25.0},        // 0.96
	{"billy", "blily", 47.0 / 50.0},       // 0.94
	{"massie", "massey", 14.0 / 15.0},     // 0.9333333333333333
	{"yvette", "yevett", 9.0 / 10.0},      // 0.9
	{"billy", "bolly", 22.0 / 25.0},       // 0.88
	{"dwayne", "duane", 21.0 / 25.0},      // 0.84
	{"dixon", "dickson", 437.0 / 525.0},   // 0.8323809523809523
	{"billy", "susan", 0.0},               // 0
	{"barnes", "anderson", 271.0 / 360.0}, // 0.7527777777777778

	// from Code Spelunking: Jaro-Winkler String Comparison
	// by Breck Baldwin
	// http://lingpipe-blog.com/2006/12/13/code-spelunking-jaro-winkler-string-comparison/
	{"ABCVWXYZ", "CABVWXYZ", 23.0 / 24.0},     // 0.9583333333333334
	{"ABCDUVWXYZ", "DBCAUVWXYZ", 29.0 / 30.0}, // 0.9666666666666667
	{"ABBBUVWXYZ", "BBBAUVWXYZ", 29.0 / 30.0}, // 0.9666666666666667
	{"ABCDUVWXYZ", "DABCUVWXYZ", 14.0 / 15.0}, // 0.9333333333333333

	// from searching the web for examples
	// http://stackoverflow.com/questions/20278373/peculiar-behaviour-of-jaro-distance-in-jellyfish
	{"Poverty", "Poervty", 101.0 / 105.0}, // 0.9619047619047619
	{"Poervty", "Poverty", 101.0 / 105.0}, // 0.9619047619047619
}

func TestWinkler(t *testing.T) {
	for i, a := range winklerTests {
		d := JaroWinkler(a.a, a.b)
		if math.Abs(d-a.d) > variance {
			t.Errorf("#%d, JaroWinkler(%q,%q) value is %v; want %v", i, a.a, a.b, d, a.d)
		}
	}
}

func BenchmarkWinkler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := &jaroTests[i%len(jaroTests)] // use same list of string pairs as BenchmarkJaro
		_ = JaroWinkler(j.a, j.b)
	}
}

var milneTests = []struct {
	a, b string  // strings
	j, w float64 // distamce
}{
	// from Rich Milne
	// https://github.com/richmilne/JaroWinkler/blob/master/jaro/jaro_tests.py
	{"SHACKLEFORD", "SHACKELFORD", 32.0 / 33.0, 54.0 / 55.0},    // 0.9696969696969697 0.9818181818181818
	{"DUNNINGHAM", "CUNNIGHAM", 121.0 / 135.0, 121.0 / 135.0},   // 0.8962962962962963 0.8962962962962963
	{"NICHLESON", "NICHULSON", 25.0 / 27.0, 43.0 / 45.0},        // 0.9259259259259259 0.9555555555555556
	{"JONES", "JOHNSON", 83.0 / 105.0, 437.0 / 525.0},           // 0.7904761904761904 0.8323809523809523
	{"MASSEY", "MASSIE", 8.0 / 9.0, 14.0 / 15.0},                // 0.8888888888888888 0.9333333333333333
	{"ABROMS", "ABRAMS", 8.0 / 9.0, 83.0 / 90.0},                // 0.8888888888888888 0.9222222222222223
	{"HARDIN", "MARTINEZ", 13.0 / 18.0, 13.0 / 18.0},            // 0.7222222222222222 0.7222222222222222
	{"ITMAN", "SMITH", 7.0 / 15.0, 7.0 / 15.0},                  // 0.4666666666666667 0.4666666666666667
	{"JERALDINE", "GERALDINE", 25.0 / 27.0, 25.0 / 27.0},        // 0.9259259259259259 0.9259259259259259
	{"MARTHA", "MARHTA", 17.0 / 18.0, 173.0 / 180.0},            // 0.9444444444444444 0.9611111111111111
	{"MICHELLE", "MICHAEL", 73.0 / 84.0, 129.0 / 140.0},         // 0.8690476190476191 0.9214285714285714
	{"JULIES", "JULIUS", 8.0 / 9.0, 14.0 / 15.0},                // 0.8888888888888888 0.9333333333333333
	{"TANYA", "TONYA", 13.0 / 15.0, 22.0 / 25.0},                // 0.8666666666666667 0.88
	{"DWAYNE", "DUANE", 37.0 / 45.0, 21.0 / 25.0},               // 0.8222222222222222 0.84
	{"SEAN", "SUSAN", 47.0 / 60.0, 161.0 / 200.0},               // 0.7833333333333333 0.805
	{"JON", "JOHN", 11.0 / 12.0, 14.0 / 15.0},                   // 0.9166666666666666 0.9333333333333333
	{"JON", "JAN", 7.0 / 9.0, 4.0 / 5.0},                        // 0.7777777777777778 0.8
	{"DWAYNE", "DYUANE", 37.0 / 45.0, 21.0 / 25.0},              // 0.8222222222222222 0.84
	{"CRATE", "TRACE", 11.0 / 15.0, 11.0 / 15.0},                // 0.7333333333333333 0.7333333333333333
	{"WIBBELLY", "WOBRELBLY", 1265.0 / 1512.0, 1433.0 / 1680.0}, // 0.8366402116402116 0.8529761904761904
	{"DIXON", "DICKSONX", 23.0 / 30.0, 61.0 / 75.0},             // 0.7666666666666667 0.8133333333333334
	{"MARHTA", "MARTHA", 17.0 / 18.0, 173.0 / 180.0},            // 0.9444444444444444 0.9611111111111111
	{"AL", "AL", 1.0, 1.0},                                      // 1 1
	{"aaaaaabc", "aaaaaabd", 11.0 / 12.0, 19.0 / 20.0},          // 0.9166666666666666 0.95
	{"ABCVWXYZ", "CABVWXYZ", 23.0 / 24.0, 23.0 / 24.0},          // 0.9583333333333334 0.9583333333333334
	{"ABCAWXYZ", "BCAWXYZ", 51.0 / 56.0, 51.0 / 56.0},           // 0.9107142857142857 0.9107142857142857
	{"ABCVWXYZ", "CBAWXYZ", 51.0 / 56.0, 51.0 / 56.0},           // 0.9107142857142857 0.9107142857142857
	{"ABCDUVWXYZ", "DABCUVWXYZ", 14.0 / 15.0, 14.0 / 15.0},      // 0.9333333333333333 0.9333333333333333
	{"ABCDUVWXYZ", "DBCAUVWXYZ", 29.0 / 30.0, 29.0 / 30.0},      // 0.9666666666666667 0.9666666666666667
	{"ABBBUVWXYZ", "BBBAUVWXYZ", 29.0 / 30.0, 29.0 / 30.0},      // 0.9666666666666667 0.9666666666666667
	{"ABCDUV11lLZ", "DBCAUVWXYZ", 563.0 / 770.0, 563.0 / 770.0}, // 0.7311688311688311 0.7311688311688311
	{"ABBBUVWXYZ", "BBB11L3VWXZ", 257.0 / 330.0, 257.0 / 330.0}, // 0.7787878787878788 0.7787878787878788
	{"-", "-", 1.0, 1.0},                                        // 1 1
	{"A", "A", 1.0, 1.0},                                        // 1 1
	{"AB", "AB", 1.0, 1.0},                                      // 1 1
	{"ABC", "ABC", 1.0, 1.0},                                    // 1 1
	{"ABCD", "ABCD", 1.0, 1.0},                                  // 1 1
	{"ABCDE", "ABCDE", 1.0, 1.0},                                // 1 1
	{"AA", "AA", 1.0, 1.0},                                      // 1 1
	{"AAA", "AAA", 1.0, 1.0},                                    // 1 1
	{"AAAA", "AAAA", 1.0, 1.0},                                  // 1 1
	{"AAAAA", "AAAAA", 1.0, 1.0},                                // 1 1
	{"A", "B", 0.0, 0.0},                                        // 0 0
	{"-", "ABC", 0.0, 0.0},                                      // 0 0
	{"ABCD", "-", 0.0, 0.0},                                     // 0 0
	{"--", "-", 5.0 / 6.0, 17.0 / 20.0},                         // 0.8333333333333334 0.85
	{"--", "---", 8.0 / 9.0, 41.0 / 45.0},                       // 0.8888888888888888 0.9111111111111111
	{"mm", "mmm", 8.0 / 9.0, 41.0 / 45.0},                       // 0.8888888888888888 0.9111111111111111
}

func TestMilne(t *testing.T) {
	for i, a := range milneTests {
		j := Jaro(a.a, a.b)
		if math.Abs(j-a.j) > variance {
			t.Errorf("#%d, Jaro(%q,%q) value is %v; want %v", i, a.a, a.b, j, a.j)
		}

		w := JaroWinkler(a.a, a.b)
		if math.Abs(w-a.w) > variance {
			t.Errorf("#%d, JaroWinkler(%q,%q) value is %v; want %v", i, a.a, a.b, w, a.w)
		}
	}
}
