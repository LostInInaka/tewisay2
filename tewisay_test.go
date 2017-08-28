package main

import (
	"testing"
)

func TestBalloon(t *testing.T) {
	tests := []string{
		"this string has lots of unicodes\n" +
			"☆てを可愛い✈testعربIԉｔëｒｎåｔíòｎɑɭｉｚɑｔｉ߀ｎ٩(͡๏̯͡๏)۶ ٩(-̮̮̃•̃).",
		"this has lots of colors\n" +
			"\x1b[32m   2.4M               \x1b[42m______\x1b[m \x1b[32m.\x1b[m\n" +
			"\x1b[33m   1.2M                  \x1b[43m___\x1b[m \x1b[33m.git/\x1b[m\n" +
			"\x1b[34m    36K                    \x1b[44m_\x1b[m \x1b[34mcows/\x1b[m\n" +
			"\x1b[35m    12K                    \x1b[45m_\x1b[m \x1b[35mswidth/\x1b[m\n" +
			"\x1b[36m   9.1M \x1b[46m____________________\x1b[m \x1b[36mtewisay-git/\x1b[m\n" +
			"\x1b[1m 12.75M                      Total\n" +
			"\x1b[m",
		"this\tone\thas\t\ttabs\ntetete\t\tte\nte\t\ttete",
	}
	expect := []string{
		`┌────────────────────────────────────────────────────────────────┐
│ this string has lots of unicodes[0m                               │
│ ☆てを可愛い✈testعربIԉｔëｒｎåｔíòｎɑɭｉｚɑｔｉ߀ｎ٩(͡๏̯͡๏)۶ ٩(-̮̮̃•̃).[0m │
└────────────────────────────────────────────────────────────────┘`,
		`┌───────────────────────────────────────────┐
│ this has lots of colors[0m                   │
│ [32m   2.4M               [42m______[m [32m.[m[0m            │
│ [32m[42m[m[32m[m[33m   1.2M                  [43m___[m [33m.git/[m[0m        │
│ [33m[43m[m[33m[m[34m    36K                    [44m_[m [34mcows/[m[0m        │
│ [34m[44m[m[34m[m[35m    12K                    [45m_[m [35mswidth/[m[0m      │
│ [35m[45m[m[35m[m[36m   9.1M [46m____________________[m [36mtewisay-git/[m[0m │
│ [36m[46m[m[36m[m[1m 12.75M                      Total[0m        │
└───────────────────────────────────────────┘`,
		`┌────────────────────────────────────┐
│ this	one	has		tabs[0m │
│ tetete		te[0m           │
│ te		tete[0m                 │
└────────────────────────────────────┘`,
	}

	for i := range tests {
		ret := balloon(tests[i], borders["unicode"])
		if ret != expect[i] {
			t.Errorf("Expected:\n%s\nGot:\n%s\n", expect[i], ret)
		}
	}
}
