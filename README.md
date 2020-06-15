# open-browser

_(untested. use at your own discretion. PRs welcomed.)_

This package provides functions for opening a URL with the user's default browser. (or a mail client.)

|OS|Strategies|
|---|---|
|macOS (darwin)|`open`|
|Windows|`start`|
|Linux/FreeBSD|`xdg-open`, `sensible-browser`, `x-www-browser`|

Note: If the scheme is not defined (empty scheme), the function will automatically append "https://" as the scheme.

```go
// example

package main

import "github.com/lexffe/go-open-browser"

func main() {
    if err := browser.Open("https://www.github.com"); err != nil {
        // handle error...
    }
}
```
