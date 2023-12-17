package main

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Player struct {
	Name              string        `json:"Name"`
	RemainingTime     string        `json:"RemainingTime"`
	RemainingTimeReal time.Duration `json:"RemainingTimeReal"`
	BonusTime         time.Duration `json:"BonusTime"`
	Turn              bool          `json:"Turn"`
}

type Game struct {
	Active          bool          `json:"Active"`
	StartTime       time.Time     `json:"StartTime"`
	TotalDuration   time.Duration `json:"TotalDuration"`
	CurrentDuration time.Duration `json:"CurrentDuration"`
	LastChange      time.Time     `json:"LastChange"`
}

var player1 *Player
var player2 *Player
var game *Game

func InstantiatePlayers(ctx context.Context) {
	player1 = &Player{
		Name:              "Player 1",
		RemainingTime:     "5m0s",
		RemainingTimeReal: time.Duration(5 * 60 * time.Second),
		BonusTime:         time.Duration(0),
		Turn:              true,
	}
	player2 = &Player{
		Name:              "Player 2",
		RemainingTime:     "5m0s",
		RemainingTimeReal: time.Duration(5 * 60 * time.Second),
		BonusTime:         time.Duration(0),
		Turn:              false,
	}
	game = &Game{
		Active:          false,
		StartTime:       time.Now(),
		TotalDuration:   time.Duration(0),
		CurrentDuration: time.Duration(0),
		LastChange:      time.Now(),
	}
	sendGameData(ctx)
}

type GameData struct {
	Players [2]*Player `json:"Players"`
	Game    *Game      `json:"Game"`
}

func sendGameData(ctx context.Context) error {
	go func() {
		for {
			game_data := &GameData{
				Players: [2]*Player{player1, player2},
				Game:    game,
			}
			runtime.EventsEmit(ctx, "game_data", game_data)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	go func() {
		for {
			if game.Active {
				game.CurrentDuration = game.TotalDuration + time.Since(game.StartTime)
				for _, player := range [2]*Player{player1, player2} {
					if player.Turn {
						new_time_remaining := player.RemainingTimeReal - time.Since(game.LastChange)
						player.RemainingTime = new_time_remaining.Round(1 * time.Second).String()
					}
				}
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return nil
}

func saveRealTime(p *Player) {
	new_time_remaining := p.RemainingTimeReal - time.Since(game.LastChange)
	if new_time_remaining <= 0 {
		p.RemainingTimeReal = 0
	}
	p.RemainingTimeReal = new_time_remaining
}

func (g *GameData) StartStopGame() {
	if game.Active {
		game.Active = false
	} else {
		game.Active = true
		game.StartTime = time.Now()
		game.LastChange = time.Now()
	}
}

func (g *GameData) EndRound() {
	if player1.Turn {
		player1.Turn = false
		player2.Turn = true
		saveRealTime(player1)
	} else {
		player1.Turn = true
		player2.Turn = false
		saveRealTime(player2)
	}
	game.LastChange = time.Now()
}
