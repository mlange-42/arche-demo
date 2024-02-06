package main

import (
	"log"

	"github.com/mlange-42/arche-demo/ants"
	"github.com/mlange-42/arche-demo/bees"
	"github.com/mlange-42/arche-demo/box2d"
	"github.com/mlange-42/arche-demo/evolution"
	"github.com/mlange-42/arche-demo/flocking"
	"github.com/mlange-42/arche-demo/logo"
	"github.com/mlange-42/arche-demo/matrix"
)

func run(demo string) {
	switch demo {
	case "logo":
		logo.Run()
	case "ants":
		ants.Run()
	case "bees":
		bees.Run()
	case "box2d":
		box2d.Run()
	case "evolution":
		evolution.Run()
	case "flocking":
		flocking.Run()
	case "matrix":
		matrix.Run()
	default:
		log.Fatal("Demo not found: ", demo)
	}
}
