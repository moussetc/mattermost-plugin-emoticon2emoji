package main

import (
	"bytes"
	"encoding/json"
	"math"
	"regexp"
	"strings"
)

type Emoticon2EmojiPluginConfiguration struct {
	MatchesChoice string
	UserMatches   string
}

func (p *Emoticon2EmojiPlugin) applyNewConfig(configuration *Emoticon2EmojiPluginConfiguration) error {
	userMatches, err := unserializeConfigMatches(configuration.UserMatches)

	if err != nil {
		return appError("Unable to parse the emoticons to emojis matches list", err)
	}

	// user mappings > slack mappings > default mappings > mattermost mappings
	effectiveMap := map[string]string{}
	if configuration.MatchesChoice == "" {
		configuration.MatchesChoice = "slack_default_custom"
	}
	if strings.Contains(configuration.MatchesChoice, "default") {
		effectiveMap = DefaultMatches
	}
	if strings.Contains(configuration.MatchesChoice, "slack") {
		for k, v := range SlackMatches {
			effectiveMap[k] = v
		}
	}
	if strings.Contains(configuration.MatchesChoice, "custom") {
		for k, v := range userMatches {
			effectiveMap[k] = v
		}
	}
	p.matches = map[string]match{}
	for emoticon, emoji := range effectiveMap {
		p.matches[emoticon] = match{
			replacement: emoji,
			regexp:      getEmoticonRegexp(emoticon),
		}
	}

	return nil
}

func getEmoticonRegexp(emoticon string) *regexp.Regexp {
	return regexp.MustCompile("(\\s)+(" + regexp.QuoteMeta(emoticon) + ")(\\s)+")
}

// Read a string of serialized matches (JSON) into a map
func unserializeConfigMatches(matches string) (map[string]string, error) {
	if matches == "" {
		return map[string]string{}, nil
	}

	var matchesMap map[string]string
	matchesSerialized := bytes.NewBufferString(matches)
	d := json.NewDecoder(matchesSerialized)
	if err := d.Decode(&matchesMap); err != nil {
		return nil, appError("Unable to parse the emoticons to emojis matches list", err)
	}

	return matchesMap, nil
}

// Translate replace all the configured emoticons in a string by their equivalent as Mattermost emojis
func (p *Emoticon2EmojiPlugin) translate(input string) (result string) {
	if p.matches != nil {
		return translate(input, p.matches)
	}
	return input
}

// Translate replace all matches that form a single word by their intended match
func translate(input string, matches map[string]match) (result string) {
	text := input
	for emoticon, match := range matches {
		result := ""
		currentText := text[:]
		index := strings.Index(currentText, emoticon)
		for index > -1 {
			startIndex := int(math.Max(0, float64(index-1)))
			endIndex := int(math.Min(float64(len(currentText)-1), float64(index+len(emoticon))))
			slice := currentText[startIndex : endIndex+1]
			replacement := emoticon
			// only replace if the match is not part of a bigger "word"
			if match.regexp.MatchString(" " + slice + " ") {
				replacement = ":" + match.replacement + ":"
			}
			result = result + currentText[:index] + replacement
			// continue the search for the rest of the string only
			currentText = currentText[index+len(emoticon):]
			index = strings.Index(currentText, emoticon)
		}
		text = result + currentText
	}
	return text
}
