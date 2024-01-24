---
title: Home
wasm: home
description: |
    Test description.

    With multiple lines and a [Markdown link](https://test.com).
---

# Arche Demo

{{< rawhtml >}}
<style>
    #canvas-container {
        display: inline-block;
        border: solid white 2px;
        min-width: 0px;
        max-width: 95%;
        overflow: auto;
    }

    #loading {
        width: 880px;
        height: 480px;
        background-color: #000000;
        color: #ffffff;
    }

    #canvas {
        max-width: 100%;
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