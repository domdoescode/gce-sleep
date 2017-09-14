# gce-sleep [![CircleCI](https://circleci.com/gh/domudall/gce-sleep.svg?style=svg)](https://circleci.com/gh/domudall/gce-sleep) [![Average time to resolve an issue](http://isitmaintained.com/badge/resolution/domudall/gce-sleep.svg)](http://isitmaintained.com/project/domudall/gce-sleep "Average time to resolve an issue") [![Percentage of issues still open](http://isitmaintained.com/badge/open/domudall/gce-sleep.svg)](http://isitmaintained.com/project/domudall/gce-sleep "Percentage of issues still open") [![Go Report Card](https://goreportcard.com/badge/github.com/domudall/gce-sleep)](https://goreportcard.com/report/github.com/domudall/gce-sleep)

Tool for shutting down/starting up Google Cloud Engine instances based on tags for savings costs when not in use.

**Current version: v0.0.2**

Concept lovingly stolen and adapted from [gce-bedtime](https://github.com/lbn/gce-bedtime).

## Getting Started

### Installing

Download a version from the [releases](https://github.com/domudall/gce-sleep/releases) page on Github, and place it into your local bin folder.

### Running

#### Config

```hcl
project "gce-sleep-testing" {
	zones = [
		"europe-west1-c"
	]
}

ruleset "weekly-sleep" {
	startTime = "06:00"
	stopTime = "19:00"
	timezone = "Europe/London"
	days = [
		1,
		2,
		3,
		4,
		5
	]
}
```
