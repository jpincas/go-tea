package main

import (
	"encoding/json"
	"os"
	"sort"
	"sync"
	"time"
)

const leaderboardFile = "leaderboard.json"

type HighScore struct {
	Initials string    `json:"initials"`
	Score    int       `json:"score"`
	Date     time.Time `json:"date"`
}

type Leaderboard struct {
	Easy   []HighScore `json:"easy"`
	Medium []HighScore `json:"medium"`
	Hard   []HighScore `json:"hard"`
	mu     sync.Mutex
}

func NewLeaderboard() *Leaderboard {
	return &Leaderboard{
		Easy:   []HighScore{},
		Medium: []HighScore{},
		Hard:   []HighScore{},
	}
}

func LoadLeaderboard() (*Leaderboard, error) {
	l := NewLeaderboard()
	data, err := os.ReadFile(leaderboardFile)
	if err != nil {
		if os.IsNotExist(err) {
			return l, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Leaderboard) Save() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(leaderboardFile, data, 0644)
}

func (l *Leaderboard) AddScore(difficulty Difficulty, score HighScore) {
	l.mu.Lock()
	defer l.mu.Unlock()

	switch difficulty {
	case Easy:
		l.Easy = append(l.Easy, score)
		sortScores(l.Easy)
		if len(l.Easy) > 10 {
			l.Easy = l.Easy[:10]
		}
	case Medium:
		l.Medium = append(l.Medium, score)
		sortScores(l.Medium)
		if len(l.Medium) > 10 {
			l.Medium = l.Medium[:10]
		}
	case Hard:
		l.Hard = append(l.Hard, score)
		sortScores(l.Hard)
		if len(l.Hard) > 10 {
			l.Hard = l.Hard[:10]
		}
	}
}

func (l *Leaderboard) IsHighScore(difficulty Difficulty, score int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	var scores []HighScore
	switch difficulty {
	case Easy:
		scores = l.Easy
	case Medium:
		scores = l.Medium
	case Hard:
		scores = l.Hard
	}

	if len(scores) < 10 {
		return true
	}

	// Check if score is better (lower) than the worst score on the leaderboard
	return score < scores[len(scores)-1].Score
}

func sortScores(scores []HighScore) {
	sort.Slice(scores, func(i, j int) bool {
		// Lower score (turns) is better
		if scores[i].Score != scores[j].Score {
			return scores[i].Score < scores[j].Score
		}
		// If scores are equal, newer date is better (or older? usually first to achieve keeps rank)
		// Let's say first to achieve keeps rank, so older date is "better" (smaller index)
		return scores[i].Date.Before(scores[j].Date)
	})
}

func (l *Leaderboard) GetScores(difficulty Difficulty) []HighScore {
	l.mu.Lock()
	defer l.mu.Unlock()

	switch difficulty {
	case Easy:
		return l.Easy
	case Medium:
		return l.Medium
	case Hard:
		return l.Hard
	}
	return []HighScore{}
}
