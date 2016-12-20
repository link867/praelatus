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

func(ms *MockTeamStore) GetAll() ([]models.Team, error) {
 	return []models.Team{
		models.Team{
			Team.ID = 1
			Team.Name = "A"
			Team.Lead = models.User{
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
			Team.Members = []models.User{
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
			},			
		},
		models.Team{
			Team.ID = 1
			Team.Name = "A"
			Team.Lead = models.User{
				2,
				"foouser3",
				"foopass",
				"foo@foo3.com",
				"Foo McFooserson3",
				"",
				"",
				false,
				true,
				models.Settings{},
			}
			Team.Members = []models.User{
				models.User{
					3,
					"foouser3",
					"foopass",
					"foo@foo3.com",
					"Foo McFooserson3",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
				models.User{
					4,
					"foouser4",
					"foopass",
					"foo@foo4.com",
					"Foo McFooserson4",
					"",
					"",
					false,
					true,
					models.Settings{},
				},		
			},
		},
	},
	nil
}

func(ms *MockTeamStore) GetForUser(m *models.User) ([]models.Team, error) {
	// If I am getting the team that the user is on,
	// why am I returning a slice? 
	return []models.Team{
		models.Team{
			Team.ID = 1
			Team.Name = "A"
			Team.Lead = models.User{
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
			Team.Members = []models.User{
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
			},			
		},
		models.Team{
			Team.ID = 1
			Team.Name = "A"
			Team.Lead = models.User{
				2,
				"foouser3",
				"foopass",
				"foo@foo3.com",
				"Foo McFooserson3",
				"",
				"",
				false,
				true,
				models.Settings{},
			}
			Team.Members = []models.User{
				models.User{
					3,
					"foouser3",
					"foopass",
					"foo@foo3.com",
					"Foo McFooserson3",
					"",
					"",
					false,
					true,
					models.Settings{},
				},
				models.User{
					4,
					"foouser4",
					"foopass",
					"foo@foo4.com",
					"Foo McFooserson4",
					"",
					"",
					false,
					true,
					models.Settings{},
				},		
			},
		},
	},
	nil

}

func(ms *MockTeamStore) AddMembers(m *models.Team models.User) error {
	return nil
}

func(ms *MockTeamStore) New(*models.Team) error {
	return nil
}

func(ms *MockTeamStore) Save() error {
	return nil
}

func(ms *MockTeamStore) Remove() error {
	return nil
}

//A Mock LabelStore struct
type MockLabelStore struct {}

func (ms *MockLabelStore) Get(l *models.Label) error {
	return nil
}


//A Mock FieldStore struct
type MockFieldStore struct {}

func (MockFieldStore) Get(f *models.Field) error {
	f.ID = 1
	f.Name = "FooField"
	f.DataType = "FooDataType"
	f.Options = "" // Correct?
	return nil
}

func (MockFieldStore) GetAll() ([]models.Field, error) {
	return []models.Field{
		models.Field{
			Field.ID = 1
			Field.Name = "FooField"
			Field.DataType = "FooDataType"
			Field.Options = "" // Correct?
		},
		models.Field{
			Field.ID = 2
			Field.Name = "FooField2"
			Field.DataType = "FooDataType2"
			Field.Options = "" //Correct?
		}
	}, nil
}

func (MockFieldStore) GetByProject(p models.Project) {
	return []models.Field{
		models.Field{
			Field.ID = 1
			Field.Name = "FooField"
			Field.DataType = "FooDataType"
			Field.Options = "" // Correct?
		},
		models.Field{
			Field.ID = 2
			Field.Name = "FooField2"
			Field.DataType = "FooDataType2"
			Field.Options = "" //Correct?
		}
	}, nil
}

func (MockFieldStore) AddToProject(p models.Project, f *models.Field, t ...models.TicketType) error {
	return nil
}

func (MockFieldStore) New(f *models.Field) error {
	return nil
}

func (MockFieldStore) Save(f *models.Field) error {
	return nil
}

func (MockFieldStore) Remove(f *models.Field) error {
	return nil
}

// //A Mock TicketStore struct
type MockTicketStore struct {}

func (MockTicketStore) Get(t *models.Ticket) error {
	t.ID = 1
	t.CreatedDate = "000000"
	t.UpdatedDate = "000000"
	t.Key = "fooKey"
	t.Summary = "fooSummary"
	t.Description = "fooDescription"
	t.Fields = []models.FieldValue{
		FieldValue{

		},
		FieldValue{

		},
	}
	t.Labels = []models.Label{
		models.Label{

		},
		models.Label{

		}
	}
	t.Type = ""
	t.Reporter = models.User{
			4,
			"foouser4",
			"foopass",
			"foo@foo4.com",
			"Foo McFooserson4",
			"",
			"",
			false,
			true,
			models.Settings{},
	}
	t.Reporter = models.User{
			4,
			"foouser4",
			"foopass",
			"foo@foo4.com",
			"Foo McFooserson4",
			"",
			"",
			false,
			true,
			models.Settings{},
	}
	t.Status = ""
	return nil
}

// //A Mock TypeStore struct
// type MockTypeStore struct {}

// //A Mock ProjectStore struct
// type MockProjectStore struct {}

// //A Mock StatusStore struct
// type MockStatusStore struct {}

// //A Mock Workflow Store
// type MoskWorkflowStore struct {}
