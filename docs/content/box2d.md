---
title: Box2D
description: |
    A simulation using the [Go port](https://github.com/ByteArena/box2d) of [Box2D](https://box2d.org/),
    a 2D physics engine for games.

    Here, [Arche](https://github.com/mlange-42/arche) only handles the graphics
    and applies forces from the mouse.
---

{{< rawhtml >}}
<style>
    #loading {
        width: 880px;
        height: 480px;
    }
</style>

<div id="canvas-container">
    <iframe id="iframe" src="/wasm.html?box2d" width="880" height="480" allow="autoplay" frameBorder="0" scrolling="no"></iframe>
</div>
{{< /rawhtml >}}
