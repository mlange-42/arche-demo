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

{{< canvas box2d 880 480 >}}

<p id="instructions">Move the mouse over the canvas! Click to pause and resume!</p>
<p class="tt">go get <a href="https://github.com/mlange-42/arche">github.com/mlange-42/arche</a>
</p>
{{< /rawhtml >}}
