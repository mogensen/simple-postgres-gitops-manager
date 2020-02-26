package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	for {
		runReconsilition()
		time.Sleep(3 * time.Second)
	}
}

func runReconsilition() {
	log.Infof("============")

	// Read Desired State
	desiredState, err := newStateFromFile("desired-state.yml")
	if err != nil {
		log.Errorf("Could not read desired state: %v", err)
		return
	}
	// Read Current State
	db, err := connect()
	if err != nil {
		log.Errorf("Could not read current state: %v", err)
		return
	}
	defer db.Close()
	currentState, err := newStateFromDB(db)
	if err != nil {
		log.Errorf("Could not read current state: %v", err)
		return
	}

	// Diff
	log.Infof("Desired state: %v", desiredState)
	log.Infof("Current state: %v", currentState)

	shouldCreate := findDatabasesToCreate(desiredState, currentState)
	shouldDelete := findDatabasesToDelete(desiredState, currentState)

	log.Infof("Planning to create : %v", shouldCreate)
	log.Infof("Planning to delete : %v", shouldDelete)

	// Act
	for _, dbToCreate := range shouldCreate {
		err := createDB(db, dbToCreate)
		if err != nil {
			log.Warnf("Failed to create db: %v", dbToCreate)
		}
	}
	for _, dbToDelete := range shouldDelete {
		err := deleteDB(db, dbToDelete)
		if err != nil {
			log.Warnf("Failed to delete db: %v", dbToDelete)
		}
	}
}

func findDatabasesToCreate(desiredState, currentState *State) []Database {
	res := []Database{}
	for _, desired := range desiredState.Databases {
		found := findByName(currentState, desired.Name)
		if !found {
			res = append(res, desired)
		}
	}
	return res
}

func findDatabasesToDelete(desiredState, currentState *State) []Database {
	res := []Database{}
	for _, current := range currentState.Databases {
		found := findByName(desiredState, current.Name)
		if !found {
			res = append(res, current)
		}
	}
	return res
}

func findByName(state *State, name string) bool {
	for _, db := range state.Databases {
		if db.Name == name {
			return true
		}

	}
	return false
}
