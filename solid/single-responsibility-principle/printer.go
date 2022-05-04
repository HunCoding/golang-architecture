package singleresponsibilityprinciple

import (
	"io/ioutil"
)

func PrintNews(filename string, journal JournalClassic) {
	_ = ioutil.WriteFile(filename, []byte(journal.String()), 0644)
}
