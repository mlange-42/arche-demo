---
title: Ant Colony
wasm: ants
description: |
    A stylized model of ants foraging and scouting,
    using trail pheromones.

    Ants lay trails of two types of pheromones:
    one when coming from the nest,
    and another one when coming from a resource.

    Using these trails, workers find their way to resources (white) and back to the nest ({{< col cyan >}}cyan{{< /col >}}).
    To find new resources, scouts ({{< col yellow >}}yellow{{< /col >}}) swarm out and do a random walk until they find something.

    In this example, the landscape is a network with nodes and edges represented by ECS entities.
---

{{< rawhtml >}}
{{< canvas 880 480 >}}

<p id="instructions">Click to pause and resume.</p>
<p class="tt">go get <a href="https://github.com/mlange-42/arche">github.com/mlange-42/arche</a>
</p>
{{< /rawhtml >}}
