# Flocks Of Blocks

<a href="https://github.com/jeremyforan/go-flocks-of-blocks/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/jeremyforan/go-flocks-of-blocks">
</a>

<a href="https://pkg.go.dev/github.com/jeremyforan/go-flocks-of-blocks"><img src="https://pkg.go.dev/badge/github.com/jeremyforan/go-flocks-of-blocks.svg" alt="Go Reference"></a>

Flocks of Blocks is a Go library that helps compose Slack messages using the [Block Framework](https://api.slack.com/block-kit). 

:warning: Warning
------------------------------------------
This package is in heavy flux at the moment as I work to incorporate feedback from various sources.


# Why

After building Slack bots in Go, I looked for a faster way of composing block messages. I started utilizing the Go templating library to build ad-hoc messages. This package provides an intuitive way of generating Slack messages using templates and a functional approach heavily inspired by the Charm


# Philosophy
Slack messages should be easy and fun to compose. Most Slack messages are simple and, as a result, less likely to violate any of Slack message's restrictions:
> 	IE: Maximum length for this field is 255 characters.

Therefore the package should try and minimize the amount of errors the user has to handle.

A functional approach to building the assets will be more concise and intuitive:

```go
button := NewButton("Click This", "button1").MakeStyleDanger().AddUrl(url)
```

# Roadmap
There are three major phases to this project:

### 1) Version 0.0.x (Develope) 
* Explore different implementation and design approaches including more composition
* Gather user feedback

### 2) Version 1.0.x (Productionized)

* Establish a consistant naming convention
* Decide on the degree of structure composition
* Complete code coverage

### 3) Version 1.1.x (Maintain)
* Add component validation
* Make a really great Logo. Like sticker worthy!

# Documentation

## Installation

## Usage
	
	
# Notes
the slack reference page is missing a type for the Video block reference https://api.slack.com/reference/block-kit/blocks#video

Extrnal data sources missing field options
https://api.slack.com/reference/block-kit/block-elements#external_multi_select

