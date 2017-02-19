package http

import (
	"net/http"
	"strings"

	jjwebdev "github.com/jjwebdev/go-template/models"

	"github.com/blockninja/ninjarouter"
)

type muxMux struct {
	files *http.ServeMux
	users *UserHandlers
}

func (mm *muxMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API") != "" {
		if strings.HasPrefix(r.URL.String(), "/users") {
			mm.users.ServeHTTP(w, r)
		}
	} else {
		mm.files.ServeHTTP(w, r)
	}
}

// UserHandlers contain all the handlers for users
type UserHandlers struct {
	*ninjarouter.Mux
	UserService jjwebdev.UserService
}

func newUserHandlers(us jjwebdev.UserService) *UserHandlers {
	h := &UserHandlers{
		Mux:         ninjarouter.New(),
		UserService: us,
	}

	// Member paths
	h.Mux.DELETE("/users/sign_out", userChain(h.sessionDelete))

	// Public Paths
	h.Mux.POST("/users/sign_in", publicChain(h.sessionCreate))
	h.Mux.POST("/users/create", publicChain(h.userCreate))
	h.Mux.PATCH("/users/forgot", publicChain(h.userForgotPassword))

	// Admin Paths
	h.Mux.GET("/users/members/all", adminChain(h.adminMembersIndex))
	h.Mux.POST("/users/members/create", adminChain(h.adminMembersCreate))
	h.Mux.GET("/users/members/read/:id", adminChain(h.adminMembersRead))
	h.Mux.PATCH("/users/members/update/:id", adminChain(h.adminMembersUpdate))
	h.Mux.DELETE("/users/members/delete/:id", adminChain(h.adminMembersDelete))

	h.Mux.GET("/users/admins/all", adminChain(h.adminAdminsIndex))
	h.Mux.POST("/users/admins/create/:id", adminChain(h.adminAdminsCreate))
	h.Mux.GET("/users/admins/read/:id", adminChain(h.adminAdminsRead))
	h.Mux.PATCH("/users/admins/update/:id", adminChain(h.adminAdminsUpdate))
	h.Mux.DELETE("/users/admins/delete/:id", adminChain(h.adminAdminsDelete))

	h.Mux.NotFound = http.HandlerFunc(notFound)

	return h
}

func fileRoutes() *http.ServeMux {
	filesMux := http.NewServeMux()
	filesMux.Handle("/", http.FileServer(http.Dir("./web/dist/")))
	return filesMux
}

func routes(us jjwebdev.UserService) *muxMux {
	mm := &muxMux{
		users: newUserHandlers(us),
		files: fileRoutes(),
	}

	return mm
}
