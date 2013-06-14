package main

import (
	"flag"
	"fmt"
	"github.com/gilesp/markov"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

var r *rand.Rand

func main() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	mainfile, secondaryfile, tweet := parseFlags()

	mainText := loadTextFromFile(mainfile)
	secondaryText := loadTextFromFile(secondaryfile)

	splitter := markov.NewDubSplitter()

	mainPhrases := splitter.Split(mainText)
	mixinPhrases := splitter.Split(secondaryText)

	remix := []string{}
	randomPhrase := ""
	randomMixin := ""
	for len(mainPhrases) > 0 {
		randomPhrase, mainPhrases = selectRandomPhrase(mainPhrases)
		remix = append(remix, randomPhrase)
		if r.Intn(4) == 3 && len(mixinPhrases) > 0 {
			//1 in 4 chance of adding a mixin phrase, if there are any left to add
			randomMixin, mixinPhrases = selectRandomPhrase(mixinPhrases)
			remix = append(remix, randomMixin)
		}
		if tweet && (utf8.RuneCountInString(strings.Join(remix, " ")) >= 160) {
			break
		}
	}
	/*
		sentences := markov.NewNaiveSplitter().Split(strings.Join(remix, " "))
		for _, sentence := range sentences {
			fmt.Println(sentence)
		}
	*/
	fmt.Println(strings.Join(remix, " "))

	if len(mixinPhrases) > 0 {
		fmt.Println("\n\nLeftovers:")
		for _, leftover := range mixinPhrases {
			fmt.Println(leftover)
		}
	}

}

func selectRandomPhrase(phrases []string) (phrase string, fewerPhrases []string) {
	phraseLength := len(phrases)
	i := r.Intn(phraseLength)
	phrase = phrases[i]
	phrases[i] = phrases[phraseLength-1]
	fewerPhrases = phrases[:phraseLength-1]
	return
}

func parseFlags() (mainfile string, secondaryfile string, tweet bool) {
	flag.StringVar(&mainfile, "main", "main.txt", "File containing main text to remix")
	flag.StringVar(&secondaryfile, "secondary", "", "File containing flavour text to mix in to the main text")
	flag.BoolVar(&tweet, "tweet", false, "Produce 160 characters or less")
	flag.Parse()
	return
}

func loadTextFromFile(filename string) string {
	if filename != "" {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Unable to open file ", filename)
			return ""
		} else {
			return string(content)
		}
	}
	return ""
}
