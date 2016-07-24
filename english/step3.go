package english

import (
	"github.com/kljensen/snowball/snowballword"
)

var (
	alize       = snowballword.MakeSuffix("alize")
	icate       = snowballword.MakeSuffix("icate")
	ative       = snowballword.MakeSuffix("ative")
	iciti       = snowballword.MakeSuffix("iciti")
	ical        = snowballword.MakeSuffix("ical")
	ness        = snowballword.MakeSuffix("ness")
	emptySuffix = snowballword.MakeSuffix("")
)

// Step 3 is the stemming of various longer sufficies
// found in R1.
//
func step3(w *snowballword.SnowballWord) bool {

	suffix := w.FirstSuffixA(
		ational, tional, alize, icate, ative,
		iciti, ical, ful, ness,
	)

	// If it is not in R1, do nothing
	if suffix == nil || len(suffix.Runes) > len(w.RS)-w.R1start {
		return false
	}

	// Handle special cases where we're not just going to
	// replace the suffix with another suffix: there are
	// other things we need to do.
	//
	if suffix == ative {

		// If in R2, delete.
		//
		if len(w.RS)-w.R2start >= 5 {
			w.RemoveLastNRunes(len(suffix.Runes))
			return true
		}
		return false
	}

	// Handle a suffix that was found, which is going
	// to be replaced with a different suffix.
	//
	var repl *snowballword.Suffix
	switch suffix {
	case ational:
		repl = ate
	case tional:
		repl = tion
	case alize:
		repl = al
	case icate, iciti, ical:
		repl = ic
	case ful, ness:
		repl = emptySuffix
	}
	if repl != nil {
		w.ReplaceSuffixRunes(suffix.Runes, repl.Runes, true)
	}
	return true

}
