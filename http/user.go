package http

import (
	"log"
	"net/http"
)

func (h *UserHandlers) sessionCreate(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "sessionCreate",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) sessionDelete(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "sessionDelete",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) userCreate(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "userCreate",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) userForgotPassword(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "userForgotPassword",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminMembersIndex(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminMembersIndex",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminMembersCreate(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminMembersCreate",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminMembersRead(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminMembersRead",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminMembersUpdate(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminMembersUpdate",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminMembersDelete(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminMembersDelete",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminAdminsIndex(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminAdminsIndex",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminAdminsCreate(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminAdminsCreate",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminAdminsRead(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminAdminsRead",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminAdminsUpdate(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminAdminsUpdate",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandlers) adminAdminsDelete(w http.ResponseWriter, r *http.Request) {
	details := &funcDetails{
		Name:        "adminAdminsDelete",
		Description: "Not implemented",
	}
	log.Println("[HANDLER]", details.Name)
	handleErrorAndRespond(details, w, "Not implemented", http.StatusNotImplemented)
}
