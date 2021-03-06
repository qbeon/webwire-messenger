package memeng

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// GetUser implements the Engine interface
func (eng *engine) GetUser(ident api.Identifier) (*api.User, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Search by user.identifier index
	if account, exists := eng.users[ident]; exists {
		return &account.Profile, nil
	}

	// Not found
	return nil, nil
}
