package db

// MemDB is an in-memory database repository implementing the DB interface
// used for testing
type MemDB struct {
	installations map[int]int // accountID -> installationID
	err           error
}

// Ensure MemDB implements DB
var _ DB = (*MemDB)(nil)

// NewMemDB returns an MemDB
func NewMemDB() *MemDB {
	return &MemDB{
		installations: make(map[int]int),
	}
}

func (db *MemDB) ForceError(err error) {
	db.err = err
}

// AddGHInstallation implements DB interface
func (db *MemDB) AddGHInstallation(installationID, accountID int) error {
	db.installations[accountID] = installationID
	return db.err
}

// RemoveGHInstallation implements DB interface
func (db *MemDB) RemoveGHInstallation(accountID int) error {
	delete(db.installations, accountID)
	return db.err
}

// FindGHInstallation implements DB interface
func (db *MemDB) FindGHInstallation(accountID int) (*GHInstallation, error) {
	if installationID, ok := db.installations[accountID]; ok {
		return &GHInstallation{AccountID: accountID, InstallationID: installationID}, db.err
	}
	return nil, db.err
}