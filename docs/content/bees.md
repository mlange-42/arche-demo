---
title: Bee Hives
description: |
    A stylized model of bee foraging and scouting.

    Bees scout ({{< col yellow >}}yellow{{< /col >}}),
    forage ({{< col "#e0e0e0" >}}light grey{{< /col >}}),
    return to the hive ({{< col cyan >}}cyan{{< /col >}})
    and potentially do the waggle dance ({{< col "#ff7070" >}}red{{< /col >}}).
    Other bees in the hive ({{< col "#b0b0b0" >}}dark grey{{< /col >}}) may fly to a patch indicated by a waggle dance (white),
    or go scouting on their own.

    Uses different components to indicate the different activities.
    Hive internal decision making uses Arche's entity relations feature.
---

{{< rawhtml >}}
{{< canvas bees 880 480 >}}

<p id="instructions">Move the mouse over the canvas! Click to pause and resume!</p>
<p class="tt">go get <a href="https://github.com/mlange-42/arche">github.com/mlange-42/arche</a>
</p>
{{< /rawhtml >}}
