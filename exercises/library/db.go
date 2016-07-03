package library

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/hashicorp/go-memdb"
)

type (
	// MemDB is an in-memory database.
	MemDB struct {
		*memdb.MemDB
	}

	// CourseModel is the database model used to persist courses.
	CourseModel struct {
		// ID is the course identifier.
		ID string
		// Name is the unique course name.
		Name string
		// Description is the course description.
		Description string
		// StatTime is the data and time the course starts.
		StartTime time.Time
		// EndTime is the date and time the course ends.
		EndTime time.Time
		// Location is the place the course is being held at.
		Location string
	}

	// RegistrationModel is the database model used to persist registrations.
	RegistrationModel struct {
		// ID is the registration identifier.
		ID string
		// CourseID is the course identifier.
		CourseID string
		// FirstName is the attendee first name.
		FirstName string
		// LastName is the attendee last name.
		LastName string
		// Address is the attendee street address.
		Address *Address
	}

	// Address is a street address.
	Address struct {
		// Number is the street number.
		Number int
		// Street is the street name.
		Street string
		// City is the city name.
		City string
		// State is the US state 2 letter code.
		State string
		// Zip is the US Zip code.
		Zip int
	}

	// UserModel is the database model used to persist users.
	UserModel struct {
		// Email is the user email.
		Email string
		// HashedPassword is the user password hashed with bcrypt.
		HashedPassword string
	}
)

// ErrNotFound is the error returned when a model cannot be found.
var ErrNotFound = errors.New("model not found")

// NewDB returns a new in-memory database pre-initialized with seed data.
func NewDB() *MemDB {
	db, err := memdb.NewMemDB(schema())
	if err != nil {
		panic(err.Error())
	}
	txn := db.Txn(true)
	defer txn.Commit()
	txn.Insert("courses", buildingAPIs)
	txn.Insert("courses", goa101)
	txn.Insert("registrations", baRegistration1)
	txn.Insert("registrations", baRegistration2)
	txn.Insert("registrations", baRegistration3)
	txn.Insert("users", brian)
	txn.Insert("users", raphael)
	txn.Insert("users", simple)

	return &MemDB{db}
}

// Insert inserts the data in the given table.
func (db *MemDB) Insert(table string, data interface{}) error {
	txn := db.Txn(true)
	defer txn.Commit()
	return txn.Insert(table, data)
}

// Get retrieves the entry from table where the field fieldName has the value fieldVal.
func (db *MemDB) Get(table string, fieldName string, fieldVal interface{}) (interface{}, error) {
	txn := db.Txn(false)
	defer txn.Abort()
	m, err := txn.First(table, fieldName, fieldVal)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, ErrNotFound
	}
	return m, nil
}

// List retrieves all entries from the table where the field fieldName has the value fieldVal.
// If fieldVal is nil then List returns all the entries.
func (db *MemDB) List(table string, fieldName string, fieldVal interface{}) ([]interface{}, error) {
	txn := db.Txn(false)
	defer txn.Abort()
	var it memdb.ResultIterator
	var err error
	if fieldVal == nil {
		it, err = txn.Get(table, fieldName)
	} else {
		it, err = txn.Get(table, fieldName, fieldVal)
	}
	if err != nil {
		return nil, err
	}
	var entries []interface{}
	e := it.Next()
	for e != nil {
		entries = append(entries, e)
		e = it.Next()
	}
	return entries, nil
}

// Delete deletes the entry from table where the field fieldName has the value fieldVal.
func (db *MemDB) Delete(table string, fieldName string, fieldVal interface{}) error {
	txn := db.Txn(true)
	defer txn.Commit()
	n, err := txn.DeleteAll(table, fieldName, fieldVal)
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

// NewEmptyDB returns an empty database, mainly intended for tests.
func NewEmptyDB() *MemDB {
	mdb, err := memdb.NewMemDB(schema())
	if err != nil {
		panic(err.Error())
	}
	return &MemDB{mdb}
}

// Wipe deletes the entire DB content, mainly intended for tests.
func (db *MemDB) Wipe() {
	mdb, err := memdb.NewMemDB(schema())
	if err != nil {
		panic(err.Error())
	}
	db.MemDB = mdb
}

// parseTime is a helper function that calls time.Parse and panics if it returns an error.
func parseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

// hashed is a helper function that hashes the given password using bcrypt.
func hashed(pass string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

var (
	buildingAPIs = &CourseModel{
		ID:          "1",
		Name:        "Building APIs in Go",
		Description: "Learn how to build HTTP APIs using the Go language",
		StartTime:   parseTime(time.RFC822, "10 Jul 16 09:00 MST"),
		EndTime:     parseTime(time.RFC822, "10 Jul 16 17:00 MST"),
		Location:    "Denver",
	}
	goa101 = &CourseModel{
		ID:          "2",
		Name:        "Mastering goa",
		Description: "Learn how to build microservices using goa",
		StartTime:   parseTime(time.RFC822, "15 Aug 16 09:00 MST"),
		EndTime:     parseTime(time.RFC822, "15 Aug 16 17:00 MST"),
		Location:    "Santa Barbara",
	}

	baRegistration1 = &RegistrationModel{
		ID:        "1",
		CourseID:  "1",
		FirstName: "Gopher",
		LastName:  "Extraordinay",
		Address: &Address{
			Number: 42,
			Street: "Main Street",
			City:   "Denver",
			State:  "CO",
			Zip:    80205,
		},
	}
	baRegistration2 = &RegistrationModel{
		ID:        "2",
		CourseID:  "1",
		FirstName: "Another",
		LastName:  "Gopher",
		Address: &Address{
			Number: 43,
			Street: "Main Street",
			City:   "Denver",
			State:  "CO",
			Zip:    80205,
		},
	}
	baRegistration3 = &RegistrationModel{
		ID:        "3",
		CourseID:  "1",
		FirstName: "Last",
		LastName:  "Gopher",
		Address: &Address{
			Number: 44,
			Street: "Main Street",
			City:   "Denver",
			State:  "CO",
			Zip:    80205,
		},
	}

	brian = &UserModel{
		Email:          "brian@gopheracademy.com",
		HashedPassword: hashed("s33kret"),
	}
	raphael = &UserModel{
		Email:          "raphael@goa.design",
		HashedPassword: hashed("s33kret2"),
	}
	simple = &UserModel{
		Email:          "email",
		HashedPassword: hashed("password"),
	}
)

var idMx = &sync.Mutex{}
var counter = 9

// NewID creates model IDs, in a real app this would be done by the DB or would use some form of GUID.
func NewID() string {
	idMx.Lock()
	defer idMx.Unlock()
	counter++
	return strconv.Itoa(counter)
}

// In-memory database schema.
func schema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"courses": &memdb.TableSchema{
				Name: "courses",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"idxName": &memdb.IndexSchema{
						Name:    "idxName",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
			"registrations": &memdb.TableSchema{
				Name: "registrations",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"courseIDIdx": &memdb.IndexSchema{
						Name:    "courseIDIdx",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "CourseID"},
					},
				},
			},
			"users": &memdb.TableSchema{
				Name: "users",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
				},
			},
		},
	}
}
