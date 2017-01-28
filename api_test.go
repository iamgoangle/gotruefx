package truefx

import (
	"encoding/json"
	"testing"
)

func TestGet(t *testing.T) {
	var ticks []Tick
	feed := NewFeed()
	result := feed.Get()
	err := json.Unmarshal(result, &ticks)
	if err != nil {
		t.Fatal(err)
	}
	if len(ticks) < 1 {
		t.Error("The returned result is incorrect.")
	}
}

func TestGetBySymbol(t *testing.T) {
	var ticks []Tick
	feed := NewFeed()
	result := feed.GetBy("EUR/USD")
	err := json.Unmarshal(result, &ticks)
	if err != nil {
		t.Fatal(err)
	}
	if len(ticks) != 1 {
		t.Error("The returned result is incorrect.")
	}
}

func TestGetBySymbolWithAuthorizedSession(t *testing.T) {
	var ticks []Tick
	feed, err := NewFeedAuthorize("user", "passwd")
	if err != nil {
		t.Fatal(err)
	}
	result := feed.GetBy("AUD/JPY")
	err = json.Unmarshal(result, &ticks)
	if err != nil {
		t.Fatal(err)
	}
	if len(ticks) == 0 {
		t.Log(feed.uri)
		t.Error("The session is incorrect.")
	}
}
