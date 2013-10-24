package main

import  "fmt"
//import "io/ioutil"
//import "log"
import "strings"
import "strconv"
import "bufio"
import "os"

func readTree(filename string) [][]int {
     file, _ := os.Open(filename)
     scanner := bufio.NewScanner(file)
     tree := make([][]int, 0)
     for scanner.Scan() {
          row := make([]int, 0)
          for _, num := range strings.Split(scanner.Text(), " ") {
               n, _ := strconv.Atoi(num)
               row = append(row, n)      
          }
          tree = append(tree, row)
     }
     return tree
}

func maxTotal(tree [][]int) int {
    for i:=len(tree)-2; i >= 0; i-- {
         row := tree[i]
         nextRow := tree[i+1]
         for i, n := range row {
             left := n + nextRow[i]
             right := n + nextRow[i+1]
             if left > right {
                  row[i] = left
             } else {
                  row[i] = right
             }
         }
    }
    return tree[0][0]
}
func main() {
     tree := readTree("67.txt")    
     fmt.Println(maxTotal(tree))
}
