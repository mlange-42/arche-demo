---
title: Home
wasm: home
description: |
    A simple simulation with 30.000 entities.
    
    Each entity has a target pixel in the logo.
    It accelerates towards this target, as well as away from the mouse pointer in case it is closeby.
---

# Arche Demo

{{< rawhtml >}}
<style>
    #loading {
        width: 880px;
        height: 480px;
    }
</style>

<div id="canvas-container">
    <div id="loading">
        <p class="centered">Loading...</p>
    </div>
</div>
<p id="instructions">Move the mouse over the logo! Click to pause and resume!</p>
<p><a href="https://github.com/mlange-42/arche"><tt style="font-size: 120%">github.com/mlange-42/arche</tt></a>
</p>
{{< /rawhtml >}}
