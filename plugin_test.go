package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mattermost/mattermost-server/model"

	"github.com/mattermost/mattermost-server/plugin/plugintest"
	"github.com/mattermost/mattermost-server/plugin/plugintest/mock"
)

func initTestPlugin(t *testing.T, config *Emoticon2EmojiPluginConfiguration) *plugintest.API {
	api := &plugintest.API{}
	api.On("LoadPluginConfiguration", mock.MatchedBy(func(x interface{}) bool { return true })).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*Emoticon2EmojiPluginConfiguration)
		dest.MatchesChoice = config.MatchesChoice
		dest.UserMatches = config.UserMatches
	}).Return(nil)
	return api
}

func TestNewPostWithDefaultEmoticon(t *testing.T) {
	p := Emoticon2EmojiPlugin{}
	config := Emoticon2EmojiPluginConfiguration{
		MatchesChoice: "slack_default_custom",
	}
	p.OnActivate(initTestPlugin(t, &config))

	post := &model.Post{
		Message: "Hello XD !!",
	}
	post, err := p.MessageWillBePosted(post)
	assert.NotNil(t, post)
	assert.Equal(t, "Hello :laughing: !!", post.Message)
	assert.Equal(t, "", err)
}

func TestUpdatedPostWithSlackEmoticon(t *testing.T) {
	p := Emoticon2EmojiPlugin{}
	config := Emoticon2EmojiPluginConfiguration{
		MatchesChoice: "slack_default_custom",
	}
	p.OnActivate(initTestPlugin(t, &config))

	post := &model.Post{
		Message: "Hello </3 !!",
	}
	post, err := p.MessageWillBeUpdated(post, nil)
	assert.NotNil(t, post)
	assert.Equal(t, "Hello :broken_heart: !!", post.Message)
	assert.Equal(t, "", err)
}

func TestReplaceWithOnlyEmoticon(t *testing.T) {
	var testMatches = map[string]match{
		"XD": match{
			replacement: "laughing",
			regexp:      getEmoticonRegexp("XD"),
		},
		"</3": match{
			replacement: "broken_heart",
			regexp:      getEmoticonRegexp("</3"),
		},
		"8)": match{
			replacement: "sunglasses",
			regexp:      getEmoticonRegexp("8)"),
		},
	}

	input := "XD"
	res := translate(input, testMatches)
	assert.Equal(t, ":laughing:", res)

	input = "  XD   \t"
	res = translate(input, testMatches)
	assert.Equal(t, "  :laughing:   \t", res)

	input = "XD  aaaaaa"
	res = translate(input, testMatches)
	assert.Equal(t, ":laughing:  aaaaaa", res)

	input = "aaaaaa XD"
	res = translate(input, testMatches)
	assert.Equal(t, "aaaaaa :laughing:", res)

	input = "XD\nXD\n XD \n XD\n\aaa XD bbbb\naaa XD\nXD bbbbb\naaaaaaXDbbbbb"
	res = translate(input, testMatches)
	assert.Equal(t, ":laughing:\n:laughing:\n :laughing: \n :laughing:\n\aaa :laughing: bbbb\naaa :laughing:\n:laughing: bbbbb\naaaaaaXDbbbbb", res)

	input = "aaaaaa XD </3 XDaaa aaa</3 XD XD 8)"
	res = translate(input, testMatches)
	assert.Equal(t, "aaaaaa :laughing: :broken_heart: XDaaa aaa</3 :laughing: :laughing: :sunglasses:", res)
}

func TestMappingsPrecedence(t *testing.T) {
	p := Emoticon2EmojiPlugin{}
	config := Emoticon2EmojiPluginConfiguration{
		MatchesChoice: "slack_default_custom",
		UserMatches:   "{\"XD\":\"cry\"}",
	}
	p.OnActivate(initTestPlugin(t, &config))

	post := &model.Post{
		Message: "XD",
	}
	post, err := p.MessageWillBeUpdated(post, nil)
	assert.NotNil(t, post)
	assert.Equal(t, ":cry:", post.Message)
	assert.Equal(t, "", err)
}

func TestMappingsPrecedenceWithoutCustom(t *testing.T) {
	p := Emoticon2EmojiPlugin{}
	config := Emoticon2EmojiPluginConfiguration{
		MatchesChoice: "slack_default",
		UserMatches:   "{\"XD\":\"cry\"}",
	}
	p.OnActivate(initTestPlugin(t, &config))

	post := &model.Post{
		Message: "XD",
	}
	post, err := p.MessageWillBeUpdated(post, nil)
	assert.NotNil(t, post)
	assert.Equal(t, ":laughing:", post.Message)
	assert.Equal(t, "", err)
}

func TestMappingsPrecedenceWithoutSlack(t *testing.T) {
	p := Emoticon2EmojiPlugin{}
	config := Emoticon2EmojiPluginConfiguration{
		MatchesChoice: "default_custom",
		UserMatches:   "{\"XD\":\"cry\"}",
	}
	p.OnActivate(initTestPlugin(t, &config))

	post := &model.Post{
		Message: "</3",
	}
	post, err := p.MessageWillBeUpdated(post, nil)
	assert.NotNil(t, post)
	assert.Equal(t, "</3", post.Message)
	assert.Equal(t, "", err)
}

func TestMappingsPrecedenceWithoutDefault(t *testing.T) {
	p := Emoticon2EmojiPlugin{}
	config := Emoticon2EmojiPluginConfiguration{
		MatchesChoice: "slack_custom",
		UserMatches:   "{\"XD\":\"cry\"}",
	}
	p.OnActivate(initTestPlugin(t, &config))

	post := &model.Post{
		Message: "8D",
	}
	post, err := p.MessageWillBeUpdated(post, nil)
	assert.NotNil(t, post)
	assert.Equal(t, "8D", post.Message)
	assert.Equal(t, "", err)
}

func TestUnserializeMatches(t *testing.T) {
	result, err := unserializeConfigMatches("{\";)\": \"broken_heart\", \"^^\\\"\":\"hehe\"}")
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "broken_heart", result[";)"])
	assert.Equal(t, "hehe", result["^^\""])
}

func TestOnLoadConfigurationError(t *testing.T) {
	p := Emoticon2EmojiPlugin{}
	api := &plugintest.API{}
	apiLoadConfigError := &model.AppError{
		Message: "argh",
	}

	api.On("LoadPluginConfiguration", mock.Anything).Return(apiLoadConfigError)
	err := p.OnActivate(api)
	assert.NotNil(t, err)
	assert.Contains(t, err, apiLoadConfigError.Message)
}
