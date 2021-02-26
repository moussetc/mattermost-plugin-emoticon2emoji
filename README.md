# Mattermost plugin Emoticon2Emoji [![Build status](https://travis-ci.org/moussetc/mattermost-plugin-emoticon2emoji.svg?branch=master)](https://travis-ci.com/github/moussetc/mattermost-plugin-emoticon2emoji)

**Maintainer:** [@moussetc](https://github.com/moussetc)

A Mattermost plugin that completes the automatic conversion from emoticon to emoji (ex: XD to :laughing:). The plugin completes the existing mappings with the [Slack mappings](https://get.slack.help/hc/en-us/articles/202931348-Use-emoji-and-emoticons#use-emoticons) and some others (see `srv/matches.go` for the default list). You can also configurer the mappings of your choice.

## Usage

Just use any emoticon from the mappings list, it will converted into the corresponding emoji when the message is posted.

## Compatibility
Use the following table to find the correct plugin version for each Mattermost server version:

| Mattermost server | Plugin release | Incompatibility |
| --- | --- | --- |
| 5.20 and higher | v2.1.x and higher | breaking plugin manifest change |
| 5.13 to 5.19 | *not implemented*  | breaking plugin API change |
| 5.2 to 5.12 | v2.1.x |  |
| 5.0| v1.0.0 | |
| below | *not supported* |  plugins can't intercept messages |

## Installation and configuration
1. Go to the [Releases page](https://github.com/moussetc/mattermost-plugin-emoticon2emoji/releases) and download the `.tar.gz` package. Supported platforms are: Linux x64, Windows x64, Darwin x64, FreeBSD x64.
2. Use the Mattermost `System Console > Plugins Management > Management` page to upload the `.tar.gz` package
3. If you wish to edit the default mappings, go to the plugin configuration page
4. **Activate the plugin** in the `System Console > Plugins Management > Management` page

### Configuration Notes in HA

If you are running Mattermost v5.11 or earlier in [High Availability mode](https://docs.mattermost.com/deployment/cluster.html), please review the following:

1. To install the plugin, [use these documented steps](https://docs.mattermost.com/administration/plugins.html#plugin-uploads-in-high-availability-mode)
2. Then, modify the config.json [using the standard doc steps](https://docs.mattermost.com/deployment/cluster.html#updating-configuration-changes-while-operating-continuously) to the following (check the [plugin.json](https://github.com/moussetc/mattermost-plugin-emoticon2emoji/blob/master/plugin.json) file to see the default mappins).

```json
 "PluginSettings": {
        // [...]
        "Plugins": {
            "com.github.moussetc.mattermost.plugin.emoticon2emoji": {
                "CustomMatches": ""
        },
        "PluginStates": {
            // [...]
            "com.github.moussetc.mattermost.plugin.emoticon2emoji": {
                "Enable": true
            },
        }
    }
```

## Development

To avoid having to manually install your plugin, build and deploy your plugin using one of the following options.

### Deploying with Local Mode

If your Mattermost server is running locally, you can enable [local mode](https://docs.mattermost.com/administration/mmctl-cli-tool.html#local-mode) to streamline deploying your plugin. Edit your server configuration as follows:

```json
{
    "ServiceSettings": {
        ...
        "EnableLocalMode": true,
        "LocalModeSocketLocation": "/var/tmp/mattermost_local.socket"
    }
}
```

and then deploy your plugin:
```
make deploy
```

You may also customize the Unix socket path:
```
export MM_LOCALSOCKETPATH=/var/tmp/alternate_local.socket
make deploy
```

If developing a plugin with a webapp, watch for changes and deploy those automatically:
```
export MM_SERVICESETTINGS_SITEURL=http://localhost:8065
export MM_ADMIN_TOKEN=j44acwd8obn78cdcx7koid4jkr
make watch
```

### Deploying with credentials

Alternatively, you can authenticate with the server's API with credentials:
```
export MM_SERVICESETTINGS_SITEURL=http://localhost:8065
export MM_ADMIN_USERNAME=admin
export MM_ADMIN_PASSWORD=password
make deploy
```

or with a [personal access token](https://docs.mattermost.com/developer/personal-access-tokens.html):
```
export MM_SERVICESETTINGS_SITEURL=http://localhost:8065
export MM_ADMIN_TOKEN=j44acwd8obn78cdcx7koid4jkr
make deploy
```