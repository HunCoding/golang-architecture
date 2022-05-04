package singleresponsibilityprinciple

import "strings"

type JournalClassic struct {
	News []string
}

func (j *JournalClassic) String() string {
	return strings.Join(j.News, "\n")
}

func (j *JournalClassic) AddNews(news string) {
	//...
}

func (j *JournalClassic) RemoveNews(news string) {
	//...
}

func (j *JournalClassic) UpdateNews(news string) {
	//...
}

func (j *JournalClassic) DeleteNews(news string) {
	//...
}
