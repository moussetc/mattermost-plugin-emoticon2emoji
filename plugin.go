package main

import (
	"strings"
	"sync/atomic"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/mattermost/mattermost-server/plugin/rpcplugin"
)

// Emoticon2EmojiPlugin is a Mattermost plugin that replace text emoticons in messages by an emoji approximation
type Emoticon2EmojiPlugin struct {
	api           plugin.API
	configuration atomic.Value
	enabled       bool
	matches       map[string]string
}

// TODO 1 : complete list with https://en.wikipedia.org/wiki/List_of_emoticons
// TODO 2 : make it configurable with serialisation in the settings
var matches = map[string]string{
	":)": "slightly_smiling_face",
	":D": "smiley",
	";)": "wink",
	":(": "slightly_frowning_face",
}

// CreatePost translate emoticons in new posts
func (p *Emoticon2EmojiPlugin)MessageWillBePosted(post *model.Post) (*model.Post, string) {
	return p.translate(post)
}

// CreatePost translate emoticons in edited posts
func (p *Emoticon2EmojiPlugin)MessageWillBeUpdated(newPost, oldPost *model.Post) (*model.Post, string) {
	return p.translate(newPost)
}

// Translate replace all the configured emoticons in a post by their equivalent as Mattermost emojis
func (p *Emoticon2EmojiPlugin) translate(post *model.Post) (*model.Post, string) {
	// Only consider emoticons with a space on each side
	for emoticon, emoji := range matches {
		post.Message = strings.Replace(post.Message, " "+emoticon+" ", " :"+emoji+": ", -1);
	}
	return post, ""
}

// Install the RCP plugin
func main() {
	rpcplugin.Main(&Emoticon2EmojiPlugin{})
}
