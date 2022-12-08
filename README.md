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


# Installation

```bash
go get github.com/jeremyforan/go-flocks-of-blocks
```

## Usage

### Basic Usage

Build the [Plain Text](https://app.slack.com/block-kit-builder/#%7B%22blocks%22:%5B%7B%22type%22:%22section%22,%22text%22:%7B%22type%22:%22plain_text%22,%22text%22:%22This%20is%20a%20plain%20text%20section%20block.%22,%22emoji%22:true%7D%7D%5D%7D) example

```go
package main

import (
	"fmt"
	fobs "github.com/jeremyforan/go-flocks-of-blocks"
)

func main() {
	text := fobs.NewPlainText("This is a plain text section block.").EnableEmoji()
	section := fobs.NewSection().SetText(text)
	basic := fobs.NewMessage().AddBlock(section)

	fmt.Println(basic)
```
Outputs:

```json
{
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "plain_text",
				"text": "This is a plain text section block.",
				"emoji": true
			}
		}
	]
}
```

This can be condensed

```go

package main

import (
	"fmt"
	fobs "github.com/jeremyforan/go-flocks-of-blocks"
)

func main() {
	text := fobs.NewPlainText("This is a plain text section block.").EnableEmoji()
	fmt.Println(fobs.NewMessage().AddBlock(text.Section()))
}
```

### Real World Example

Here is an Example 

![real world example](https://github.com/jeremyforan/go-flocks-of-blocks/blob/master/assets/real-world-1.png?raw=true)

To build that you would write

```go
package main

import (
	"fmt"
	fobs "github.com/jeremyforan/go-flocks-of-blocks"
)

func main() {

	// create a new message
	msg := fobs.NewMessage()

	// Add a header
	header := fobs.NewHeader("Device Summary")
	msg = msg.AddBlock(header)

	// Add some info
	info := fobs.NewSection().AddMrkdownField("*IP:* 192.168.0.1").AddMrkdownField("*Area:* basement")
	msg = msg.AddBlock(info)

	// Add some more info but in a single line
	msg = msg.AddBlock(fobs.NewSection().AddMrkdownField("*Hardware:* Linksys WRT-54G").AddMrkdownField("*Uptime:* 7 Days, 3 Months"))

	// Add the info message to
	ssid := fobs.NewSection().AddMrkdownField("*SSID:* FreshPrinceOfDonair")
	msg = msg.AddBlock(ssid)

	// make a "reset" button
	resetButton := fobs.NewButton("Reboot Device", "actionId-0").SetValue("click_me_123")

	// Let's assume we want to add a note based on some arbitrary bool value
	rfIssue := true
	if rfIssue {
		note := fobs.NewPlainText("*high levels of RF interference detected consider scan")
		msg = msg.AddBlock(note.Context())

		// We want to add the Danger styleing to the button due to the 'issue'
		resetButton = resetButton.MakeStyleDanger()
	}

	// Add the reset button to the message
	msg = msg.AddBlock(resetButton.Actions())

	// Generate a link that paste the body into the Slack interactive Block Kit Builder for validation
	fmt.Println(msg.GenerateKitBuilderUrl())
}
```

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
