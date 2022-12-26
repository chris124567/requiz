package web

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chris124567/requiz/api"
)

type handler struct {
	template string
	client   api.Client
}

func (h handler) handleUser(writer http.ResponseWriter, user api.User) {
	sets, err := h.client.SetsByID(strconv.Itoa(user.ID))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var ourSets []api.Set
	for _, item := range sets {
		if item.CreatorID == user.ID {
			ourSets = append(ourSets, item)
		}
	}

	if err := h.writeTemplateData(writer, filepath.Join(h.template, "user.html.tmpl"), struct {
		User api.User
		Set  []api.Set
	}{user, ourSets}); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h handler) handleSet(writer http.ResponseWriter, set api.Set, mode string) {
	user, err := h.client.UserByID(strconv.Itoa(set.CreatorID))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	terms, err := h.client.Terms(strconv.Itoa(set.ID))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Set  api.Set
		Term []api.Term
		User api.User
	}{set, terms, user}
	if mode == "learn" {
		if err := h.writeTemplateData(writer, filepath.Join(h.template, "learn.html.tmpl"), data); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if mode == "write" {
		if err := h.writeTemplateData(writer, filepath.Join(h.template, "write.html.tmpl"), data); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := h.writeTemplateData(writer, filepath.Join(h.template, "set.html.tmpl"), data); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h handler) handleRoot(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		if err := h.writeTemplateData(writer, filepath.Join(h.template, "home.html.tmpl"), nil); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		var id, mode string
		if split := strings.Split(request.URL.Path, "/"); len(split) == 2 {
			id = split[1]
		} else if len(split) > 2 {
			id, mode = split[1], split[len(split)-1]
		}

		resolved, err := h.client.ResolveURL(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		} else if len(resolved.Responses) == 0 || (len(resolved.Responses[0].Models.Set) == 0 && len(resolved.Responses[0].Models.Term) == 0 && len(resolved.Responses[0].Models.User) == 0) {
			http.Error(writer, "no such page", http.StatusInternalServerError)
			return
		}

		if len(resolved.Responses[0].Models.User) > 0 {
			h.handleUser(writer, resolved.Responses[0].Models.User[0])
		} else if len(resolved.Responses[0].Models.Set) > 0 {
			h.handleSet(writer, resolved.Responses[0].Models.Set[0], mode)
		}
	}
}

func (h handler) handleSearch(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("q")
	typ := request.URL.Query().Get("type")

	var result struct {
		Type  string
		Query string

		Set  []api.Set
		User []api.User
	}
	result.Query = query
	if typ == "set" {
		results, err := h.client.SearchSets(query)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		result.Type = "set"
		result.Set = results
	} else if typ == "user" {
		results, err := h.client.SearchUsers(query)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		result.Type = "user"
		result.User = results
	}

	if err := h.writeTemplateData(writer, filepath.Join(h.template, "search.html.tmpl"), result); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h handler) handleRobotsTxt(writer http.ResponseWriter, _ *http.Request) {
	writer.Write([]byte(`User-Agent: *
Disallow: /
`))
}

func (h handler) handlePrivacyPolicy(writer http.ResponseWriter, _ *http.Request) {
	if err := h.writeTemplateData(writer, filepath.Join(h.template, "privacy-policy.html.tmpl"), nil); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h handler) handleTermsOfService(writer http.ResponseWriter, _ *http.Request) {
	if err := h.writeTemplateData(writer, filepath.Join(h.template, "terms-of-service.html.tmpl"), nil); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewHandler(static, template string, client api.Client) *http.ServeMux {
	h := handler{template, client}

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.handleRoot)
	mux.HandleFunc("/robots.txt", h.handleRobotsTxt)
	mux.HandleFunc("/search", h.handleSearch)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(static))))

	mux.HandleFunc("/privacy-policy", h.handlePrivacyPolicy)
	mux.HandleFunc("/terms-of-service", h.handleTermsOfService)

	return mux
}
