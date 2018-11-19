package main

import (
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// Emoticon2EmojiPlugin is a Mattermost plugin that replace text emoticons in messages by an emoji approximation
type Emoticon2EmojiPlugin struct {
	plugin.MattermostPlugin
	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex
	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *Emoticon2EmojiPluginConfiguration
}

// OnActivate register the plugin command
func (p *Emoticon2EmojiPlugin) OnActivate() error {
	return p.OnConfigurationChange()
}

// MessageWillBePosted converts emoticons in new posts
func (p *Emoticon2EmojiPlugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	post.Message = p.translate(post.Message)
	return post, ""
}

// MessageWillBeUpdated converts emoticons in edited posts
func (p *Emoticon2EmojiPlugin) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
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
	plugin.ClientMain(&Emoticon2EmojiPlugin{})
}
