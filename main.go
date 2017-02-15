package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type muxMux struct {
	files *http.ServeMux
	api   *httprouter.Router
}

func (mm *muxMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API") != "" {
		mm.api.ServeHTTP(w, r)
	} else {
		mm.files.ServeHTTP(w, r)
	}
}

func secureChain(next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return withLogging(http.HandlerFunc(next)).ServeHTTP
}

func insecureChain(next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return withLogging(http.HandlerFunc(next)).ServeHTTP
}

func main() {
	apiMux := httprouter.New()

	apiMux.POST("/users/sign_in", sessionCreate)
	apiMux.DELETE("/users/sign_out", sessionDelete)

	apiMux.POST("/users", userCreate)
	apiMux.PATCH("/users", userForgotPassword)

	apiMux.GET("/admin/members/index/", adminMembersIndex)
	apiMux.POST("/admin/members/create/:id", adminMembersCreate)
	apiMux.GET("/admin/members/read/:id", adminMembersRead)
	apiMux.PATCH("/admin/members/update/:id", adminMembersUpdate)
	apiMux.DELETE("/admin/members/delete/:id", adminMembersDelete)

	apiMux.GET("/admin/admins/index/", adminAdminsIndex)
	apiMux.POST("/admin/admins/create/:id", adminAdminsCreate)
	apiMux.GET("/admin/admins/read/:id", adminAdminsRead)
	apiMux.PATCH("/admin/admins/update/:id", adminAdminsUpdate)
	apiMux.DELETE("/admin/admins/delete/:id", adminAdminsDelete)

	filesMux := http.NewServeMux()
	filesMux.Handle("/", http.FileServer(http.Dir("./web/dist/")))

	mm := &muxMux{
		api:   apiMux,
		files: filesMux,
	}

	log.Println("Your server is listening on port: 8080")
	log.Fatalln(http.ListenAndServe(":8080", mm))
}

func sessionCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func sessionDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func userCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func userForgotPassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func entryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminMembersIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminMembersCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminMembersRead(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminMembersUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminMembersDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminAdminsIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminAdminsCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminAdminsRead(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminAdminsUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}

func adminAdminsDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Not implemented")
}
