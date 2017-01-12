package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.POST("/users/sign_in", sessionCreate)
	r.DELETE("/users/sign_out", sessionDelete)

	r.POST("/users", userCreate)
	r.PATCH("/users", userForgotPassword)

	r.GET("/admin/members/index/", adminMembersIndex)
	r.POST("/admin/members/create/:id", adminMembersCreate)
	r.GET("/admin/members/read/:id", adminMembersRead)
	r.PATCH("/admin/members/update/:id", adminMembersUpdate)
	r.DELETE("/admin/members/delete/:id", adminMembersDelete)

	r.GET("/admin/admins/index/", adminAdminsIndex)
	r.POST("/admin/admins/create/:id", adminAdminsCreate)
	r.GET("/admin/admins/read/:id", adminAdminsRead)
	r.PATCH("/admin/admins/update/:id", adminAdminsUpdate)
	r.DELETE("/admin/admins/delete/:id", adminAdminsDelete)

	r.GET("/", entryHandler)

	http.ListenAndServe(":8080", r)
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

func adminMembersCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Paramzs) {
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
