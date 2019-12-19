<h1 align="center">Welcome to go-soundcloud 👋 *** Work In Progress ***</h1>
<p align="center">
  <a href="https://github.com/tsirysndr/go-soundcloud/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-soundcloud.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-soundcloud">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-soundcloud">
  <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/go-soundcloud">
  <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-soundcloud">
  <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-soundcloud">
  <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-soundcloud">
  <a href="https://github.com/tsirysndr/go-soundcloud/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>

> go-soundcloud is a Go client library for accessing the [SoundCloud API](https://developers.soundcloud.com/docs/api/reference)

## 🚚 Install

```sh
go get github.com/tsirysndr/go-soundcloud
```

## 🚀 Usage

Import the package into your project.

```Go
import "github.com/tsirysndr/go-soundcloud"
```

Construct a new SoundCloud client, then use the various services on the client to access different parts of the SoundCloud API. For example:

```Go
client := soundcloud.NewClient()
track, _ := client.Track.Get("13158665")
```

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
