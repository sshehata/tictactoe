package gym

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"tictactoe/game"
)

// Env a tictacoe gym env
type Env struct {
	baseURL   *url.URL
	sessionID int
}

func startNewGame(u *url.URL) int {
	resp, err := http.Get(u.String())
	if err != nil {
		panic(fmt.Sprintf("could not start new game: %v", err.Error()))
	}
	body := parseBody(resp)
	return int(body["sessionID"].(float64))
}

// NewEnv Create new gym tictactoe env
func NewEnv(base string) (*Env, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("invalid base url: %v", err.Error()))
	}

	sessionID := startNewGame(baseURL)
	return &Env{
		baseURL,
		sessionID,
	}, nil
}

func (e *Env) getURL(path string) *url.URL {
	u, err := url.Parse(fmt.Sprintf("%v/%v", e.sessionID, path))
	if err != nil {
		panic(fmt.Sprintf("bad urls: %v", err.Error()))
	}
	return e.baseURL.ResolveReference(u)
}

func parseBody(resp *http.Response) map[string]interface{} {
	var body map[string]interface{}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("request to server failed: %v", resp.StatusCode))
	}

	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(fmt.Sprintf("could not parse response: %v", err.Error()))
	}

	return body
}

// GameOver Returns true if the game is over
func (e *Env) GameOver() bool {
	u := e.getURL("gameover")

	resp, err := http.Get(u.String())
	if err != nil {
		panic(fmt.Sprintf("failed to contact game server: %v", err.Error()))
	}

	body := parseBody(resp)
	return body["gameOver"].(bool)
}

// Observation Returns current board state
func (e *Env) Observation() Board {
	u := e.getURL("")

	resp, err := http.Get(u.String())
	if err != nil {
		panic(fmt.Sprintf("failed to contact game server: %v", err.Error()))
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("request to server failed: %v", resp.StatusCode))
	}

	var body map[string]Board
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(fmt.Sprintf("could not parse response: %v", err.Error()))
	}

	return body["board"]
}

// Moves Returns possible moves according to current state
func (e *Env) Moves() []game.Position {
	u := e.getURL("moves")

	resp, err := http.Get(u.String())
	if err != nil {
		panic(fmt.Sprintf("failed to contact game server: %v", err.Error()))
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("request to server failed: %v", resp.StatusCode))
	}

	var body map[string][]game.Position
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(fmt.Sprintf("could not parse response: %v", err.Error()))
	}

	return body["moves"]
}
