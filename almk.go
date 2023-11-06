package main

import (
	"almk/rend"
	"log"
	"os"
	"reflect"
)

type almk struct {
	name     name
	location location
	skills   skills
	work     work
	links    links
}

type name struct {
	first string
	last  string
}

type location struct {
	city  string
	state string
}

type skills struct {
	languages  []string
	tools      []string
	experience []string
}

type work struct {
	position string
	company  string
}

type links struct {
	email    string
	github   string
	linkedin string
}

func (almk *almk) Render() error {
	fields := reflect.VisibleFields(reflect.TypeOf(*almk))
	for _, field := range fields {
		err := rend.Load(field)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	almk := Self()
	err := almk.Render()
	if err != nil {
		logger := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime)
		logger.Println(err)
	}
}

func Self() *almk {
	na := name{
		first: "Abdul",
		last:  "Malek",
	}
	lo := location{
		city:  "New York",
		state: "NY",
	}
	sk := skills{
		languages: []string{
			"Python", "Go", "JavaScript", "Java", "Scala", "C++", "x86", "Lisp/Scheme", "Bash",
		},
		tools: []string{
			"Linux", "Git", "AWS", "FastAPI", "Django", "React", "CSS", "PostgreSQL", "NoSQL", "PySpark", "Airflow", "OpenGL", "LWJGL",
		},
		experience: []string{
			"distributed systems", "web", "financial systems", "security", "crypto", "operating systems", "graphics", "various math",
		},
	}
	wo := work{
		position: "Senior Software Engineer",
		company:  "Alpaca",
	}
	li := links{
		email:    "me@almk.dev",
		github:   "almk-dev",
		linkedin: "almk",
	}

	return &almk{
		name:     na,
		location: lo,
		skills:   sk,
		work:     wo,
		links:    li,
	}
}
