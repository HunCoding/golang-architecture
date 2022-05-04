package singleresponsibilityprinciple

import (
	"io/ioutil"
	"net/url"
	"strings"
)

type Journal struct {
	News []string
}

func (j *Journal) String() string {
	return strings.Join(j.News, "\n")
}

func (j *Journal) AddNews(news string) {
	//...
}

func (j *Journal) RemoveNews(news string) {
	//...
}

func (j *Journal) UpdateNews(news string) {
	//...
}

func (j *Journal) DeleteNews(news string) {
	//...
}

func (j *Journal) PrintNews(filename string, news string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) LoadNewsFromFile(filename string) {
	//...
}

func (j *Journal) LoadNewsFromWeb(website *url.URL) {
	//...
}
