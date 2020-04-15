package client

import (
	"strings"
)

// Project part of repo
type Project struct {
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
}

// Repo image
type Repo struct {
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
	TagsCount int    `json:"tags_count"`
}

// Tag image
type Tag struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// Login auth cookies
func (e Ep) Login() error {
	const path = "/c/login"
	r := strings.NewReader(e.buildCredential())
	customHeader := map[string]string{
		//"Content-Type": "text/plain; charset=utf-8",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	_, err := e.reqHTTPWithCookie("POST", e.Domain+path, r, customHeader)
	if err != nil {
		return err
	}
	return nil
}

// SearchProject by keyword
func (e Ep) SearchProject(kw string) ([]Project, error) {
	path := "/api/projects?page=1&page_size=15&name=" + kw
	ps := []Project{}
	err := plainGetJSON(e.Domain+path, &ps)
	return ps, err
}

// SearchImg by keywod under a project
func (e Ep) SearchImg(pID, kw string) ([]Repo, error) {
	path := "/api/repositories?page=1&page_size=15&q=" + kw + "&project_id=" + pID
	rs := []Repo{}
	err := plainGetJSON(e.Domain+path, &rs)
	return rs, err
}

// ListImgTags list all tags
func (e Ep) ListImgTags(repo string) ([]Tag, error) {
	path := "/api/repositories/" + repo + "/tags?detail=1"
	ts := []Tag{}
	err := plainGetJSON(e.Domain+path, &ts)
	return ts, err
}
