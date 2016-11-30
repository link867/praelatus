package store

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/praelatus/backend/models"
)

var seedFuncs = []func(s Store) error{
	SeedUsers,
	SeedTeams,
	SeedProjects,
	SeedTicketTypes,
	SeedFields,
	SeedStatuses,
	SeedLabels,
	SeedTickets,
	SeedComments,
}

// SeedAll will run all of the seed functions
func SeedAll(s Store) error {
	fmt.Println("Seeding All")
	for _, f := range seedFuncs {
		e := f(s)
		if e != nil {
			return e
		}
	}

	return nil
}

// SeedLabels will add some test labesl to the database
func SeedLabels(s Store) error {
	labels := []models.Label{
		models.Label{
			Name: "test",
		},
		models.Label{
			Name: "duplicate",
		},
		models.Label{
			Name: "wontfix",
		},
	}

	for _, l := range labels {
		e := s.Labels().New(&l)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedTickets will add some test tickets to the database
func SeedTickets(s Store) error {
	fmt.Println("Seeding tickets")
	for i := 0; i < 50; i++ {
		t := &models.Ticket{
			Summary:     "This is a test ticket. #" + strconv.Itoa(i),
			Description: "No really, this is just a test",
			Reporter:    models.User{ID: 1},
			Assignee:    models.User{ID: 1},
			Labels: []models.Label{
				models.Label{1, "test"},
			},
			Fields: []models.FieldValue{
				models.FieldValue{
					Name:  "Story Points",
					Value: rand.Int(),
				},
				models.FieldValue{
					Name: "Priority",
					Value: models.FieldOption{
						Default:  "LOW",
						Selected: []string{"HIGH", "MEDIUM", "LOW"}[rand.Intn(3)],
						Options:  []string{"HIGH", "MEDIUM", "LOW"},
					},
				},
			},
			Type: models.TicketType{ID: 1},
		}

		e := s.Tickets().New(models.Project{ID: 1}, t)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedStatuses will add some ticket statuses to the database
func SeedStatuses(s Store) error {
	statuses := []models.Status{
		models.Status{
			Name: "Open",
		},
		models.Status{
			Name: "In Progress",
		},
		models.Status{
			Name: "Done",
		},
	}

	fmt.Println("Seeding statuses")
	for _, st := range statuses {
		e := s.Statuses().New(&st)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedComments will add some comments to all tickets
func SeedComments(s Store) error {
	fmt.Println("Seeding comments")
	t, se := s.Tickets().GetAll()
	if se != nil {
		return se
	}

	for _, tk := range t {
		for x := 0; x < 25; x++ {
			c := &models.Comment{
				Body: fmt.Sprintf(`This is the %d th comment
				# Yo Dawg
				**I** *heard* you
				> like markdown
				so I put markdown in your comments`, x),
				Author: models.User{ID: 1},
			}

			e := s.Tickets().NewComment(tk, c)
			if e != nil && e != ErrDuplicateEntry {
				return e
			}

			if e == ErrDuplicateEntry {
				return nil
			}
		}

	}

	return nil
}

// SeedFields will seed the given store with some test Fields.
func SeedFields(s Store) error {
	fields := []models.Field{
		models.Field{
			Name:     "Story Points",
			DataType: "INT",
		},
		models.Field{
			Name:     "TestField2",
			DataType: "FLOAT",
		},
		models.Field{
			Name:     "TestField3",
			DataType: "INT",
		},
		models.Field{
			Name:     "TestField4",
			DataType: "DATE",
		},
		models.Field{
			Name:     "Priority",
			DataType: "OPT",
			Options: models.FieldOption{
				Default:  "LOW",
				Selected: "LOW",
				Options:  []string{"HIGH", "MEDIUM", "LOW"},
			},
		},
	}

	fmt.Println("Seeding fields")
	for _, f := range fields {
		e := s.Fields().New(&f)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			continue
		}

		e = s.Fields().AddToProject(models.Project{ID: 1}, &f)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
		}
	}

	return nil
}

// SeedProjects will seed the given store with some test projects.
func SeedProjects(s Store) error {
	projects := []models.Project{
		models.Project{
			Name: "TEST Project",
			Key:  "TEST",
			Lead: models.User{ID: 1},
		},
		models.Project{
			Name: "TEST Project 2",
			Key:  "TEST2",
			Lead: models.User{ID: 2},
		},
	}

	fmt.Println("Seeding projects")
	for _, p := range projects {
		e := s.Projects().New(&p)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
		}
	}

	return nil
}

// SeedTeams will seed the database with some test Teams.
func SeedTeams(s Store) error {
	teams := []models.Team{
		models.Team{
			Name: "The A Team",
			Lead: models.User{
				ID: 1,
			},
			Members: []models.User{
				models.User{ID: 1},
				models.User{ID: 2},
			},
		},
		models.Team{
			Name: "The B Team",
			Lead: models.User{
				ID: 2,
			},
			Members: []models.User{
				models.User{ID: 1},
				models.User{ID: 2},
			},
		},
	}

	fmt.Println("Seeding teams")
	for _, team := range teams {
		team.Lead = models.User{ID: 1}

		e := s.Teams().New(&team)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
		}
	}

	return nil
}

// SeedTicketTypes will seed the database with some test TicketTypes.
func SeedTicketTypes(s Store) error {
	types := []models.TicketType{
		models.TicketType{
			Name: "Bug",
		},
		models.TicketType{
			Name: "Epic",
		},
		models.TicketType{
			Name: "Story",
		},
		models.TicketType{
			Name: "Feature",
		},
		models.TicketType{
			Name: "Question",
		},
	}

	fmt.Println("Seeding ticket types")
	for _, t := range types {
		e := s.Types().New(&t)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
		}
	}

	return nil
}

// SeedUsers will seed the database with some test users.
func SeedUsers(s Store) error {
	t1, be := models.NewUser("testuser", "test", "Test Testerson",
		"test@example.com", false)
	if be != nil {
		return be
	}

	t2, be := models.NewUser("testadmin", "test", "Test Testerson II",
		"test1@example.com", false)
	if be != nil {
		return be
	}

	users := []models.User{
		*t1,
		*t2,
	}

	fmt.Println("Seeding users")
	for _, u := range users {
		e := s.Users().New(&u)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
		}
	}

	return nil
}
