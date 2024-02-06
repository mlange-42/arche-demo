# Arche Demo &ndash; Flocking

Classical boids model,
mimicking the behaviour of bird flocks or fish schools.

Each boid reacts with simple rules to it's 8 nearest neighbors.
Gonum's [`kdtree`](https://pkg.go.dev/gonum.org/v1/gonum/spatial/kdtree)
is used to speed up the lookup of nearby individuals.

See the [live demo](https://mlange-42.github.io/arche-demo/flocking/) for this example.
