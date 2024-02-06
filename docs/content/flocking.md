---
title: Flocking
wasm: flocking
description: |
    Classical boids model,
    mimicking the behaviour of bird flocks or fish schools.

    Each boid reacts with simple rules to it's 8 nearest neighbors.
    Gonum's [{{< tt >}}kdtree{{< /tt >}}](https://pkg.go.dev/gonum.org/v1/gonum/spatial/kdtree)
    is used to speed up the lookup of nearby individuals.
---

{{< rawhtml >}}
{{< canvas 880 480 >}}

<p id="instructions">Move the mouse over the canvas! Use <span class="tt">PageUp</span> and <span class="tt">PageDown</span> to adjust simulation speed. Click to pause and resume.</p>
<p class="tt">go get <a href="https://github.com/mlange-42/arche">github.com/mlange-42/arche</a>
</p>
{{< /rawhtml >}}
