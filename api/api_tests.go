package api

import (
	"testing"

	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

func TestMain(m *testing.M) {

}

type MockStore struct{}

func (ms *MockStore) Users() store.UserStore {
	return MockUsersStore{}
}

func (ms *MockStore) Teams() store.TeamStore {
	return MockTeamStore{}
}

func (ms *MockStore) Labels() store.LabelStore {
	return MockLabelStore{}
}

func (ms *MockStore) Fields() store.FieldStore {
	return MockFieldStore{}
}

func (ms *MockStore) Tickets() store.TicketStore {
	return MockTicketStore{}
}

func (ms *MockStore) Projects() store.ProjectStore {
	return MockProjectStore{}
}

func (ms *MockStore) Statuses() store.StatusStore {
	return MockStatusStore{}
}

func (ms *MockStore) Workflows() store.WorkflowStore {
	return MockWorkflowStore{}
}

//A Mock UsersStore struct
type MockUsersStore struct{}

func (ms *MockUsersStore) Get(u *models.User) error {
	u.ID = 1
	u.Username = "foouser"
	u.Password = "foopass"
	u.Email = "foo@foo.com"
	u.FullName = "Foo McFooserson"
	u.IsActive = true
	return nil
}

func (ms *MockUsersStore) GetAll() ([]models.User, error) {
	return []models.User{
		models.User{
			1,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
		models.User{
			2,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
	}, nil
}

func (ms *MockUsersStore) New() error {
	return nil
}

func (ms *MockUsersStore) Save() error {
	return nil
}

func (ms *MockUsersStore) Remove() error {
	return nil
}

//A Mock TeamStore struct
type MockTeamStore struct{}

func (ms *MockTeamStore) Get(t *models.Team) error {
	t.ID = 1
	t.Name = "A"
	t.Lead = models.User{
		1,
		"foouser",
		"foopass",
		"foo@foo.com",
		"Foo McFooserson",
		"",
		"",
		false,
		true,
		models.Settings{},
	}
	t.Members = []models.User{
		models.User{
			1,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
		models.User{
			2,
			"foouser",
			"foopass",
			"foo@foo.com",
			"Foo McFooserson",
			"",
			"",
			false,
			true,
			models.Settings{},
		},
	}
	return nil
}

// func(ms *MockTeamStore) GetAll() ([]models.Team, error) {
// 	return []models.Team, error
// }

// func(ms *MockTeamStore) GetForUser() ([]models.Team, error) {
// 	return []models.Team, error
// }

// func(ms *MockTeamStore) AddMembers() error {
// 	return error
// }

// func(ms *MockTeamStore) New() error {
// 	return error
// }

// func(ms *MockTeamStore) Save() error {
// 	return error
// }

// func(ms *MockTeamStore) Remove() error {
// 	return error
// }

// //A Mock LabelStore struct
// type MockLabelStore struct {}

// func (ms *MockLabelStore) Get() error {
// 	return nil
// }

// func (ms)

// //A Mock FieldStore struct
// type MockFieldStore struct {}

// //A Mock TicketStore struct
// type MockTicketSotre struct {}

// //A Mock TypeStore struct
// type MockTypeStore struct {}

// //A Mock ProjectStore struct
// type MockProjectStore struct {}

// //A Mock StatusStore struct
// type MockStatusStore struct {}

// //A Mock Workflow Store
// type MoskWorkflowStore struct {}
