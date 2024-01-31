---
title: Evolution
wasm: evolution
description: |
    An evolutionary model of grazer behavior.

    Grazers have genes for 4 traits:  
    {{< tt >}}MaxAngle{{< /tt >}} is the maximum turning angle,
    {{< tt >}}MinGrass{{< /tt >}} is the grass threshold to stop grazing and start searching,
    {{< tt >}}Investment{{< /tt >}} is the energy invested in offspring,
    {{< tt >}}Offspring{{< /tt >}} is the number of offspring.

    Colors of grazers are also inherited and mutated, but have no influence on fitness.
    Grazers can only mate with other grazers with a similar color.
---

{{< rawhtml >}}
{{< canvas 880 480 >}}

<p id="instructions">Use <span class="tt">PageUp</span> and <span class="tt">PageDown</span> to adjust simulation speed. Click to pause and resume.</p>
<p class="tt">go get <a href="https://github.com/mlange-42/arche">github.com/mlange-42/arche</a>
</p>
{{< /rawhtml >}}
