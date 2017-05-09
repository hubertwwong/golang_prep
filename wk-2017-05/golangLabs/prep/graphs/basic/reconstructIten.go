package main

import "fmt"

// ADJ LIST
// ======================================================
type NodeG struct {
  next *NodeG
  val int
}

type Graph struct {
  vertex []*NodeG
}

// constrct a node.
// assumes that its linear. in the vertex labeling...
func (g *Graph) init(numVertex int) {
  if g.vertex == nil {
    g.vertex = make([]*NodeG, numVertex+1)
  }
}

// insert stuff into the Graph
func (g *Graph) insert(adjList [][]int) {
  //fmt.Println(">insert")
  for i:=0 ; i<len(adjList) ; i++ {
    if g.vertex[adjList[i][0]] == nil {
      //fmt.Println(">insert new")
      g.vertex[adjList[i][0]] = &NodeG{nil, adjList[i][1]}
    } else {
      fmt.Println(">insert old")
      n := g.vertex[adjList[i][0]]
      for ; n.next != nil ; {
        n = n.next
      }
      (*n).next = &NodeG{nil, adjList[i][1]}
    }
  }
}

// return all adj vertex as an array of ints if you want to process it.
// so if you hit a cycle...
func (g *Graph) adj(v int) []int {
  //fmt.Println("> g adj", v, g.vertex[v])
  if g.vertex[v] == nil {
    return nil
  } else {
    adjNodes := make([]int, 0)
    n := g.vertex[v]
    for {
      //fmt.Println(">>>>>", (*n))
      adjNodes = append(adjNodes, (*n).val)
      //fmt.Println(">1", n)
      n = n.next
      //fmt.Println(">2", n)
      if n == nil {
        break
      }
      if n.next == nil {
        adjNodes = append(adjNodes, (*n).val)
        break
      }
    }
    //fmt.Println("> adjNodes", adjNodes)
    return adjNodes
  }
}

// find all graphs given a specific vertex...
// returns a 3d array...
// 1st one contains the path.
func (g *Graph) findAllPaths(v int) [][][]int {
  curPath := make([][]int, 0)
  paths := make([][][]int, 0)
  g.findPathDFS(v, curPath, &paths)

  //fmt.Println(">paths ")
  //fmt.Println(paths)

  return paths
}

// return a dfs path.
func (g *Graph) findPathDFS(v int, curPath [][]int, paths *[][][]int) {
  // append current node to path
  // curPath = append(curPath, v)
  
  // find all adj verticies
  adj := g.adj(v)
  fmt.Println("> findPathDFS", v, "adj", adj, "curPath", curPath)
  // if there are no adj verticies, you are at the end of a path.
  if len(adj) == 0 && len(curPath) == len(g.vertex){
    (*paths) = append((*paths), curPath)
  } else {
    // flag to mark if you have taken a path
    pathTaken := false

    // dfs for each node. need to check if you haven't taken a path yet.
    for i:=0 ; i<len(adj) ; i++ {
      curEdge := []int{v, adj[i]}
      fmt.Println(intInList(curEdge, curPath))
      if intInList(curEdge, curPath) == false {
        curPath = append(curPath, curEdge)
        g.findPathDFS(adj[i], curPath, paths)
        pathTaken = true
      }
    }

    // if you did not take any path, you are done constructing a possible path.
    if pathTaken == false && len(curPath) == len(g.vertex) {
      (*paths) = append((*paths), curPath)
    }
  }
}

// returns if a edge is in the edge list.
func intInList(v []int, l [][]int) bool {
  for i:=0 ; i<len(l) ; i++ {
    if v[0] == l[i][0] && v[1] == l[i][1] {
      return true
    }
  }
  return false
}

// SORT FN STRING
// ======================================================
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

// SORT FN INTS. Wonder if you can use an interface.
// ======================================================
func mergeSortI(s []int) []int {
    if len(s) <= 1 {
        return s
    }
    
    l := mergeSortI(s[:len(s)/2])
    r := mergeSortI(s[len(s)/2:])
    return merge2I(l, r)
}

