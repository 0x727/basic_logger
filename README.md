# lark_logger

This is a Lark (飞书) webhook logging container for Mythic v3.0.0+. This container connects to RabbitMQ queues for logging and sends messages to a Lark webhook in real-time.

## Features

- Sends log events to Lark webhook as messages are received
- Supports: Callback, Task, Payload, Keylog, Credential, Artifact, File, Response events
- Falls back to local logging if `LARK_WEBHOOK_URL` is not set
- Logs are still written to `mythic.log` file

## Configuration

Set the `LARK_WEBHOOK_URL` environment variable with your Lark webhook URL:

```bash
export LARK_WEBHOOK_URL=https://open.feishu.cn/open-apis/bot/v2/hook/your-webhook-url
```

To create a Lark webhook:
1. Open a Lark group
2. Click "Group Settings" > " Bots"
3. Add a custom bot and copy the webhook URL

## How to install in Mythic

When it's time for you to test out your install or for another user to install your c2 profile, it's pretty simple. Within Mythic you can run the `mythic-cli` binary to install this in one of three ways:

* `sudo ./mythic-cli install github https://github.com/user/repo` to install the main branch
* `sudo ./mythic-cli install github https://github.com/user/repo branchname` to install a specific branch of that repo
* `sudo ./mythic-cli install folder /path/to/local/folder/cloned/from/github` to install from an already cloned down version of an agent repo

Now, you might be wondering _when_ should you or a user do this to properly add your profile to their Mythic instance. There's no wrong answer here, just depends on your preference. The three options are:

* Mythic is already up and going, then you can run the install script and just direct that profile's containers to start (i.e. `sudo ./mythic-cli c2 start profileName`.
* Mythic is already up and going, but you want to minimize your steps, you can just install the profile and run `sudo ./mythic-cli mythic start`. That script will first _stop_ all of your containers, then start everything back up again. This will also bring in the new profile you just installed.
* Mythic isn't running, you can install the script and just run `sudo ./mythic-cli mythic start`.
