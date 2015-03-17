package grad

import (
	"strings"
	"testing"
)

func TestDetectFeedWithEmpty(t *testing.T) {
	html := ""
	actual := DetectFeed(strings.NewReader(html))
	var expected []string

	if len(actual) != len(expected) {
		t.Errorf("got %v\n want %v", len(actual), len(expected))
	}
}

func TestDetectFeedWithNoFeed(t *testing.T) {
	html := "<html><head></haed></html>"
	actual := DetectFeed(strings.NewReader(html))
	var expected []string

	if len(actual) != len(expected) {
		t.Errorf("got %v\n want %v", len(actual), len(expected))
	}
}

func TestDetectFeedWithOneRssFeed(t *testing.T) {
	html := "<html><head><link rel=\"alternate\" type=\"application/rss+xml\" title=\"test_rss_feed_title\" href=\"http://test_rss_feed_url/\"/></head></html>"
	actual := DetectFeed(strings.NewReader(html))
	expected := []string{"http://test_rss_feed_url/"}

	if len(actual) != len(expected) {
		t.Errorf("got %v\n want %v", len(actual), len(expected))
	}

	for i, _ := range expected {
		if actual[i] != expected[i] {
			t.Errorf("got %v\n want %v", actual, expected)
		}
	}
}

func TestDetectFeedWithOneAtomFeed(t *testing.T) {
	html := "<html><head><link rel=\"alternate\" type=\"application/atom+xml\" title=\"test_atom_feed_title\" href=\"http://test_atom_feed_url/\"/></head></html>"
	actual := DetectFeed(strings.NewReader(html))
	expected := []string{"http://test_atom_feed_url/"}

	if len(actual) != len(expected) {
		t.Errorf("got %v\n want %v", len(actual), len(expected))
	}

	for i, _ := range expected {
		if actual[i] != expected[i] {
			t.Errorf("got %v\n want %v", actual, expected)
		}
	}
}

func TestDetectFeedWithAtomFeed(t *testing.T) {
	html := "<html><head><link rel=\"alternate\" type=\"application/rss+xml\" title=\"test_rss_feed_title\" href=\"http://test_rss_feed_url/\"/><link rel=\"alternate\" type=\"application/atom+xml\" title=\"test_atom_feed_title\" href=\"http://test_atom_feed_url/\"/></head></html>"
	actual := DetectFeed(strings.NewReader(html))
	expected := []string{"http://test_rss_feed_url/", "http://test_atom_feed_url/"}

	if len(actual) != len(expected) {
		t.Errorf("got %v\n want %v", len(actual), len(expected))
	}

	for i, _ := range expected {
		if actual[i] != expected[i] {
			t.Errorf("got %v\n want %v", actual, expected)
		}
	}
}
