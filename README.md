# mattermost-plugin-emoticon2emoji
A plugin that add some emoticon to emoji mappings to Mattermost to transform. For example, a plain text emoticon like XD will be automatically converted to  :laughing:. The Slack mappings that are not done by default by Mattermost are added in this plugin, plus some more.

## Requirements
- Mattermost 5.0 (to allow plugins to intercept posts).

## Installation and configuration
1. Go to the [Releases page](https://github.com/moussetc/mattermost-plugin-emoticon2emoji/releases) and download the package for your OS and architecture.
2. Use the Mattermost `System Console > Plugins Management > Management` page to upload the `.tar.gz` package
3. **Activate the plugin** in the `System Console > Plugins Management > Management` page

## Manual configuration
If you need to enable & configure this plugin directly in the Mattermost configuration file `config.json`, for example if you are doing a [High Availability setup](https://docs.mattermost.com/deployment/cluster.html), you can use the following lines (remember to set the API key!):
```json
 "PluginSettings": {
        // [...]
        "Plugins": {
            "com.github.moussetc.mattermost.plugin.emoticon2emoji": {
                "Matches": ""
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

You can define your own emoji mappings in the plugin configuration screen (it will replace entirely the default mapping).
 To see the default emoji mappings, see the `default_matches.go` file.

## Development
Run make vendor to install dependencies, then develop like any other Go project, using go test, go build, etc.

If you want to create a fully bundled plugin that will run on a local server, you can use `make mattermost-plugin-emoticon2emoji.tar.gz`.

## What's next?
- add some auto tests
- allow admin to change some mappings without replacing ALL the default mappings
- have someone check the default mappings
- option to replace emoticon by unicode symbol
- make the plugin opt-in or opt-out with a list of users
