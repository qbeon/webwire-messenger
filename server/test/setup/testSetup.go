package setup

import (
	"testing"
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv"
	"github.com/qbeon/webwire-example-postboard/server/client"
)

// TestSetup represents the prepared setup of an individual test
type TestSetup struct {
	t       *testing.T
	conf    *Config
	clients []client.ApiClient

	ApiServer apisrv.ApiServer
	Helper    *Helper
}

// newTestSetup creates a new test setup
func newTestSetup(t *testing.T, conf *Config) *TestSetup {
	// Start recording test setup time
	start := time.Now()

	// Setup test database and connections to it here

	// Launch API server
	apiServer := launchApiServer(t, conf.serverConfig)

	// Create a new test setup instance
	testSetup := &TestSetup{
		t:         t,
		conf:      conf,
		clients:   make([]client.ApiClient, 0, 3),
		ApiServer: apiServer,
	}

	// Initialize test helper
	testSetup.Helper = &Helper{
		t:         t,
		ts:        testSetup,
		apiServer: apiServer,
	}

	// Setup test database state here

	// Record test setup time
	conf.statisticsRecorder.Set(t, func(stat *TestStatistics) {
		stat.SetupTime = time.Since(start)
	})

	return testSetup
}

// MaxCreationTimeDeviation returns the configured maximum accepted
// entity creation time deviation duration
func (ts *TestSetup) MaxCreationTimeDeviation() time.Duration {
	return ts.conf.defaultMaxCreationTimeDeviation
}

// Teardown gracefully terminates the test,
// this method MUST BE DEFERRED until the end of the test!
func (ts *TestSetup) Teardown() {
	// Start recording test teardown time
	start := time.Now()

	// Disconnect API clients
	for _, clt := range ts.clients {
		clt.Close()
	}

	// Stop the API server instance
	if err := ts.ApiServer.Shutdown(); err != nil {
		// Don't break on shutdown failure, remove database before quitting!
		ts.t.Errorf("API server shutdown failed: %s", err)
	}

	// Delete test database here

	// Record test teardown time
	ts.conf.statisticsRecorder.Set(ts.t, func(stat *TestStatistics) {
		stat.TeardownTime = time.Since(start)
	})
}
