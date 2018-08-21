package setup

import (
	"testing"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv"
	"github.com/qbeon/webwire-messenger/server/apisrv/config"
	"github.com/qbeon/webwire-messenger/server/client"
)

// TestSetup represents the prepared setup of an individual test
type TestSetup struct {
	t       *testing.T
	stats   *StatisticsRecorder
	clients []client.ApiClient

	ApiServer apisrv.ApiServer
	Helper    *Helper
}

// newTestSetup creates a new test setup
func newTestSetup(
	t *testing.T,
	statsRec *StatisticsRecorder,
	config config.Config,
) *TestSetup {
	// Start recording test setup time
	start := time.Now()

	// Setup test database and connections to it here

	// Launch API server
	apiServer := launchApiServer(t, config)

	// Create a new test setup instance
	testSetup := &TestSetup{
		t:         t,
		stats:     statsRec,
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
	statsRec.Set(t, func(stat *TestStatistics) {
		stat.SetupTime = time.Since(start)
	})

	return testSetup
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
	ts.stats.Set(ts.t, func(stat *TestStatistics) {
		stat.TeardownTime = time.Since(start)
	})
}
