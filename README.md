<p align="center">
  <a href="https://goreportcard.com/report/github.com/fyne-io/examples"><img src="https://goreportcard.com/badge/github.com/fyne-io/examples" alt="Code Status" /></a>
  <a href="https://travis-ci.org/fyne-io/examples"><img src="https://travis-ci.org/fyne-io/examples.svg" alt="Build Status" /></a>
  <a href='https://coveralls.io/github/fyne-io/examples?branch=develop'><img src='https://coveralls.io/repos/github/fyne-io/examples/badge.svg?branch=develop' alt='Coverage Status' /></a>
</p>

# Fyne Examples

Here we will gather example apps that use the [Fyne](http://fyne.io) toolkit.

You can start the main example app that links to all the others by running

```bash 
go run main.go
```

or you can specify a particular example by naming it in the parameter list, like:

```bash
go run main.go -calculator
```

Alternatively each app has a direct main executable in the cmd/* folders.

All these examples are fully scalable - try setting the `FYNE_SCALE`
environment variable to override the detection of your screen's density.
Many also respond to the current theme (this is default behaviour for
apps built using Fyne widgets) - you can try setting `FYNE_THEME=light`
to change from the default dark theme.

## Widget based examples

The following examples use mostly built in widgets making applications
trivial to build :).

### Calculator

![](img/calc-dark.png) &nbsp; ![](img/calc-light.png)

### Bugs game (like MineSweeper)

Hunt the squares to reveal everything apart from the bugs!

![](img/bugs.png)

### XKCD

An XKCD comic browser with random and lookup features.

![](img/xkcd.png)

## Graphics based examples

These examples use the Fyne canvas API to draw primitive shapes,
text and images to create custom user interfaces.

### Clock

A simple analog clock that matches the current theme.

![](img/clock-dark.png) &nbsp; ![](img/clock-light.png)

### Fractal

A fratal viewer that can be panned and zoomed

![](img/fractal.png)

### Solitaire

A simple game of solitaire.

![](img/solitaire.png)

### Life

A basic visualisation of Conway's game of life.

![](img/life.png)
