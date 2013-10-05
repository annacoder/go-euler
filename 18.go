package main

import  "fmt"
import "io/ioutil"
import "log"
import "strings"
import "strconv"

type Tree struct {
    Left *Tree
    Value int
    Right *Tree
}

func maxTotal(tree *Tree) int {
    if tree.Left == nil {
        return tree.Value
    }

    leftTotal := maxTotal(tree.Left)
    rightTotal := maxTotal(tree.Right)
    if leftTotal > rightTotal {
        return leftTotal + tree.Value
    }
    return rightTotal + tree.Value
}

func main() {
    content, err := ioutil.ReadFile("18.txt")
    if err!=nil {
        log.Fatal("couldn't read file", err)
    }

    tree := new(Tree)
    prevLevel := make([]*Tree, 0)
    lines := strings.Split(string(content), "\n")
    for i,line := range lines {
        if line == "" { continue }
        numbers := make([]int,0)
        for _, num := range strings.Split(line, " ") {
            n, _ := strconv.Atoi(num)
            numbers = append(numbers,n)
        }
        if i == 0 {
            tree.Value = numbers[0]
            prevLevel = append(prevLevel, tree)
            continue
        }

        newLevel := make([]*Tree, 0)
        for _, n := range numbers {
            newLevel = append(newLevel, &Tree{Value: n})
        }

        for j, node := range prevLevel {
            node.Left = newLevel[j]
            node.Right = newLevel[j+1]
        }
        prevLevel = newLevel
    }
    fmt.Println(maxTotal(tree))
}

