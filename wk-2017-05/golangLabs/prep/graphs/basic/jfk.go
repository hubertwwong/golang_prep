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

func findItinerary(tickets [][]string) []string {
    // 1. Construct a list of verticies.
    sortedCitiesList := orderedVertexList(tickets)
    
    // 4. Convert pairs to ints instead of string.
    ticketInts := convertCityToInts(tickets, sortedCitiesList)
    //fmt.Println(ticketInts)
    
    // 5. Constrct the graph.
    var g Graph
    g.init(len(sortedCitiesList))
    g.insert(ticketInts)
    //fmt.Println("5. >>>", g.vertex[0])
    
    // 5.1 mark path.
    // Use the int order to dictate the path that you used.
    // fmt.Println("> 5.1 >")
    paths := make([]int, len(ticketInts))
    pathTraveled := make([]bool, len(ticketInts))
    currentPos := getCityIndex("JFK", sortedCitiesList)
    
    // 6. Go thru the graph. perfer lex ordering.
    fmt.Println("> 6 >")
    for i:=0 ; i<len(pathTraveled) ; i++ {
        fmt.Println("> 6.1 >", i, currentPos)
        // return a sorted adj list.
        // fmt.Println(currentPos)
        adj := g.adj(currentPos)
        //fmt.Println("> 6.2 >")
        adj = mergeSortI(adj)
        
        // figure out the first untraveled. edge
        fmt.Println("> 6.3 >")
        for j:=0 ; j<len(adj) ; j++ {
            if pathTraveled[adj[j]] == false {
                // mark the current path
                pathTraveled[adj[j]] = true
                paths[adj[j]] = i
                
                // update the dest city.
                currentPos = adj[j]
            }
        }
        fmt.Println("> 6.5 end>")
    }
    
    // 7. Convert results back to array of string.
    fmt.Println("> 7.")
    fmt.Println(paths)
    fmt.Println(sortedCitiesList)
    
    return sortedCitiesList
}

func main() {
	
}