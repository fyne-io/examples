<p align="center">
  <a href="https://goreportcard.com/report/github.com/fyne-io/examples"><img src="https://goreportcard.com/badge/github.com/fyne-io/examples" alt="Code Status" /></a>
  <a href='http://gophers.slack.com/messages/fyne'><img src='https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=blue' alt='Join us on Slack' /></a>
  <a href='https://fossfi.sh/support-fyneio'><img src='https://img.shields.io/badge/$-support_us-orange.svg?labelWidth=20&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAkCAYAAADPRbkKAAAABmJLR0QA7wAyAD/CTveyAAAACXBIWXMAAA9hAAAPYQGoP6dpAAAAB3RJTUUH4wMVCQ4LeuPReAAABDFJREFUWMPVmX9oVWUYxz/3tr5jbX9oYiZWSmNgoYYlUVqraYvIamlCg8pqUBTZr38qMhZZ0A8oIqKIiLJoJtEIsR+zki1mGqQwZTYkImtlgk5Wy7jPnbv90XPgcD3n3Lvdy731wOGc+77nec73Oe/zfp/nOTeVy+WolphZC3CVH9N9OAsMAzuBLkkHkmykKu2AmQGsAjYAC4pQ+QJ4WNJg1R0ws3pgI3DTJFVPAM8AGyRNVMUBB98DLCvBzLtAh6QTwUC6gmGzuUTwAGuBR8MD6QpFzwPAypi5v4HXgRXAEg+vLQm2njSzCyoWQmbWAPwAzIqY/gm4UdJAhN5qoAuojdCbAD4BVlViBdbGgJ8A2gPwZtZsZvea2XwASd1AZ4zNNHA90FgJB9ri6FHStw7+WaAPeA3YbWaBzovAiF/nfMUC+QD4sRIOLIkZ73Hws4BHQuOneY7A2aY/CHdgL3AtsBhYD9TVVMCB0+O2h59nRJBJfej6WOj6Bj8C2VcpFoqSC/38PbArb24zQDabBWhOsFFbNAuZ2UygxZcvcHwPsEPScILecaAuYmoUaJR01MymAU8AFzm7vCxp3MxagW0xpv8EzinogLPC40A7cGpMKHQB6yX9FqE/ACyKMb8VWC0pG6E3D+gF5sbo9km6Ml0A/N3Ad8BtMeABBNwB9JrZ7Ij5PQmPuA7YZWZXmFnan1ljZrf65p2boDuQmInNrAN4I29DJUmT02C+bCtiL/QCo2Y2BIwB7wFzCuj1xGZiM7vcb9gOTAMunUTZMVPSkZCtJuBAmQngIHCupImaCPApoBVokvSrj7UDm4o0Pg84Evo97EkoVUYHXgjK6pNWwB1I5dfdZtbtjUghWSppZ0ivDvirjA4MAosCfCeFhaRcPvhQLV6QbYGhvLHZZQRvwO1hfJNJZFu9qkySNyUdyxu7pkzgc8BDknZPuSMzs8XA58AZEdPbgTZJY3ld2GABOixWXpH0YMk9sZlN96zZ6hn2F+AdYFM4IZlZLfBRQiMzGXlaUmdFmnpPSC3A814alCJjwDpJG+NuqCkC0HzgLi+L651RvgTeAu4ELnHazAJnedkwpwzvoh+4RdLPid+FMpnMDElHY8CvAd4GGiKmR4A1DvwpYHkZabLTO7KCkgYeiwG/3Iu0hoQ6fwvwh6QVwFKn2tEpgB4BPgSaJS0oFnywAnuB5yR15bHNV6HPfUnyqaSVIV0BlwEXe51zNtAYehGHgMNOyfu9F+iXND6V5UplMplXgfuAj4HPgIXAPcXsj5CcKelwNbqiVCaTOd/jrhRZJumbajiQlrQf6OZ/KkEpsQ74vQQ7w1V1QNIh4Gbg+BRs9BXi6kqsAJK+Bq52hihWxoH7/wshFDixw786vF9kmm+TtK+aDsTWQma2EOjg379/zgNO8akh73NfknSw2pv4H3Ayg0FmbTMRAAAAAElFTkSuQmCC' alt='Support Fyne.io' /></a>

  <br />
  <a href="https://travis-ci.com/github/fyne-io/examples"><img src="https://travis-ci.com/github/fyne-io/examples.svg" alt="Build Status" /></a>
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
go run main.go -fractal
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

Moved to [calculator repository](https://github.com/fyne-io/calculator/)

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

A fractal viewer that can be panned and zoomed

![](img/fractal.png)

### Solitaire

Moved to [solitaire repository](https://github.com/fyne-io/solitaire/)

### Life

Moved to [life repository](https://github.com/fyne-io/life/)
