package main

import (
    "fmt"
    "flag"
    "math"
)

type coordinates struct {
    x, y int
}


func parseArguments() *coordinates {
    x := flag.Int("x", 3, "X coordinate")
    y := flag.Int("y", 2, "Y coordinate")
    flag.Parse()
    coords := new(coordinates)
    coords.x = *x
    coords.y = *y
    return coords
}

func getStartingValueFromRing(ringNumber int) int {
    // If the highest dimension is 0 it is at the origin point
    if ringNumber == 0 {
        return 0
    }
    // Starts with one point whose value is 0, therefore first ring starts with
    // dotsInPreviousRing = 1, whose previousRingLastValue = 0 and innerDots = 1
    dotsInPreviousRing := 1
    innerDots := 1
    previousRingLastValue := 0
    for i := 1; i < ringNumber; i++ {
        dotsInPreviousRing = 4 * (innerDots + 1) 
        previousRingLastValue += dotsInPreviousRing
        innerDots += 2
    }
    // This is the starting value of the ring, 
    return previousRingLastValue + 1
}

func getInitialCoordinatesFromRing(ringNumber int) *coordinates {
    // If the highest dimension is 0 it is at the origin point
    if ringNumber == 0 {
        return &coordinates{
            x: 0,
            y: 0,
        }
    }
    // Start from 1, 0 ring 1
    coord := &coordinates{
        x: 1,
        y: 0,
    }
    for i := 1; i < ringNumber; i++ {
        coord.x = coord.x + 1
        coord.y = coord.y - 1
    }
    return coord
}

func getRingNumberFromCoordinate(coord *coordinates) int {
    var ringNumber int
    xCoord := int(math.Abs(float64(coord.x)))
    yCoord := int(math.Abs(float64(coord.y)))
    if xCoord > yCoord {
        ringNumber = xCoord
    } else {
        ringNumber = yCoord
    }
    return ringNumber
}

func isSameCoordinate(coordA, coordB *coordinates) bool {
    if coordA.x == coordB.x && coordA.y == coordB.y {
        return true
    }
    return false
}

func calculateValueFromCoordinate(coord *coordinates) int {
    ringNumber := getRingNumberFromCoordinate(coord)
    startValue := getStartingValueFromRing(ringNumber)
    initCoord := getInitialCoordinatesFromRing(ringNumber)
    addValue := 0
    currentCoord := initCoord

    // right side, from starting pos to right upper corner
    for j := initCoord.y; j <= ringNumber; j++ {
        currentCoord.y = j
        if isSameCoordinate(coord, currentCoord) {
            return startValue + addValue
        }
        addValue++
    }

    // top side, from right upper corner to left upper corner
    for i := initCoord.x - 1; i >= -1 * ringNumber; i-- {
        currentCoord.x = i
        if isSameCoordinate(coord, currentCoord) {
            return startValue + addValue
        }
        addValue++
    }

    // left side, from left upper corner to left lower corner
    for j := currentCoord.y - 1; j >= -1 * ringNumber; j-- {
        currentCoord.y = j
        if isSameCoordinate(coord, currentCoord) {
            return startValue + addValue
        }
        addValue++
    }

    // bottom side, from left bottom corner to right bottom corner
    for i := currentCoord.x + 1; i <= ringNumber; i++ {
        currentCoord.x = i
        if isSameCoordinate(coord, currentCoord) {
            return startValue + addValue
        }
        addValue++
    }
    return 0
}

func main() {
    coords := parseArguments()
    fmt.Printf("Using coordinates x=%v, y=%v\n", coords.x, coords.y)
    fmt.Println(calculateValueFromCoordinate(coords))
}
