package main

import "fmt"

type NodeG struct {
  next *NodeG
  val int
}

type Graph struct {
  vertex []*NodeG
  edgeCount map[Edge]int
  numEdges int
}

type Edge struct {
  s int
  d int
}

// ===== Merge Sort =====
func mergeSort(s []string) []string {
    if len(s) <= 1 {
        return s
    }
    
    l := mergeSort(s[:len(s)/2])
    r := mergeSort(s[len(s)/2:])
    return merge2(l, r)
}

func merge2(l, r []string) []string {
    merged := make([]string, len(l) + len(r))
    
    li:=0
    ri:=0
    
    for i := 0 ; i<len(merged) ; i++ {
        if li == len(l) {
            merged[i] = r[ri]
            ri++
        } else if ri == len(r) {
            merged[i] = l[li]
            li++
        } else if l[li] < r[ri] {
            merged[i] = l[li]
            li++
        } else {
            merged[i] = r[ri]
            ri++
        }
    }
    
    return merged
}

// ===== CONVERSION =====

func getListOfCityFromTickets(tickets [][]string) []string {
    // stores the cities in a hash map
    cityMap := make(map[string]bool)
    for i := 0 ; i<len(tickets) ; i++ {
        cityMap[tickets[i][0]] = true
        cityMap[tickets[i][1]] = true
    }
    
    // grabs the list of cities in a hash
    cityList := make([]string, 0)
    cityI := 0
    for k, _ := range(cityMap) {
        cityList = append(cityList, k)
        cityI++
    }
    return cityList
}

// convert a list of tickets to a list of ints to be fed into the graph.
func convertCitiesToInt(tickets [][]string, m map[string]int) [][]int {
  ticketInts := make([][]int, len(tickets))
  for i := 0 ; i<len(tickets) ; i++ {
    srcI := m[tickets[i][0]]
    destI := m[tickets[i][1]]
    ticketInts[i] = []int{srcI, destI}
  }
  return ticketInts
}

// return 2 hashes that are based off a sorted list...
// should get you any easy conversion mechanism between ints and 
func getOrderCitiesHash(cities []string) (map[string]int, map[int]string) {
  citiesList := mergeSort(cities)
  citiesIntHash := make(map[string]int)
  intCitiesHash := make(map[int]string)

  for k, v := range citiesList {
      citiesIntHash[v] = k
      intCitiesHash[k] = v
  }

  return citiesIntHash, intCitiesHash
}

func convertEdgeIntsToListOfCities(cities [][]int, m map[int]string) []string {
  listCities := make([]string, 0)

  for i := 0 ; i < len(cities) ; i++ {
    if i == 0 {
      listCities = append(listCities, m[cities[i][0]])
    }
    listCities = append(listCities, m[cities[i][1]])
  }

  return listCities
}

// ===== ONE OFF =====

// Do you need this..
func copyMap(s map[Edge]int) map[Edge]int {
  d := make(map[Edge]int)
  for k,v := range(s) {
    d[k] = v
  }
  return d
}

// ===== GRAPH =====
func (g *Graph) init(numVertex int) {
  if g.vertex == nil {
    g.vertex = make([]*NodeG, numVertex)
  }
  g.edgeCount = make(map[Edge]int)
}

func (g *Graph) insert(adjList [][]int) {
  // storing the number of edges.
  g.numEdges = len(adjList)

  for i:=0 ; i<len(adjList) ; i++ {
    curEdge := Edge{adjList[i][0], adjList[i][1]}

    if g.vertex[adjList[i][0]] == nil {
      g.vertex[adjList[i][0]] = &NodeG{nil, adjList[i][1]}      
      g.edgeCount[curEdge]++
    } else {
      if _, ok := g.edgeCount[curEdge]; ok {
        g.edgeCount[curEdge]++
      } else {
        oldHead := g.vertex[adjList[i][0]]
        newNode := &NodeG{nil, adjList[i][1]}
        // insert node at the start.
        (*newNode).next = oldHead
        g.vertex[adjList[i][0]] = newNode

        // bubble the value up to the correct position
        prevNode := newNode
        prevNode = nil
        for curNode := g.vertex[adjList[i][0]] ; (*curNode).next != nil ; curNode = (*curNode).next {
          nextNode := (*curNode).next
          if (*curNode).val > (*nextNode).val && curNode == g.vertex[adjList[i][0]] {
            g.vertex[adjList[i][0]] = nextNode
            (*curNode).next = (*nextNode).next
            (*nextNode).next = curNode
            curNode = nextNode
          } else if (*curNode).val > (*nextNode).val {
            (*prevNode).next = nextNode
            (*curNode).next = (*nextNode).next
            (*nextNode).next = curNode
            curNode = nextNode
          } else {
            break
          }
          // set the previous node.
          prevNode = curNode

          // not sure why you need this.
          if (*curNode).next == nil {
            break
          }
        }
        //(*n).next = &NodeG{nil, adjList[i][1]}
        g.edgeCount[curEdge]++
      }
    }
  }
}

// you need to update...
func (g *Graph) adj(v int) []int {
  if g.vertex[v] == nil {
    return nil
  } else {
    adjNodes := make([]int, 0)
    n := g.vertex[v]
    for {
      adjNodes = append(adjNodes, (*n).val)
      n = n.next
      if n == nil {
        break
      }
      if n.next == nil {
        adjNodes = append(adjNodes, (*n).val)
        break
      }
    }
    return adjNodes
  }
}

