# mattermost-plugin-emoticon2emoji
A Mattermost plugin to transform plain text emoticons like XD to picture emojis like :laughing:

# PLUGIN STATUS : the plugin is being developed...

## Requirements
- Mattermost 5.0 (to allow plugins to intercept posts).

## Installation and configuration
1. Go to the [Releases page](https://github.com/moussetc/mattermost-plugin-emoticon2emoji/releases) and download the package for your OS and architecture.
2. Use the Mattermost `System Console > Plugins Management > Management` page to upload the `.tar.gz` package
3. **Activate the plugin** in the `System Console > Plugins Management > Management` page

## Manual configuration TODO
If you need to enable & configure this plugin directly in the Mattermost configuration file `config.json`, for example if you are doing a [High Availability setup](https://docs.mattermost.com/deployment/cluster.html), you can use the following lines (remember to set the API key!):
```json
 "PluginSettings": {
        // [...]
        "Plugins": {
            "com.github.moussetc.mattermost.plugin.emoticon2emoji": {
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
TODO : list of supported emoticons

## Development
Run make vendor to install dependencies, then develop like any other Go project, using go test, go build, etc.

If you want to create a fully bundled plugin that will run on a local server, you can use `make mattermost-plugin-emoticon2emoji.tar.gz`.

## What's next?
TODO
