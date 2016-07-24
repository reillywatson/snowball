package english

import (
	"github.com/kljensen/snowball/snowballword"
)

var (
	eedly = snowballword.MakeSuffix("eedly")
	ingly = snowballword.MakeSuffix("ingly")
	edly  = snowballword.MakeSuffix("edly")
	ing   = snowballword.MakeSuffix("ing")
	eed   = snowballword.MakeSuffix("eed")
	ed    = snowballword.MakeSuffix("ed")
	at    = snowballword.MakeSuffix("at")
	bl    = snowballword.MakeSuffix("bl")
	iz    = snowballword.MakeSuffix("iz")
	bb    = snowballword.MakeSuffix("bb")
	dd    = snowballword.MakeSuffix("dd")
	ff    = snowballword.MakeSuffix("ff")
	gg    = snowballword.MakeSuffix("gg")
	mm    = snowballword.MakeSuffix("mm")
	nn    = snowballword.MakeSuffix("nn")
	pp    = snowballword.MakeSuffix("pp")
	rr    = snowballword.MakeSuffix("rr")
	tt    = snowballword.MakeSuffix("tt")
)

// Step 1b is the normalization of various "ly" and "ed" sufficies.
//
func step1b(w *snowballword.SnowballWord) bool {

	suffix := w.FirstSuffixA(eedly, ingly, edly, ing, eed, ed)

	switch suffix {

	case nil:
		// No suffix found
		return false

	case eed, eedly:

		// Replace by ee if in R1
		if len(suffix.Runes) <= len(w.RS)-w.R1start {
			w.ReplaceSuffixRunes(suffix.Runes, []rune("ee"), true)
		}
		return true

	case ed, edly, ing, ingly:
		hasLowerVowel := false
		for i := 0; i < len(w.RS)-len(suffix.Runes); i++ {
			if isLowerVowel(w.RS[i]) {
				hasLowerVowel = true
				break
			}
		}
		if hasLowerVowel {

			// This case requires a two-step transformation and, due
			// to the way we've implemented the `ReplaceSuffix` method
			// here, information about R1 and R2 would be lost between
			// the two.  Therefore, we need to keep track of the
			// original R1 & R2, so that we may set them below, at the
			// end of this case.
			//
			originalR1start := w.R1start
			originalR2start := w.R2start

			// Delete if the preceding word part contains a vowel
			w.RemoveLastNRunes(len(suffix.Runes))

			// ...and after the deletion...

			newSuffix := w.FirstSuffixA(at, bl, iz, bb, dd, ff, gg, mm, nn, pp, rr, tt)
			switch newSuffix {

			case nil:

				// If the word is short, add "e"
				if isShortWord(w) {

					// By definition, r1 and r2 are the empty string for
					// short words.
					w.RS = append(w.RS, []rune("e")...)
					w.R1start = len(w.RS)
					w.R2start = len(w.RS)
					return true
				}

			case at, bl, iz:

				// If the word ends "at", "bl" or "iz" add "e"
				w.ReplaceSuffixRunes(newSuffix.Runes, append(newSuffix.Runes, rune('e')), true)

			case bb, dd, ff, gg, mm, nn, pp, rr, tt:

				// If the word ends with a double remove the last letter.
				// Note that, "double" does not include all possible doubles,
				// just those shown above.
				//
				w.RemoveLastNRunes(1)
			}

			// Because we did a double replacement, we need to fix
			// R1 and R2 manually. This is just becase of how we've
			// implemented the `ReplaceSuffix` method.
			//
			rsLen := len(w.RS)
			if originalR1start < rsLen {
				w.R1start = originalR1start
			} else {
				w.R1start = rsLen
			}
			if originalR2start < rsLen {
				w.R2start = originalR2start
			} else {
				w.R2start = rsLen
			}

			return true
		}

	}

	return false
}
