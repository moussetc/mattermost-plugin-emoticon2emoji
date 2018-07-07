# mattermost-plugin-emoticon2emoji
A plugin that completes the automatic conversion from emoticon to emoji (:) to :smile:) in Mattermost messages, by adding [Slack mappings](https://get.slack.help/hc/en-us/articles/202931348-Use-emoji-and-emoticons#use-emoticons) and some other (see `matches.go` for the full list) and permitting the administrator to define custom mappings.

## Requirements
- Mattermost 5.0 (to allow plugins to intercept posts).

## Installation and configuration
1. Go to the [Releases page](https://github.com/moussetc/mattermost-plugin-emoticon2emoji/releases) and download the package for your OS and architecture.
2. Use the Mattermost `System Console > Plugins Management > Management` page to upload the `.tar.gz` package
3. Go to the plugin configuration page and follow the instructions
4. **Activate the plugin** in the `System Console > Plugins Management > Management` page

## Manual configuration
If you need to enable & configure this plugin directly in the Mattermost configuration file `config.json`, for example if you are doing a [High Availability setup](https://docs.mattermost.com/deployment/cluster.html), you can use the following lines (remember to set the API key!):
```json
 "PluginSettings": {
        // [...]
        "Plugins": {
            "com.github.moussetc.mattermost.plugin.emoticon2emoji": {
                "MatchesChoice": "slack_default_custom", // other choices are combinations of slack, default and custom
                "UserMatches": "{}" // custom emoticons->emoji mappings in JSON format, for example "{\":)\":\"grin\", \":(\": \"cry\"}"
            },
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