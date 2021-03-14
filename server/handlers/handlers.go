package handlers

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"html/template"
	"microblog/pkg"
	"microblog/server/service"
	"microblog/types"
	"net/http"
	"time"
)

type Handler struct {
	cnf *service.Service
}

func NewHandler(srv *service.Service) (*Handler, error) {
	return &Handler{
		cnf: srv,
	}, nil
}

func (h *Handler) Blog(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/blog.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Blog template.ParseFiles"))
		return
	}
	err = t.ExecuteTemplate(w, "blog", nil)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Blog t.ExecuteTemplate"))
		return
	}
}

func (h *Handler) Saveparamsfromblog(w http.ResponseWriter, r *http.Request) {
	var text types.Blog
	Loc, err := time.LoadLocation("Europe/Simferopol")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Saveparamsfromblog time.LoadLocation"))
		return
	}
	Now := time.Now().In(Loc).Format("2 Jan 2006 15:04:05")
	text.Text = r.FormValue("title")
	text.Anous = r.FormValue("anons")
	text.FullText = r.FormValue("full_text")
	text.Now = Now
	text.Username = r.FormValue("username")
	err = h.cnf.Saveparamsfromblog(&text)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Saveparamsfromblog h.cnf.Saveparamsfromblog"))
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Aboutus(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/aboutus.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Aboutus template.ParseFiles"))
		return
	}
	err = t.ExecuteTemplate(w, "aboutus", nil)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Aboutus t.ExecuteTemplate"))
		return
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Index template.ParseFiles"))
		return
	}
	posts := []types.Blog{}
	posts, err = h.cnf.Index(posts)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Index h.cnf.Index"))
		return
	}
	err = t.ExecuteTemplate(w, "index", posts)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Index t.ExecuteTemplate"))
		return
	}
}

func (h *Handler) Showpost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Showpost template.ParseFiles"))
		return
	}
	vars := mux.Vars(r)
	showPost := &types.Blog{}
	showPost, err = h.cnf.Showpost(showPost, vars)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Showpost h.cnf.Showpost"))
		return
	}
	err = t.ExecuteTemplate(w, "show", showPost)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with Handler in Showpost t.ExecuteTemplate"))
		return
	}
}
