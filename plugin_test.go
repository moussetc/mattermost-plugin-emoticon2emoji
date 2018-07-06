package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mattermost/mattermost-server/model"

	"github.com/mattermost/mattermost-server/plugin/plugintest"
	"github.com/mattermost/mattermost-server/plugin/plugintest/mock"
)

func initTestPlugin(t *testing.T) *plugintest.API {

	api := &plugintest.API{}
	api.On("RegisterCommand", mock.Anything).Return(nil)
	api.On("UnregisterCommand", mock.Anything, mock.Anything).Return(nil)
	api.On("GetUser", mock.Anything).Return(&model.User{
		Id:       "userid",
		Nickname: "User",
	}, (*model.AppError)(nil))

	return api
}

func TestUnserializeMatches(t *testing.T) {
	p := Emoticon2EmojiPlugin{}

	result, err := p.unserializeMatches("{\";)\": \"wink\", \"^^\\\"\":\"hehe\"}")
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "wink", result[";)"])
	assert.Equal(t, "hehe", result["^^\""])
}