func merge2I(l, r []int) []int {
    merged := make([]int, len(l) + len(r))
    
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

// Return a sorted list of cities.
func orderedVertexList(tickets [][]string) []string {
    // stores the cities in a hash map
    cityMap := make(map[string]bool)
    for i := 0 ; i<len(tickets) ; i++ {
        cityMap[tickets[i][0]] = true
        cityMap[tickets[i][1]] = true
    }
    
    // grabs the list of cities in a hash
    cityList := make([]string, 0)
    cityI := 0
    //fmt.Println("<")
    for k, _ := range(cityMap) {
        cityList = append(cityList, k)
        //fmt.Println(cityList[cityI])
        cityI++
    }
    
    return mergeSort(cityList)
}

// Convert a list of cities to a list of ints.
func convertCityToInts(tickets [][]string, cities []string) [][]int {
    ticketsInt := make([][]int, len(tickets))
    
    for i:=0 ; i<len(tickets) ; i++ {
        srcCityI := getCityIndex(tickets[i][0], cities)
        destCityI := getCityIndex(tickets[i][1], cities)
        ticketsInt[i] = []int{srcCityI, destCityI}
    }
    
    return ticketsInt
}

// helper to convert city paris to int pairs...
func getCityIndex(s string, sList []string ) int {
    for k, v := range(sList) {
        if v == s {
            return k
        }
    }
    return -1
}

// find a src dest pair in the ticket its array.
func findIndexInTicketInts(src, dest int, ticketsInts [][]int) int {
  index := -1

  if src < 0 || src > len(ticketsInts) {
    return -1
  } else if dest < 0 || dest > len(ticketsInts) {
    return -1
  }

  for i:=0 ; i<len(ticketsInts) ; i++ {
    if ticketsInts[i][0] == src && ticketsInts[i][1] == dest {
      index = i
      //fmt.Println(ticketsInts)
      //fmt.Println("> found index", src, dest, ">", index)
      break
    }
  }

  return index
}

// broken...
// convert a list of ticket ints back to a path of cities in string format
// this is assuming the list of cities are in sorted order.
func convertIntsToCities(path []int, ticketInts [][]int, sortedCitiesList []string) []string {
  finalCities := make([]string, 0)
  // fmt.Println("> convertIntsToCities")
  // fmt.Println("> ticketIts", ticketInts)
  // fmt.Println("> paths", path)

  curCity := ""
  for i:=0 ; i<len(path) ; i++ {
    //fmt.Println("final cities", finalCities, i, path[i])
    //fmt.Println("final cities ints", ticketInts[path[i]][0], ticketInts[path[i]][1]);
    //fmt.Println(ticketInts[path[i]][1])
    if i==0 {
      curCity = sortedCitiesList[ticketInts[path[i]][0]]
      finalCities = append(finalCities, curCity)
    }
    curCity = sortedCitiesList[ticketInts[path[i]][1]]
    finalCities = append(finalCities, curCity)
  }

  return finalCities
}

func findItinerary(tickets [][]string) []string {
    fmt.Println("> findItin")
    fmt.Println(tickets)
    // 1. Construct a list of verticies.
    sortedCitiesList := orderedVertexList(tickets)
    fmt.Println(sortedCitiesList)

    // 4. Convert pairs to ints instead of string.
    ticketInts := convertCityToInts(tickets, sortedCitiesList)
    fmt.Println(ticketInts)
    
    // 5. Constrct the graph.
    var g Graph
    g.init(len(sortedCitiesList))
    g.insert(ticketInts)
    // fmt.Println("5. >>>", g.vertex[0])
    // fmt.Println("5. >>>", g.vertex[1])
    // fmt.Println("5. >>>", g.vertex[2])
    
    // 5.1 mark path.
    // Use the int order to dictate the path that you used.
    // fmt.Println("> 5.1 >")
    paths := make([]int, len(ticketInts))
    pathTraveled := make([]bool, len(ticketInts))
    currentPos := getCityIndex("JFK", sortedCitiesList)
    
    // all path test.
    g.findAllPaths(currentPos)

    // 6. Go thru the graph. perfer lex ordering.
    //
    //fmt.Println("> 6 start >", currentPos)
    for i:=0 ; i<len(pathTraveled) ; i++ {
        fmt.Println("> 6.1 > i", i, "cPos", currentPos, "> ===========================")
        // return a sorted adj list.
        // fmt.Println(currentPos)
        adj := g.adj(currentPos)
        //fmt.Println("> 6.2 >")
        adj = mergeSortI(adj)
        
        // figure out the first untraveled. edge
        fmt.Println("> 6.2 > adj", adj)
        for j:=0 ; j<len(adj) ; j++ {
            // find the end in ticketsIntIndex.
            curPathIndex := findIndexInTicketInts(currentPos, adj[j], ticketInts)

            fmt.Println("> 6.3 > curPathIndex", curPathIndex, "> src", currentPos, "dest", adj[j], adj)
            if pathTraveled[curPathIndex] == false {
                // mark the current path
                pathTraveled[curPathIndex] = true
                paths[i] = curPathIndex
                fmt.Println("> 6.3 > inserting ", paths, "i", i, "curPathIndex", curPathIndex)
                
                // update the dest city.
                currentPos = adj[j]
                
                // found a city break out.
                break
            }
        }
        fmt.Println("> 6.5 end>")
    }
    
    // 7. Convert results back to array of string.
    fmt.Println("> paths ", paths)
    iten := convertIntsToCities(paths, ticketInts, sortedCitiesList)
    fmt.Println("> 7.")
    fmt.Println(iten)
    //fmt.Println(sortedCitiesList)
    
    return iten
}

//
// ===========================================================================

func main() {
  // tickets := make([][]string, 4)
  // tickets[0] = []string{"MUC", "LHR"}
  // tickets[1] = []string{"JFK", "MUC"}
  // tickets[2] = []string{"SFO", "SJC"}
  // tickets[3] = []string{"LHR", "SFO"}
  
  // [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]

  // tickets := make([][]string, 5)
  // tickets[0] = []string{"JFK", "SFO"}
  // tickets[1] = []string{"JFK", "ATL"}
  // tickets[2] = []string{"SFO", "ATL"}
  // tickets[3] = []string{"ATL", "JFK"}
  // tickets[4] = []string{"ATL", "SFO"}


  // ALGO error... Bascially There is only one valid path...
  // I pick the wrong path by using a bad shortcut.
  // [["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]
  tickets := make([][]string, 3)
  tickets[0] = []string{"JFK", "KUL"}
  tickets[1] = []string{"JFK", "NRT"}
  tickets[2] = []string{"NRT", "JFK"}
  //["JFK","NRT","JFK","KUL"]

  r := findItinerary(tickets)
  fmt.Println(r)
}