func (g *Graph) findFirstDFS(curVertex int) [][] int {
  finalPath := make([][]int, 0)
  // not sure if you need this.
  existingEdges := copyMap(g.edgeCount)
  done := false
  g.findFirstDFS2(curVertex, &finalPath, existingEdges, &done)
  fmt.Println("> findFirstDFS", finalPath)

  return finalPath
}

func (g *Graph) findFirstDFS2(curVertex int, finalPath *[][]int, existingEdges map[Edge]int, done *bool) {
  fmt.Println("\n> findFirstDFS2 > start", len(*finalPath), "numEdges", g.numEdges, "existingEdges", existingEdges)

  if len(*finalPath) == g.numEdges && !(*done) {
    (*done) = true
  } else if (!(*done)) {
    // figure out legal adj
    curAdjVertex := g.adj(curVertex)
    for i := 0 ; i < len(curAdjVertex) ; i++ {
      curEdge := Edge{curVertex, curAdjVertex[i]}
      v, ok := existingEdges[curEdge]
      if ok && v != 0 {
        // update final path
        if !(*done) {
          (*finalPath) = append((*finalPath), []int{curEdge.s, curEdge.d})
        }

        fmt.Println(">findFirstDFS2 > finalPath before traversal", (*finalPath))
        existingEdges[curEdge]--
        fmt.Println(">findFirstDFS2 > traversing to > curEdge", curEdge)
        g.findFirstDFS2(curEdge.d, finalPath, existingEdges, done)
        existingEdges[curEdge]++

        // pop off the last item.
        if !(*done) {
          (*finalPath) = (*finalPath)[:len((*finalPath))-1]
        }
      }
    }
  }

  // fmt.Println(finalPath)
  // fmt.Println(existingEdges)
  // fmt.Println(*done)
}

// ===== MAIN =====
func findItinerary(tickets [][]string) []string {
  var g Graph
  
  listOfCities := getListOfCityFromTickets(tickets)
  citiesIntHash, intCityHash := getOrderCitiesHash(listOfCities)
  ticketInts := convertCitiesToInt(tickets, citiesIntHash)
  g.init(len(listOfCities))
  g.insert(ticketInts)
  jfkInt := citiesIntHash["JFK"]

  // find DFP
  resultInt := g.findFirstDFS(jfkInt)
  
  // convert the values back.
  resultCities := convertEdgeIntsToListOfCities(resultInt, intCityHash)

  // PRINT TO NOT GET AN ERROR.
  fmt.Println("\n>>> FIND ITEN")
  // fmt.Println(citiesIntHash)
  //fmt.Println(intCityHash)
  // fmt.Println(ticketInts)
  fmt.Println(resultInt)
  fmt.Println(resultCities)

  return resultCities
}

func main() {
  var g Graph
  
  // ["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]
  // g.init(5)
  // tickets := make([][]string, 4)
  // tickets[0] = []string{"MUC", "LHR"}
  // tickets[1] = []string{"JFK", "MUC"}
  // tickets[2] = []string{"SFO", "SJC"}
  // tickets[3] = []string{"LHR", "SFO"}

  // [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
  g.init(3)
  tickets := make([][]string, 5)
  tickets[0] = []string{"JFK","SFO"}
  tickets[1] = []string{"JFK","ATL"}
  tickets[2] = []string{"SFO","ATL"}
  tickets[3] = []string{"ATL","JFK"}
  tickets[4] = []string{"ATL","SFO"}

  // [["EZE","AXA"],["TIA","ANU"],["ANU","JFK"],["JFK","ANU"],["ANU","EZE"],["TIA","ANU"],["AXA","TIA"],["TIA","JFK"],["ANU","TIA"],["JFK","TIA"]]
  // ["JFK","ANU","EZE","AXA","TIA","ANU","JFK","TIA","ANU","TIA","JFK"] 
  // g.init(5)
  // tickets := make([][]string, 10)
  // tickets[0] = []string{"EZE","AXA"}
  // tickets[1] = []string{"TIA","ANU"}
  // tickets[2] = []string{"ANU","JFK"}
  // tickets[3] = []string{"JFK","ANU"}
  // tickets[4] = []string{"ANU","EZE"}
  // tickets[5] = []string{"TIA","ANU"}
  // tickets[6] = []string{"AXA","TIA"}
  // tickets[7] = []string{"TIA","JFK"}
  // tickets[8] = []string{"ANU","TIA"}
  // tickets[9] = []string{"JFK","TIA"}

  // [["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]
  // ["JFK","NRT","JFK","KUL"]
  // g.init(3)
  // tickets := make([][]string, 3)
  // tickets[0] = []string{"JFK","KUL"}
  // tickets[1] = []string{"JFK","NRT"}
  // tickets[2] = []string{"NRT","JFK"}
  

  // and the end
  findItinerary(tickets)

  // myInts := make([][]int, 8)
  // myInts[0] = []int{0, 4}
  // myInts[1] = []int{0, 2}
  // myInts[2] = []int{0, 3}
  // myInts[3] = []int{0, 20}
  // myInts[4] = []int{0, 15}
  // myInts[5] = []int{0, 100}
  // myInts[6] = []int{0, 3}
  // myInts[7] = []int{0, 1}
  // myInts[2] = []int{0, 1}
  // myInts[3] = []int{0, 2}
  //g.insert(myInts)

  // fmt.Println(">", g.vertex)
  //fmt.Println(g.adj(0))
  // fmt.Println(g.edgeCount)
}

/*

80 test cases on leetcode for 58 ms
so its pretty good...
.
Optimizations.
Using a double hash to go from int to strings and vice versa in O(1)
Sorted the adjancy list  adjacent edges for the lex minimal ones goes first. Non need to sort by picking the adj edge.


*/