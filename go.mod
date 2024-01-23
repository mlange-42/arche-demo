module github.com/mlange-42/arche-web

go 1.21.6

require (
	github.com/llgcode/draw2d v0.0.0-20231212091825-f55e0c776b44
	github.com/markfarnan/go-canvas v0.0.0-20200722235510-6971ccd00770
	github.com/mlange-42/arche v0.10.0
	github.com/mlange-42/arche-model v0.6.0
)

replace github.com/mlange-42/arche-model v0.6.0 => ../arche-model

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	golang.org/x/exp v0.0.0-20230418202329-0354be287a23 // indirect
	golang.org/x/image v0.14.0 // indirect
)
