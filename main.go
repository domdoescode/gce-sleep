package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

type Config struct {
	Project map[string]Project    `hcl:"project"`
	Ruleset map[string]RawRuleset `hcl:"ruleset"`
}

type Ruleset struct {
	StartTime time.Time
	StopTime  time.Time
	Timezone  *time.Location
	Days      []int
	Instances []Instance
}

type RawRuleset struct {
	StartTime string `hcl:"startTime"`
	StopTime  string `hcl:"stopTime"`
	Timezone  string `hcl:"timezone"`
	Days      []int  `hcl:"days"`
}

type Instance struct {
	Name    string
	Project string
	Zone    string
	Status  string
}

type Project struct {
	Zones []string `hcl:"zones"`
}

func newRuleset(r RawRuleset) (rs Ruleset, err error) {
	timezone, locationErr := time.LoadLocation(r.Timezone)
	if locationErr != nil {
		err = multierror.Append(err, errors.New("timezone is not valid"))
	}

	rs.Timezone = timezone

	if r.StartTime == "" {
		err = multierror.Append(err, errors.New("startTime cannot be blank"))
	}

	if r.StopTime == "" {
		err = multierror.Append(err, errors.New("stopTime cannot be blank"))
	}

	if locationErr == nil {
		startTime, startTimeErr := time.ParseInLocation("15:04", r.StartTime, timezone)
		if startTimeErr != nil {
			err = multierror.Append(err, errors.New("startTime is not in valid 24 hour HH:mm format"))
		}

		rs.StartTime = startTime

		stopTime, stopTimeErr := time.ParseInLocation("15:04", r.StopTime, timezone)
		if stopTimeErr != nil {
			err = multierror.Append(err, errors.New("stopTime is not in valid 24 hour HH:mm format"))
		}

		rs.StopTime = stopTime
	}

	if len(r.Days) > 7 {
		err = multierror.Append(err, errors.New("days must be valid"))
	}

	for _, day := range r.Days {
		if day > 7 || day < 1 {
			err = multierror.Append(err, errors.New("days must be an int between 1 and 7"))
		}
	}
	rs.Days = r.Days

	return
}

func main() {
	var configLocation, labelName string
	flag.StringVar(&configLocation, "c", "/etc/gce-sleep.conf", "Location of the gce-sleep config file")
	flag.StringVar(&labelName, "l", "gce-sleep", "Label for filtering instances to determine if gce-sleep is active")
	flag.Parse()

	configContent, err := ioutil.ReadFile(configLocation)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = hcl.Unmarshal(configContent, &config)
	if err != nil {
		log.Fatal(err)
	}

	activeRules := make(map[string]Ruleset)
	for index, rawRuleset := range config.Ruleset {
		ruleset, err := newRuleset(rawRuleset)
		if err != nil {
			log.Fatal(err)
		} else {
			activeRules[index] = ruleset
		}
	}

	ctx := context.Background()

	client, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(client)
	if err != nil {
		log.Fatal(err)
	}

	for projectName, project := range config.Project {
		for _, zoneName := range project.Zones {
			instancesReq := computeService.Instances.List(projectName, zoneName).Filter(fmt.Sprintf("labels.%s eq on", labelName))
			if err := instancesReq.Pages(ctx, func(page *compute.InstanceList) error {
				for _, instance := range page.Items {
					for _, metadata := range instance.Metadata.Items {
						if metadata.Key == "gce-sleep-group" {
							actionableInstances := activeRules[*metadata.Value]
							actionableInstances.Instances = append(actionableInstances.Instances, Instance{
								Project: projectName,
								Zone:    zoneName,
								Name:    instance.Name,
								Status:  instance.Status,
							})
							activeRules[*metadata.Value] = actionableInstances
						}
					}
				}

				return nil
			}); err != nil {
				log.Fatal(err)
			}
		}
	}

	for _, ruleset := range activeRules {
		now := time.Now().In(ruleset.Timezone)

		shouldBeRunning := shouldBeRunning(now, ruleset.StartTime, ruleset.StopTime)

		for _, instance := range ruleset.Instances {
			if shouldBeRunning && instance.Status == "TERMINATED" {
				call := computeService.Instances.Start(instance.Project, instance.Zone, instance.Name)
				call.Do()

				log.Println(fmt.Sprintf("Instance %q starting", instance.Name))
			} else if !shouldBeRunning && instance.Status == "RUNNING" {
				call := computeService.Instances.Stop(instance.Project, instance.Zone, instance.Name)
				call.Do()

				log.Println(fmt.Sprintf("Instance %q stopping", instance.Name))
			}
		}
	}
}

func shouldBeRunning(now, startTime, stopTime time.Time) bool {
	if startTime.Hour() <= now.Hour() && stopTime.Hour() >= now.Hour() {
		if startTime.Hour() == now.Hour() && startTime.Minute() > now.Minute() {
			return false
		}

		if stopTime.Hour() == now.Hour() && stopTime.Minute() < now.Minute() {
			return false
		}
	}

	return true
}
