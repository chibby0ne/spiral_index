package main

import (
    "testing"
)

func TestGetStartingNumberFromRing(t *testing.T) {
    tables := []struct {
        ringNumber int
        expectedValue int
    }{
        {0, 0},
        {1, 1},
        {2, 9},
        {3, 25},
    }
    for _, table := range tables {
        actualValue := getStartingValueFromRing(table.ringNumber)
        if actualValue != table.expectedValue {
            t.Errorf("ExpectedValue: %v, ActualValue: %v\n", table.expectedValue, actualValue)
        }
    }
}


func TestGetRingNumberFromCoordinate(t *testing.T) {
    tables := []struct {
        coord *coordinates
        expectedValue int
    }{
        {&coordinates{0, 0}, 0},
        {&coordinates{1, 0}, 1},
        {&coordinates{2, -1}, 2},
        {&coordinates{3, -2}, 3},
    }
    for _, table := range tables {
        actualValue := getRingNumberFromCoordinate(table.coord)
        if actualValue != table.expectedValue {
            t.Errorf("ExpectedValue: %v, ActualValue: %v\n", table.expectedValue, actualValue)
        }
    }
}


func TestGetInitialCoordinatesFromRing(t *testing.T) {
    tables := []struct {
        ringNumber int
        expectedValue *coordinates
    }{
        {0, &coordinates{0, 0}},
        {1, &coordinates{1, 0}},
        {2, &coordinates{2, -1}},
        {3, &coordinates{3, -2}},
    }
    for _, table := range tables {
        actualValue := getInitialCoordinatesFromRing(table.ringNumber)
        if !isSameCoordinate(actualValue, table.expectedValue) {
            t.Errorf("ExpectedValue: %v, ActualValue: %v\n", *table.expectedValue, *actualValue)
        }
    }
}

func TestCalculateValueFromCoordinate(t *testing.T) {
    tables := []struct {
        coord *coordinates
        expectedValue int
    }{
        {&coordinates{0, 0}, 0},
        {&coordinates{1, 0}, 1},
        {&coordinates{1, 2}, 13},
        {&coordinates{2, -1}, 9},
        {&coordinates{2, -2}, 24},
        {&coordinates{3, -2}, 25},
        {&coordinates{3, 2}, 29},
    }
    for _, table := range tables {
        actualValue := calculateValueFromCoordinate(table.coord)
        if actualValue != table.expectedValue {
            t.Errorf("ExpectedValue: %v, ActualValue: %v\n", table.expectedValue, actualValue)
        }
    }
}

