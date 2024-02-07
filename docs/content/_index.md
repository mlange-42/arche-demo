---
title: Home
wasm: logo
description: |
    A simple simulation with 30.000 entities.
    
    Each entity has a target pixel in the logo.
    It accelerates towards this target, as well as away from the mouse pointer
    in case it is within a certain distance.
---

{{< rawhtml >}}
{{< canvas 880 480 >}}

<p id="instructions">Move the mouse over the logo! Click to pause and resume! {{< fullscreen >}}</p>
<p><tt style="font-size: 120%">go get <a href="https://github.com/mlange-42/arche">github.com/mlange-42/arche</a></tt>
</p>
{{< /rawhtml >}}
