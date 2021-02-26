package main

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initTestPlugin(t *testing.T, config *configuration) *plugintest.API {
	api := &plugintest.API{}
	api.On("LoadPluginConfiguration", mock.AnythingOfType("*main.configuration")).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*configuration)
		dest.CustomMatches = config.CustomMatches
	}).Return(func(dest interface{}) error {
		*dest.(*configuration) = *config
		return nil
	})
	return api
}

func TestNewPostWithDefaultEmoticon(t *testing.T) {
	p := Plugin{}
	config := configuration{
		CustomMatches: "{\"XD\":\"laughing\"}",
	}
	p.API = initTestPlugin(t, &config)
	_ = p.OnConfigurationChange()

	post := &model.Post{
		Message: "Hello XD !!",
	}
	post, err := p.MessageWillBePosted(nil, post)
	assert.NotNil(t, post)
	assert.Equal(t, "Hello :laughing: !!", post.Message)
	assert.Equal(t, "", err)
}

func TestUpdatedPostWithSlackEmoticon(t *testing.T) {
	p := Plugin{}
	config := configuration{}
	p.API = initTestPlugin(t, &config)
	_ = p.OnConfigurationChange()

	post := &model.Post{
		Message: "Hello </3 !!",
	}
	post, err := p.MessageWillBeUpdated(nil, post, nil)
	assert.NotNil(t, post)
	assert.Equal(t, "Hello :broken_heart: !!", post.Message)
	assert.Equal(t, "", err)
}

func TestReplaceWithOnlyEmoticon(t *testing.T) {
	var testMatches = map[string]match{
		"XD": {
			replacement: "laughing",
			regexp:      getEmoticonRegexp("XD"),
		},
		"</3": {
			replacement: "broken_heart",
			regexp:      getEmoticonRegexp("</3"),
		},
		"8)": {
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
	p := Plugin{}
	config := configuration{
		CustomMatches: "{\"8)\":\"cry\"}",
	}
	p.API = initTestPlugin(t, &config)
	_ = p.OnConfigurationChange()

	post := &model.Post{
		Message: "8)",
	}
	post, err := p.MessageWillBeUpdated(nil, post, nil)
	assert.NotNil(t, post)
	assert.Equal(t, ":cry:", post.Message)
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
