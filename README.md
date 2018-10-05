# Fyne Examples

Here we will gather example apps that use the [Fyne](http://fyne.io) toolkit.
This is a new repository and many existing examples can still be found
in the main Fyne repo at [https://github.com/fyne-io/fyne/tree/develop/examples].

You can start the main example app by running 

```bash 
go run main.go
```

or you can specify a particular exapmple by naming it in the parameter list, like:

```bash
go run main.go -calculator
```

Alternatively each app has a direct main executable in the cmd/* folders.

All these examples are fully scalable - try setting the `FYNE_SCALE`
environment variable to override the detection of your screen's density.

## Calculator

![](img/calc-linux-dark.png) &nbsp; ![](img/calc-linux-light.png)


## Fractal

![](img/fractal-dark.png)

## Solitaire

This is a work in progress

![](img/solitaire.png)


