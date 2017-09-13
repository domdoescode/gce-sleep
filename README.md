# gce-sleep

Tool for shutting down/starting up Google Cloud Engine instances based on tags for savings costs when not in use.

**Current version: v0.0.0**

Concept lovingly stolen and adapted from (https://github.com/lbn/gce-bedtime)[gce-bedtime].

## Getting Started

### Installing

Download a version from the (https://github.com/domudall/gce-sleep/releases)[releases] page on Github, and place it into your local bin folder.

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
