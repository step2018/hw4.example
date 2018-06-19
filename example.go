// An example Go program that simply reads the contents of the graph
// data files and stores them in memory using a simple map[int][]int
// for an adjacency list.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func readEdges() map[int][]int {
	f, err := os.Open("links.txt")
	if err != nil {
		panic(err)
	}
	graph := map[int][]int{}
	scanner := bufio.NewScanner(f)
	fmt.Println("Memory usage prior to reading graph edges")
	PrintMemUsage()
	startTime := time.Now()
	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.SplitN(line, "\t", 2)
		if len(nodes) < 2 {
			log.Fatalf("bad line %v", line)
		}
		var from, to int
		if from, err = strconv.Atoi(nodes[0]); err != nil {
			log.Fatalf("bad 'from' value in line %q: %v", line, err)
		}
		if to, err = strconv.Atoi(nodes[1]); err != nil {
			log.Fatalf("bad 'to' value in line %q: %v", line, err)
		}
		graph[from] = append(graph[from], to)
	}
	endTime := time.Now()
	fmt.Println("Finished reading edges in", endTime.Sub(startTime))
	fmt.Println("Memory usage after reading graph edges")
	PrintMemUsage()
	return graph
}

func readNames() []string {
	f, err := os.Open("pages.txt")
	if err != nil {
		panic(err)
	}
	names := []string{}
	scanner := bufio.NewScanner(f)
	fmt.Println("Memory usage prior to reading graph node names")
	PrintMemUsage()
	startTime := time.Now()
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) < 2 {
			log.Fatalf("bad line: %v", line)
		}
		names = append(names, parts[1])
	}
	endTime := time.Now()
	fmt.Println("Finished reading node names in", endTime.Sub(startTime))
	fmt.Println("Memory usage after reading graph node names")
	PrintMemUsage()
	return names
}

func main() {
	graph := readEdges()
	names := readNames()
	// TODO: do something interesting with graph + names.
	_, _ = graph, names
}

// From https://golangcode.com/print-the-current-memory-usage/ (MIT
// license). PrintMemUsage outputs the current, total and OS memory
// being used. As well as the number of garage collection cycles
// completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
