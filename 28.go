package main

import "fmt"

func main() {
    const gridSize int = 1001
    var grid [gridSize][gridSize]int

    middle := gridSize/2

    upLeftX := middle-1
    upLeftY := middle-1
    downRightX := middle+1
    downRightY := middle+1
    curX := middle
    curY := middle

    const (
        left int = iota
        right
        up
        down
    )


    direction := right
    i := 1
    sum := 0
    for i <= gridSize*gridSize {
        grid[curX][curY] = i
        if curX == curY || (curX + curY) == gridSize-1{
            sum += i
        }
        i++
        switch (direction) {
            case right:
                    curY++
                    if (curY >= downRightY) {
                        direction = down
                    }
            case down:
                    curX++
                    if (curX >= downRightX) {
                        direction = left
                    }
            case left:
                    curY--
                    if (curY <=upLeftY) {
                        direction = up
                    }
            case up:
                    curX--
                    if (curX <= upLeftX) {
                        direction = right
                        upLeftX--
                        upLeftY--
                        downRightX++
                        downRightY++
                    }
        }
    }
    fmt.Println(sum)

}
