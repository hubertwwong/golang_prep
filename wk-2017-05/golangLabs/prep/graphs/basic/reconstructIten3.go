package main

import "fmt"
import "log"
import "time"

type NodeG struct {
  next *NodeG
  val int
}

type Graph struct {
  vertex []*NodeG
}

func (g *Graph) init(numVertex int) {
  if g.vertex == nil {
    g.vertex = make([]*NodeG, numVertex)
  }
}

func (g *Graph) insert(adjList [][]int) {
  for i:=0 ; i<len(adjList) ; i++ {
    if g.vertex[adjList[i][0]] == nil {
      g.vertex[adjList[i][0]] = &NodeG{nil, adjList[i][1]}
    } else {
      n := g.vertex[adjList[i][0]]
      for ; n.next != nil ; {
        n = n.next
      }
      (*n).next = &NodeG{nil, adjList[i][1]}
    }
  }
}

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

func (g *Graph) findAllPaths(v int, numPaths int, ticketInts [][]int) [][][]int {
  curPath := make([][]int, 0)
  paths := make([][][]int, 0)
  g.findPathDFS(v, curPath, &paths, ticketInts)

  return paths
}

func (g *Graph) findPathDFS(v int, curPath [][]int, paths *[][][]int, ticketInts [][]int) {
  numPaths := len(ticketInts)
  
  // find all adj verticies
  adj := g.adj(v)
  // if there are no adj verticies, you are at the end of a path.
  if len(adj) == 0 && len(curPath) == numPaths {
    (*paths) = append((*paths), curPath)
  } else {
    // flag to mark if you have taken a path
    pathTaken := false

    // dfs for each node. need to check if you haven't taken a path yet.
    for i:=0 ; i<len(adj) ; i++ {
      curEdge := []int{v, adj[i]}
      curPathEdgeCount := countInList(curEdge, curPath) 
      maxPathEdgeCount := countInList(curEdge, ticketInts) 
      if curPathEdgeCount < maxPathEdgeCount {
        dCurPath := append(curPath, curEdge)
        g.findPathDFS(adj[i], dCurPath, paths, ticketInts)
        pathTaken = true
      }
    }

    // if you did not take any path, you are done constructing a possible path.
    // note the number of edges if 1 smaller than vertex for a full path.
    if pathTaken == false && len(curPath) == numPaths {
      (*paths) = append((*paths), curPath)
    }
  }
}

func countInList(v []int, l [][]int) int {
  count := 0
  for i:=0 ; i<len(l) ; i++ {
    if v[0] == l[i][0] && v[1] == l[i][1] {
      count++
    }
  }
  return count
}

// SKIP THIS...
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
    for k, _ := range(cityMap) {
        cityList = append(cityList, k)
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
      break
    }
  }

  return index
}

func convertEdgesToIntPath(edges [][][]int) [][]int {
  paths := make([][]int, 0)

  for i:=0 ; i<len(edges) ; i++ {
    paths = append(paths, convertIntEdgeToPath(edges[i])) 
  }

  return paths
}

// you need this.....
func convertIntEdgeToPath(edges [][]int) []int {
  finalInts := make([]int, 0)

  for i, curInt := 0, 0 ; i<len(edges) ; i++ {
    if i==0 {
      curInt = edges[i][0]
      finalInts = append(finalInts, curInt)
    }
    curInt = edges[i][1]
    finalInts = append(finalInts, curInt)
  }

  return finalInts
}

func convertIntPathToCities(citiesInt []int, cities []string) []string {
  result := make([]string, len(citiesInt))
  for i:=0 ; i<len(citiesInt) ; i++ {
    for j:=0 ; j<len(cities) ; j++ {
      if citiesInt[i] == j {
        result[i] = cities[j]
        break
      }
    }
  }
  return result
}

// min path...
func minPath(paths [][]int) []int {
  if len(paths) == 0 {
    return nil
  }

  // init min path to the zeroth one.
  min := paths[0]
  
  // compare min with exi
  for i:=1 ; i<len(paths) ; i++ {
    // compare each value in the path.
    for j:=0 ; j<len(paths[i]) ; j++ {
      if paths[i][j] < min [j] {
        // current path is smaller than min path. new min.
        min = paths[i]
      } else if paths[i][j] > min [j] {
        // current path is larger than min. break out.
        break
      } else {
        // a tie occured, create a new one.
        continue
      }
    }
  }

  return min
}

func findItinerary(tickets [][]string) []string {
  // 1. Construct a list of verticies.
  sortedCitiesList := orderedVertexList(tickets)
  
  // 2. Convert pairs to ints instead of string.
  ticketInts := convertCityToInts(tickets, sortedCitiesList)
  
  // 3. Constrct the graph.
  var g Graph
  g.init(len(sortedCitiesList))
  g.insert(ticketInts)
  
  // 4. Get int position of JFK.
  currentPos := getCityIndex("JFK", sortedCitiesList)
  
  // 5. Find all paths starting at JFK
  edgePaths := g.findAllPaths(currentPos, len(tickets), ticketInts)

  // 6. Convert int ticket to int paths. Easier to sort int vs. int pairs.
  paths := convertEdgesToIntPath(edgePaths)

  // 7. Get min int path .
  minPath := minPath(paths)
  
  // 8. convert path back to cities.
  iten :=convertIntPathToCities(minPath, sortedCitiesList)

  return iten
}

//
// ===========================================================================

func timeTrack(start time.Time, name string) {
  elapsed := time.Since(start)
  log.Printf("%s took %s", name, elapsed)
}

