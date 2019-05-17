package admin

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Tairos/strame/admin/command"
	"github.com/gorilla/mux"
)

type Panel struct {
	resourcesPath string
	commands      []*command.Command
}

func NewPanel(resourcesPath string) *Panel {
	return &Panel{
		resourcesPath: resourcesPath,
	}
}

func (p *Panel) Run(ctx context.Context) {
	p.runWeb(ctx)
	p.runCommunication(ctx)
}

func (p *Panel) runWeb(ctx context.Context) {
	go func() {
		r := mux.NewRouter()
		tmpl, err := template.ParseFiles(p.resourcesPath + "/resources/html/admin_panel.html")
		if err != nil {
			panic(err)
		}

		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Cookies())
			tmpl.Execute(w, nil)
		}).Methods("GET")

		if err := http.ListenAndServeTLS(":443", p.resourcesPath+"/resources/server.crt", p.resourcesPath+"/resources/server.key", r); err != nil {
			panic(err)
		}
	}()
}

func (p *Panel) runCommunication(ctx context.Context) {
			wsHandler := func handler(w http.ResponseWriter, r *http.Request) {
		    conn, err := upgrader.Upgrade(w, r, nil)
		    if err != nil {
		        log.Println(err)
		        return
		    }

			if err := http.ListenAndServeTLS(":44300", p.resourcesPath+"/resources/server.crt", p.resourcesPath+"/resources/server.key", wsHandler); err != nil {
				panic(err)
			}
		}
}
