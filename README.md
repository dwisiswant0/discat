# discat

A simple way of sending messages from the CLI output to your Discord channel with webhook.
> Actually, this is a fork version of [slackcat](https://github.com/dwisiswant0/slackcat) that I made too!

## Installation

- Download a prebuilt binary from [releases page](https://github.com/dwisiswant0/discat/releases/latest), unpack and run! or
- If you have go1.13+ compiler installed: `go get dw1.みんな/discat`.

## Configuration

**Step 1:** Get yours Discord channel webhook URL [here](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks).

**Step 2** _(optional)_**:** Set `DISCORD_WEBHOOK_URL` environment variable.
```bash
export DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/xnxx/xxx-xxx"
```

## Usage

It's very simple!

```bash
▶ echo -e "Hello,\nworld!" | discat
```

### Flags

```
Usage of discat:
  -1    Send message line-by-line
  -u string
        Discord Webhook URL
  -v    Verbose mode
```

### Workaround

The goal is to get automated alerts for interesting stuff!

```bash
▶ assetfinder twitter.com | anew | discat -u https://hooks.discord.com/services/xxx/xxx/xxx
```

The `-u` flag is optional if you've defined `DISCORD_WEBHOOK_URL` environment variable.

Discat also strips the ANSI colors from stdin to send messages, so you'll receive a clean message on your Discord!

```bash
▶ nuclei -l urls.txt -t cves/ | discat
```

![Proof](https://user-images.githubusercontent.com/25837540/108782401-1571e380-759e-11eb-8d20-dfcc9294a30a.png)

### Line-by-line

Instead of have to wait for previously executed program to finish, use the `-1` flag if you want to send messages on a line by line _(default: false)_.

```bash
▶ amass track -d domain.tld | discat -1
```

## License

`discat` is distributed under MIT License.