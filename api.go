package gotruefx

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Tick struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Bid       float64 `json:"bid"`
	Offer     float64 `json:"offer"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	Spread    float64 `json:"spread"`
}

type Feed struct {
	session string
	baseURI string
	uri     string
}

func (f *Feed) Get() []Tick {
	return f.fetch(f.baseURI)
}

func (f *Feed) GetBySymbol(s string) []Tick {
	if len(s) < 6 {
		return []Tick{}
	}
	var symbol string
	var symbols string
	for i, s := range strings.Split(s, ",") {
		symbol = s
		if !strings.Contains(s, "/") {
			symbol = s[:3] + "/" + s[3:]
		}
		if i > 0 {
			symbols += "," + symbol
		} else {
			symbols = symbol
		}
	}
	uri := f.baseURI + "&c=" + strings.ToUpper(symbols)
	return f.fetch(uri)
}

func (f *Feed) fetch(uri string) []Tick {
	var ticks []Tick
	var t Tick

	if strings.Compare(f.session, "") != 0 {
		uri = uri + "&id=" + f.session
		f.uri = uri
	}

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(body))
	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range record {
		t.Symbol = v[0]
		t.Timestamp, _ = strconv.ParseInt(v[1], 10, 64)
		t.Bid, _ = strconv.ParseFloat(v[2]+v[3], 64)
		t.Offer, _ = strconv.ParseFloat(v[4]+v[5], 64)
		t.Low, _ = strconv.ParseFloat(v[6], 64)
		t.High, _ = strconv.ParseFloat(v[7], 64)
		t.Open, _ = strconv.ParseFloat(v[8], 64)
		t.Spread = calcSpread(t)
		ticks = append(ticks, t)
	}
	return ticks
}

func calcSpread(t Tick) float64 {
	if strings.Contains(t.Symbol, "JPY") {
		return round((t.Offer-t.Bid)*100, 1)
	} else {
		return round((t.Offer-t.Bid)*10000, 1)
	}
}

func round(num float64, precision uint) float64 {
	pow := math.Pow(10, float64(precision))
	n := int((num * pow) + math.Copysign(0.5, num))
	return float64(n) / pow
}

func NewFeed() *Feed {
	var feed Feed
	feed.baseURI = "https://webrates.truefx.com/rates/connect.html?f=csv"
	return &feed
}

func NewFeedAuthorized(username string, password string) (*Feed, error) {
	feed := NewFeed()
	uri := feed.baseURI + "&u=" + username + "&p=" + password + "&q=fx"
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if bytes.Compare(bytes.TrimSpace(body), []byte("not authorized")) == 0 {
		err = errors.New("Not Authorized")
	} else {
		feed.session = string(bytes.TrimSpace(body))
	}
	return feed, err
}

func NewFeedBypass(username string) *Feed {
	feed := NewFeed()
	feed.session = username + ":p:q:0"
	return feed
}