func main() {
  defer timeTrack(time.Now(), "Main") 
  // tickets := make([][]string, 4)
  // tickets[0] = []string{"JFK", "LHR"}
  // tickets[1] = []string{"JFK", "MUC"}
  // tickets[2] = []string{"SFO", "SJC"}
  // tickets[3] = []string{"LHR", "SFO"}
  
  // tickets := make([][]string, 5)
  // tickets[0] = []string{"JFK", "B"}
  // tickets[1] = []string{"B", "C"}
  // tickets[2] = []string{"C", "B"}
  // tickets[3] = []string{"C", "JFK"}
  // tickets[4] = []string{"JFK", "C"}

	// tickets := make([][]string, 6)
  // tickets[0] = []string{"JFK", "B"}
  // tickets[1] = []string{"B", "C"}
  // tickets[2] = []string{"C", "D"}
  // tickets[3] = []string{"D", "C"}
  // tickets[4] = []string{"B", "D"}
	// tickets[5] = []string{"D", "B"}
	

  // [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]

  // more bugs here.... path is not 5...
	// THIS IS BROKEN AGAIN
  // tickets := make([][]string, 5)
  // tickets[0] = []string{"JFK", "SFO"}
  // tickets[1] = []string{"JFK", "ATL"}
  // tickets[2] = []string{"SFO", "ATL"}
  // tickets[3] = []string{"ATL", "JFK"}
  // tickets[4] = []string{"ATL", "SFO"}

  // ALGO error... Bascially There is only one valid path...
  // I pick the wrong path by using a bad shortcut.
  // [["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]
	//["JFK","NRT","JFK","KUL"]

  // tickets := make([][]string, 3)
  // tickets[0] = []string{"JFK", "KUL"}
  // tickets[1] = []string{"JFK", "NRT"}
  // tickets[2] = []string{"NRT", "JFK"}

  // tickets := make([][]string, 4)
  // tickets[0] = []string{"JFK", "KUL"}
  // tickets[1] = []string{"JFK", "KUL"}
  // tickets[2] = []string{"KUL", "JFK"}
  // tickets[3] = []string{"KUL", "JFK"}


  //[["EZE","AXA"],["TIA","ANU"],["ANU","JFK"],["JFK","ANU"],["ANU","EZE"],["TIA","ANU"],["AXA","TIA"],["TIA","JFK"],["ANU","TIA"],["JFK","TIA"]]
  //["JFK","ANU","EZE","AXA","TIA","ANU","JFK","TIA","ANU","TIA","JFK"]

  // This has repeat flights...
  tickets := make([][]string, 10)
  tickets[0] = []string{"EZE","AXA"}
  tickets[1] = []string{"TIA","ANU"}
  tickets[2] = []string{"ANU","JFK"}
  tickets[3] = []string{"JFK","ANU"}
  tickets[4] = []string{"ANU","EZE"}
  tickets[5] = []string{"TIA","ANU"}
  tickets[6] = []string{"AXA","TIA"}
  tickets[7] = []string{"TIA","JFK"}
  tickets[8] = []string{"ANU","TIA"}
  tickets[9] = []string{"JFK","TIA"}

  //[["AXA","EZE"],["EZE","AUA"],["ADL","JFK"],["ADL","TIA"],["AUA","AXA"],["EZE","TIA"],["EZE","TIA"],["AXA","EZE"],["EZE","ADL"],["ANU","EZE"],["TIA","EZE"],["JFK","ADL"],["AUA","JFK"],["JFK","EZE"],["EZE","ANU"],["ADL","AUA"],["ANU","AXA"],["AXA","ADL"],["AUA","JFK"],["EZE","ADL"],["ANU","TIA"],["AUA","JFK"],["TIA","JFK"],["EZE","AUA"],["AXA","EZE"],["AUA","ANU"],["ADL","AXA"],["EZE","ADL"],["AUA","ANU"],["AXA","EZE"],["TIA","AUA"],["AXA","EZE"],["AUA","SYD"],["ADL","JFK"],["EZE","AUA"],["ADL","ANU"],["AUA","TIA"],["ADL","EZE"],["TIA","JFK"],["AXA","ANU"],["JFK","AXA"],["JFK","ADL"],["ADL","EZE"],["AXA","TIA"],["JFK","AUA"],["ADL","EZE"],["JFK","ADL"],["ADL","AXA"],["TIA","AUA"],["AXA","JFK"],["ADL","AUA"],["TIA","JFK"],["JFK","ADL"],["JFK","ADL"],["ANU","AXA"],["TIA","AXA"],["EZE","JFK"],["EZE","AXA"],["ADL","TIA"],["JFK","AUA"],["TIA","EZE"],["EZE","ADL"],["JFK","ANU"],["TIA","AUA"],["EZE","ADL"],["ADL","JFK"],["ANU","AXA"],["AUA","AXA"],["ANU","EZE"],["ADL","AXA"],["ANU","AXA"],["TIA","ADL"],["JFK","ADL"],["JFK","TIA"],["AUA","ADL"],["AUA","TIA"],["TIA","JFK"],["EZE","JFK"],["AUA","ADL"],["ADL","AUA"],["EZE","ANU"],["ADL","ANU"],["AUA","AXA"],["AXA","TIA"],["AXA","TIA"],["ADL","AXA"],["EZE","AXA"],["AXA","JFK"],["JFK","AUA"],["ANU","ADL"],["AXA","TIA"],["ANU","AUA"],["JFK","EZE"],["AXA","ADL"],["TIA","EZE"],["JFK","AXA"],["AXA","ADL"],["EZE","AUA"],["AXA","ANU"],["ADL","EZE"],["AUA","EZE"]]

  r := findItinerary(tickets)
  fmt.Println(r)
}

/*
I think the issue now is that this is too slow...
One big thing is to cache all of the calucations on the graph.
One big thing is count in list..
the one about 
*/