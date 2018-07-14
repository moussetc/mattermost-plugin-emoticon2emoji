package main

import (
	"net/http"
	"sync/atomic"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/mattermost/mattermost-server/plugin/rpcplugin"
)

// Emoticon2EmojiPlugin is a Mattermost plugin that replace text emoticons in messages by an emoji approximation
type Emoticon2EmojiPlugin struct {
	api           plugin.API
	configuration atomic.Value
	matches       map[string]match
}

// OnActivate register the plugin command
func (p *Emoticon2EmojiPlugin) OnActivate(api plugin.API) error {
	p.api = api

	return p.OnConfigurationChange()
}

// OnConfigurationChange apply a new plugin configuration
func (p *Emoticon2EmojiPlugin) OnConfigurationChange() error {
	var configuration Emoticon2EmojiPluginConfiguration
	if err := p.api.LoadPluginConfiguration(&configuration); err != nil {
		return appError("Unable to load new plugin configuration", err)
	}

	p.configuration.Store(&configuration)
	if err := p.applyNewConfig(&configuration); err != nil {
		return appError("Unable to apply new plugin configuration", err)
	}
	return nil
}

// MessageWillBePosted converts emoticons in new posts
func (p *Emoticon2EmojiPlugin) MessageWillBePosted(post *model.Post) (*model.Post, string) {
	post.Message = p.translate(post.Message)
	return post, ""
}

// MessageWillBeUpdated converts emoticons in edited posts
func (p *Emoticon2EmojiPlugin) MessageWillBeUpdated(newPost, oldPost *model.Post) (*model.Post, string) {
	newPost.Message = p.translate(newPost.Message)
	return newPost, ""
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
