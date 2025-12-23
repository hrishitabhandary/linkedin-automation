package auth

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFile = "data/cookies.json"

// SaveCookies saves browser cookies to disk
func SaveCookies(page *rod.Page) error {
	cookies := page.MustCookies()

	file, err := os.Create(cookieFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(cookies)
}

// LoadCookies loads cookies from disk into browser
func LoadCookies(page *rod.Page) error {
	file, err := os.Open(cookieFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var cookies []*proto.NetworkCookie
	if err := json.NewDecoder(file).Decode(&cookies); err != nil {
		return err
	}

	// Convert to NetworkCookieParam (minimal fields)
	var cookieParams []*proto.NetworkCookieParam
	for _, c := range cookies {
		cookieParams = append(cookieParams, &proto.NetworkCookieParam{
			Name:   c.Name,
			Value:  c.Value,
			Domain: c.Domain,
			Path:   c.Path,
			Expires: c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure: c.Secure,
			SameSite: c.SameSite,
		})
	}

	page.MustSetCookies(cookieParams...)
	return nil
}
