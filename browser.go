/*
Package browser provides functions for opening a URL with the user's default browser.

If the operating system is darwin (macOS), it will use open.

If the operating system is Windows, it will use start.

If the operating system is Linux / FreeBSD, it will try these strategies (in this order): xdg-open, sensible-browser, x-www-browser.
 */
package browser

import (
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

/*
Open takes a string and tries to open the browser with the string.

The function will try to parse the URL and pass it into OpenURL.

See https://golang.org/pkg/net/url/#Parse for details.
 */
func Open(rawURLStr string) error {

	urlObj, err := url.Parse(rawURLStr)

	if err != nil {
		return err
	}

	return OpenURL(urlObj)
}

/*
OpenURL takes a url struct instead of a raw string.

If the scheme is not defined (empty scheme), the function will automatically append "https://" as the scheme.
 */
func OpenURL(urlObj *url.URL) error {

	// append scheme if it does not exist.

	if !urlObj.IsAbs() {
		urlObj.Scheme = "https"
	}

	switch runtime.GOOS {
	case "darwin":
		return handleDarwin(urlObj)
	case "windows":
		return handleWindows(urlObj)
	case "linux", "freebsd":
		return handleNix(urlObj)
	default:
		return errors.New(fmt.Sprintf("OS type %v not implemented.", runtime.GOOS))
	}

}

func handleDarwin(url *url.URL) error {
	// use `open`
	cmd := exec.Command("open", url.String())
	return cmd.Run()
}

func handleWindows(url *url.URL) error {
	// use `start`
	cmd := exec.Command("start", url.String())
	return cmd.Run()
}

func handleNix(url *url.URL) error {

	strategies := []string{"xdg-open", "sensible-browser", "x-www-browser"}

	var strategy string

	// use `which` to determine if the command exists.
	// which returns 1 if command is not found.

	for _, potential := range strategies {
		cmd := exec.Command("which", potential)
		if err := cmd.Run(); err != nil {
			continue // next iteration
		} else {
			strategy = potential
			break
		}
	}

	if strategy == "" {
		return errors.New("no strategy available.")
	}

	cmd := exec.Command(strategy, url.String())

	return cmd.Run()
}
