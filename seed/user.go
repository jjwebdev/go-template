package seed

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"sync"

	"github.com/betacraft/yaag/yaag"
	"github.com/jjwebdev/go-template/http"
	"github.com/jjwebdev/go-template/models"
	"github.com/syscrusher/fake"
)

func (ss *SeedService) SeedRoles() {
	log.Println("Seeding Roles...")
	handler := http.NewServer(":8080", ss.UserService)
	roles := []*models.Role{
		{Name: "admin"},
		{Name: "member"},
	}

	for _, role := range roles {
		b, err := json.Marshal(role)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", "http://localhost:8080/users/roles/create", bytes.NewReader(b))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API", "v1")
		rr := httptest.NewRecorder()

		handler.Handler.ServeHTTP(rr, req)
	}
}

// SeedUsers will seed the users
func (ss *SeedService) SeedUsers(num int, isAdmin bool) {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Go Template", DocPath: "apidoc.html"})
	handler := http.NewServer(":8080", ss.UserService)
	roleID := 2
	if isAdmin {
		log.Println("Seeding Admins...")
		roleID = 1
	} else {
		log.Println("Seeding Members...")
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			u := &models.User{
				FirstName:    fake.FirstName(),
				LastName:     fake.LastName(),
				Email:        fake.EmailAddress(),
				Password:     fake.SimplePassword(),
				SessionToken: fake.CharactersN(10),
				RoleID:       roleID,
			}
			b, err := json.Marshal(u)
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest("POST", "http://localhost:8080/users/create", bytes.NewReader(b))
			if err != nil {
				panic(err)
			}
			req.Header.Set("X-API", "v1")
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.Handler.ServeHTTP(rr, req)
			wg.Done()
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
