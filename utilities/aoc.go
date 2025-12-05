package utilities

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadAOCURL(url string) ([]byte, error) {
	session, err := getAOCSession()
	if err != nil {
		return nil, fmt.Errorf("get aoc session: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:     "session",
		Value:    session,
		Path:     "/",
		Domain:   ".adventofcode.com",
		Secure:   true,
		HttpOnly: true,
	})

	client := http.DefaultClient
	resp, reqErr := client.Do(req)
	if reqErr != nil {
		return nil, fmt.Errorf("send http request: %w", reqErr)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status %d", resp.StatusCode)
	}

	read, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("read response body: %w", readErr)
	}

	return read, nil
}

func getAOCSession() (string, error) {
	session, exists := os.LookupEnv("AOC_SESSION")
	if !exists || len(session) == 0 {
		bytes, err := os.ReadFile(".session")
		if err != nil {
			return "", fmt.Errorf("read .session file: %w", err)
		}

		session = string(bytes)
	}

	return session, nil
}
