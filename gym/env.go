package gym

import (
	"bytes"
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

func (e *Env) resolve(path string) *url.URL {
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

// Done Returns true if the game is over
func (e *Env) Done() bool {
	u := e.resolve("gameover")

	resp, err := http.Get(u.String())
	if err != nil {
		panic(fmt.Sprintf("failed to contact game server: %v", err.Error()))
	}

	body := parseBody(resp)
	return body["gameOver"].(bool)
}

// Observation Returns current board state
func (e *Env) Observation() Board {
	u := e.resolve("")

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

// ActionSpace Returns possible moves according to current state
func (e *Env) ActionSpace() []game.Position {
	u := e.resolve("moves")

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

// Step perform the chosen action at the current step
func (e *Env) Step(action game.Position) (Board, bool) {
	u := e.resolve("")

	data, _ := json.Marshal(action)
	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(data))

	if err != nil {
		panic(fmt.Sprintf("failed to contact game server: %v", err.Error()))
	}

	var body map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(fmt.Sprintf("could not parse response: %v", err.Error()))
	}

	return *MakeBoard(body["board"].([]interface{})), body["status"].(bool)
}

// Reward get the reward for the particular player
func (e *Env) Reward(t game.Tile) float64 {
	u := e.resolve("winner")
	resp, err := http.Get(u.String())

	if err != nil {
		panic(fmt.Sprintf("failed to contact game server: %v", err.Error()))
	}

	var body map[string]game.Tile
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(fmt.Sprintf("could not parse response: %v", err.Error()))
	}

	if body["winner"] == game.Undefined {
		return 0.1
	} else if body["winner"] == t {
		return 1
	} else {
		return 0
	}
}

// Reset reset the game
func (e *Env) Reset() Board {
	u := e.resolve("reset")
	resp, err := http.Post(u.String(), "application/json", nil)

	if err != nil {
		panic(fmt.Sprintf("failed to contact game server: %v", err.Error()))
	}

	var body map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		panic(fmt.Sprintf("could not parse response: %v", err.Error()))
	}

	return *MakeBoard(body["board"].([]interface{}))
}
