# mattermost-plugin-emoticon2emoji [![Build Status](https://api.travis-ci.com/moussetc/mattermost-plugin-emoticon2emoji.svg?branch=master)](https://travis-ci.com/moussetc/mattermost-plugin-emoticon2emoji)
A plugin that completes the automatic conversion from emoticon to emoji (:) to :smile:) in Mattermost messages, by adding [Slack mappings](https://get.slack.help/hc/en-us/articles/202931348-Use-emoji-and-emoticons#use-emoticons) and some other (see `matches.go` for the default list) which can be configured.

## Compatibility
- for Mattermost 5.2 or higher: use v2.x.x release
- for Mattermost 5.0: use v1.0.0 release
- for Mattermost below: unsupported versions (plugins can't intercept posts)


## Installation and configuration
1. Go to the [Releases page](https://github.com/moussetc/mattermost-plugin-emoticon2emoji/releases) and download the package for your OS and architecture.
2. Use the Mattermost `System Console > Plugins Management > Management` page to upload the `.tar.gz` package
3. Go to the plugin configuration page to edit the custom mappings if needed
4. **Activate the plugin** in the `System Console > Plugins Management > Management` page

## Manual configuration
If you need to enable & configure this plugin directly in the Mattermost configuration file `config.json`, for example if you are doing a [High Availability setup](https://docs.mattermost.com/deployment/cluster.html), you can use the following lines:
```json
 "PluginSettings": {
        // [...]
        "Plugins": {
            "com.github.moussetc.mattermost.plugin.emoticon2emoji": {
                "CustomMatches": "" // custom emoticons->emoji mappings in JSON format, see plugin.yaml for the default value
        },
        "PluginStates": {
            // [...]
            "com.github.moussetc.mattermost.plugin.emoticon2emoji": {
                "Enable": true
            },
        }
    }
```

## Usage
Use the usual emoticons and the post will be automatically updated to replace the emoticon by the emoji code.

## Development
Run make vendor to install dependencies, then develop like any other Go project, using go test, go build, etc.

If you want to create a fully bundled plugin that will run on a local server, you can use `make mattermost-plugin-emoticon2emoji.tar.gz`.
