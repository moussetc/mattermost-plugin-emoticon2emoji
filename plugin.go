package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/mattermost/mattermost-server/plugin/rpcplugin"
)

// Emoticon2EmojiPlugin is a Mattermost plugin that replace text emoticons in messages by an emoji approximation
type Emoticon2EmojiPlugin struct {
	api              plugin.API
	configuration    atomic.Value
	enabled          bool
	effectiveMatches atomic.Value
}

type Emoticon2EmojiPluginConfiguration struct {
	MatchesChoice string
	UserMatches   string
}

// OnActivate register the plugin command
func (p *Emoticon2EmojiPlugin) OnActivate(api plugin.API) error {
	p.api = api
	p.enabled = true

	return p.OnConfigurationChange()
}

// OnDeactivate handles plugin deactivation
func (p *Emoticon2EmojiPlugin) OnDeactivate() error {
	p.enabled = false
	return nil
}

// Load and return the plugin configuration
func (p *Emoticon2EmojiPlugin) config() *Emoticon2EmojiPluginConfiguration {
	return p.configuration.Load().(*Emoticon2EmojiPluginConfiguration)
}

// Load and return the configured matches
func (p *Emoticon2EmojiPlugin) matches() map[string]string {
	return p.effectiveMatches.Load().(map[string]string)
}

func (p *Emoticon2EmojiPlugin) OnConfigurationChange() error {
	var configuration Emoticon2EmojiPluginConfiguration
	if err := p.api.LoadPluginConfiguration(&configuration); err != nil {
		return appError("Unable to load new plugin configuration", err)
	}

	p.configuration.Store(&configuration)
	userMatches, err := p.unserializeMatches(configuration.UserMatches)

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
	p.effectiveMatches.Store(effectiveMap)
	return nil
}

// Unserialize a JSON string to a map of matches
func (p *Emoticon2EmojiPlugin) unserializeMatches(matches string) (map[string]string, error) {
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

// CreatePost translate emoticons in new posts
func (p *Emoticon2EmojiPlugin) MessageWillBePosted(post *model.Post) (*model.Post, string) {
	return p.translate(post)
}

// CreatePost translate emoticons in edited posts
func (p *Emoticon2EmojiPlugin) MessageWillBeUpdated(newPost, oldPost *model.Post) (*model.Post, string) {
	return p.translate(newPost)
}

// Translate replace all the configured emoticons in a post by their equivalent as Mattermost emojis
func (p *Emoticon2EmojiPlugin) translate(post *model.Post) (*model.Post, string) {
	for emoticon, emoji := range p.matches() {
		if strings.TrimSpace(post.Message) == emoticon {
			// Only an emoticon + whitespace
			post.Message = strings.Replace(post.Message, emoticon, ":"+emoji+":", -1)
		} else {
			// Only consider emoticons with a space on each side
			post.Message = strings.Replace(post.Message, " "+emoticon+" ", " :"+emoji+": ", -1)
		}
	}
	return post, ""
}

func appError(message string, err error) *model.AppError {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	return model.NewAppError("Emoticon2Emoji Plugin", message, nil, errorMessage, http.StatusBadRequest)
}

// Install the RCP plugin
func main() {
	rpcplugin.Main(&Emoticon2EmojiPlugin{})
}
