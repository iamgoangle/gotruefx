package truefx

import (
	"testing"
)

func TestGetAll(t *testing.T) {
	feed := NewFeed()
	ticks := feed.Get()
	if len(ticks) < 1 {
		t.Error("Cannot get")
	}
}

func TestGetBySymbol(t *testing.T) {
	feed := NewFeed()
	ticks := feed.GetBySymbol("eur/usd")
	if len(ticks) != 1 {
		t.Error("'eur/usd' failed")
	}

	ticks = feed.GetBySymbol("eurusd")
	if len(ticks) != 1 {
		t.Error("'eurusd' failed")
	}

	ticks = feed.GetBySymbol("eurusd,gbpusd")
	if len(ticks) != 2 {
		t.Error("'euruse,gbpusd' failed")
	}

	ticks = feed.GetBySymbol("EUR/USD,GBP/USD")
	if len(ticks) != 2 {
		t.Error("'EUR/USD,GBP/USD' failed")
	}
}

func TestGetBySymbolWithBypassAuthorizedSession(t *testing.T) {
	feed := NewFeedBypass("admin")
	ticks := feed.GetBySymbol("AUD/JPY")
	if len(ticks) == 0 {
		t.Log(feed.uri)
		t.Error("The session is incorrect.")
	}
}
