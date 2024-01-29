# Arche Demo -- Ant Colony

A stylized model of ants foraging and scouting, using trail pheromones.

Ants lay trails of two types of pheromones:
one when coming from the nest,and another one when coming from a resource.

Using these trails, workers find their way to resources and back to the nest.
To find new resources, scouts swarm out and do a random walk until they find something.

In this example, the landscape is a network with nodes and edges represented by ECS entities.
