package brokenspider

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestAbsolutize(t *testing.T) {
	testCases := []struct {
		baseURL  string
		url      string
		expected string
	}{
		{"http://example.com", "./path/to/file", "http://example.com/path/to/file"},
		{"", "/path/to/file", "/path/to/file"},
		{"http://example.com", "http://example.com/file.html", "http://example.com/file.html"},
	}

	for _, tc := range testCases {
		actual := absolutize(tc.baseURL, tc.url)
		if actual != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, actual)
		}
	}
}

func TestIsAbsolute(t *testing.T) {
	testCases := []struct {
		url string
		exp bool
	}{
		{"http://example.com", true},
		{"/path/to/file.html", false},
		{"../path/to/file.html", false},
		{"https://thealamu.tech", true},
		{".", false},
		{"./", false},
		{"/", false},
	}
	for _, tc := range testCases {
		if got := isAbsolute(tc.url); got != tc.exp {
			t.Errorf("isAbsolute(%s) = %v, want %v", tc.url, got, tc.exp)
		}
	}
}

func TestIsBroken(t *testing.T) {
	testCases := []struct {
		url      string
		isBroken bool
	}{
		{"https://www.google.com", false},
		{"https://www.google.com/", false},
		{"https://thealamu.com", true},
		{"https://broken.go", true},
		{"https://", true},
		{"h", true},
		{"https://thealamu.tech", false},
	}

	for _, testCase := range testCases {
		isBroken := isBroken(testCase.url)
		if isBroken != testCase.isBroken {
			t.Errorf("isBroken(%s) = %t, want %t", testCase.url, isBroken, testCase.isBroken)
		}
	}
}

func TestWalk(t *testing.T) {
	// serve the test fixtures folder
	testErrChan := make(chan error)
	serveTestFixtures(testErrChan)

	// give test server 4 seconds to start unless it errors
outer:
	for {
		select {
		case <-testErrChan:
			t.Error("test fixtures server failed to start")
			t.FailNow()
		case <-time.After(4 * time.Second):
			break outer
		}
	}

	resultChan := make(chan LinkStatus)
	spider := New()
	go func() {
		spider.Walk("http://localhost:6969", resultChan)
	}()

	linkStatuses := []LinkStatus{}
	broken := 0
	for status := range resultChan {
		fmt.Println(status)
		linkStatuses = append(linkStatuses, status)
		// calculate broken links ahead of time
		if status.Broken {
			broken++
		}
	}

	if len(linkStatuses) != 10 {
		t.Errorf("Expected 10 link statuses, got %d", len(linkStatuses))
	}

	if broken != 4 {
		t.Errorf("Expected 4 broken links, got %d", broken)
	}
}

func serveTestFixtures(errChan chan error) {
	testMux := http.NewServeMux()
	testMux.Handle("/", http.FileServer(http.Dir("testfixtures")))

	go func() {
		fmt.Println("Starting testfixtures server")
		err := http.ListenAndServe(":6969", testMux)
		if err != nil {
			errChan <- err
		}
	}()
}
