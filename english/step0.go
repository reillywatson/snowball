package english

import (
	"github.com/reillywatson/snowball/snowballword"
)

var (
	sApostS = snowballword.MakeSuffix("'s'")
	apostS  = snowballword.MakeSuffix("'s")
	apost   = snowballword.MakeSuffix("'")
)

// Step 0 is to strip off apostrophes and "s".
//
func step0(w *snowballword.SnowballWord) bool {
	suffix := w.FirstSuffixA(sApostS, apostS, apost)
	if suffix == nil {
		return false
	}
	w.RemoveLastNRunes(len(suffix.Runes))
	return true
}
