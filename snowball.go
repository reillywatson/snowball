package snowball

import (
	"fmt"
	"github.com/reillywatson/snowball/english"
	"github.com/reillywatson/snowball/french"
	"github.com/reillywatson/snowball/russian"
	"github.com/reillywatson/snowball/spanish"
)

const (
	VERSION string = "v0.3.4"
)

// Stem a word in the specified language.
//
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {

	var f func(string, bool) string
	switch language {
	case "english":
		f = english.Stem
	case "spanish":
		f = spanish.Stem
	case "french":
		f = french.Stem
	case "russian":
		f = russian.Stem
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	stemmed = f(word, stemStopWords)
	return

}
