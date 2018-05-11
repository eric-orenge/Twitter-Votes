package main

import "strings"

type Path struct {
	Path string
	ID   string
}

func NewPath(p string) *Path {
	var id string
	var pathSeparator = "/"
	p = strings.Trim(p, pathSeparator)
	s := strings.Split(p, pathSeparator)
	if len(s) > 1 {
		id = s[len(s)-1]
		p = strings.Join(s[:len(s)-1], pathSeparator)
	}

	return &Path{Path: p, ID: id}
}
func (p *Path) HasID() bool {
	return len(p.Path) > 0
}