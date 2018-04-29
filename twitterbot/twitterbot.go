package twitterbot

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
)

func ContainsHashTag(t anaconda.Tweet, h string, stripHash bool) bool {
	for _, tag := range t.Entities.Hashtags {
		if strings.Contains(h, "#") && stripHash {
			h = h[1:]
		}
		if strings.ToLower(tag.Text) == strings.ToLower(h) {
			return true
		}
	}
	return false
}

type TwitterBot struct {
	Track  []string
	api    *anaconda.TwitterApi
	log    *twitterLogger
	me     anaconda.User
	stream *anaconda.Stream
}

func (tb *TwitterBot) ShouldRetweet(t anaconda.Tweet) (bool, string) {
	// Ignore retweets and my own tweets
	if t.User.Id == tb.me.Id || t.RetweetedStatus != nil {
		return false, "my own tweet or a retweet"
	}

	if ContainsHashTag(t, tb.Track[0], true) { // Retweet if it contains the main hashtag
		return true, ""
	}

	for _, h := range tb.Track[1:] {
		if ContainsHashTag(t, h, true) && len(t.Entities.Urls) > 0 {
			return true, ""
		}
	}

	return false, "no subscribed hashtag"
}

func (tb *TwitterBot) Tweet(msg string) error {
	_, err := tb.api.PostTweet(fmt.Sprintf("%s #golang #golangbg", msg), nil)
	return err
}

func (tb *TwitterBot) Start() {
	// Start the stream
	tb.stream = tb.api.PublicStreamFilter(url.Values{
		"track":            tb.Track,
		"include_entities": []string{"true"},
	})

	tb.log.Infof("[twitterbot] listening for: %v", tb.Track)
	// Keep looping over the stream
	for v := range tb.stream.C {
		// Assertion to tweet
		t, ok := v.(anaconda.Tweet)
		if !ok {
			tb.log.Warningf("[twitterbot] received unexpected value of type %T", v)
			continue
		}
		tb.log.Infof("[twitterbot] received a tweet %d: %s", t.Id, t.Text)

		// Check if this should be retweeted
		shouldRetweet, reason := tb.ShouldRetweet(t)
		if !shouldRetweet {
			tb.log.Infof("[twitterbot] shouldn't retweet %d: %s", t.Id, reason)
			continue
		}

		// All fine, so retweet
		_, err := tb.api.Retweet(t.Id, false)
		if err != nil {
			tb.log.Errorf("[twitterbot] could not retweet %d: %v", t.Id, err)
			continue
		}
		tb.log.Infof("[twitterbot] retweeted %d", t.Id)
	}
}

func (tb *TwitterBot) Shutdown() {
	if tb.stream != nil {
		tb.stream.Stop()
	}
}

func New(log *logrus.Logger) (*TwitterBot, error) {
	logger := &twitterLogger{log}

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	t := os.Getenv("TWITTER_TRACK")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" || t == "" {
		return nil, fmt.Errorf("TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET, TWITTER_ACCESS_TOKEN, TWITTER_ACCESS_TOKEN_SECRET or TWITTER_TRACK not set")
	}

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	api.SetLogger(logger)

	me, err := api.GetSelf(nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't get self: %v", err)
	}

	return &TwitterBot{
		api:   api,
		log:   logger,
		Track: strings.Split(strings.Trim(t, " "), ","),
		me:    me,
	}, nil
}
