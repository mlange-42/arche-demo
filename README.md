# Arche Demo

[![Test status](https://img.shields.io/github/actions/workflow/status/mlange-42/arche-demo/tests.yml?branch=main&label=Tests&logo=github)](https://github.com/mlange-42/arche-demo/actions/workflows/tests.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mlange-42/arche-demo.svg)](https://pkg.go.dev/github.com/mlange-42/arche-demo)
[![GitHub](https://img.shields.io/badge/github-repo-blue?logo=github)](https://github.com/mlange-42/arche-demo)
[![MIT license](https://img.shields.io/github/license/mlange-42/arche-demo)](https://github.com/mlange-42/arche-demo/blob/main/LICENSE)

Demo for the [Arche](https://github.com/mlange-42/arche) Entity Component System (ECS).

See the [live demo here](https://mlange-42.github.io/arche-demo/).

<div align="center">

<a href="https://github.com/mlange-42/arche">
<img src="https://user-images.githubusercontent.com/44003176/236701164-28178d13-7e52-4449-baa4-41b764183cbd.png" alt="Arche (logo)" width="500px" />
</a>

</div>

## Usage

Besides the [live demo](https://mlange-42.github.io/arche-demo/), you can run the examples locally.

Clone the repository, and navigate into it:

```
git clone https://github.com/mlange-42/arche-demo.git
cd arche-demo
```

Then, run individual examples like this:

```
go run . <example>
```

## Dependencies

Due to the use of [Ebitengine](https://github.com/hajimehoshi/ebiten) for rendering, the dependencies of [go-gl/gl](https://github.com/go-gl/gl) and [go-gl/glfw](https://github.com/go-gl/glfw) apply. For Ubuntu/Debian-based systems, these are:

- `libgl1-mesa-dev`
- `xorg-dev`
