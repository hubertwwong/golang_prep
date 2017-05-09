package main

import "fmt"

// == GRAPH

type Graph struct {
  vertex []NodeG
}

type NodeG struct {
  name string
  children []NodeG
}

// helper func
// should this be in graph "class"?
func createVertex(children []string, name string) (NodeG) {
  if len(children) == 0 {
    // this is wrong...
    // you can't this return a nil.
    // you can declare a var and return it but that seems weird...
    return NodeG{name, nil}
  }

  // this has to be a make cant declare a slice...
  result := make([]NodeG, len(children))
  for i:= 0 ; i<len(children) ; i++ {
    n := NodeG{children[i], nil}
    result[i] = n
  }

  return NodeG{name, result}
}

// == QUEUE

type NodeQ struct {
  next *NodeQ
  val NodeG
  // this i'm not sure about leaving this as a pointer...
  // I this was wrong... Should not be pointer.
}

type Queue struct {
  root *NodeQ
  tail *NodeQ
}

func (q *Queue) Add(val NodeG) {
  if q.root == nil {
    q.root = &NodeQ{nil, val}
    q.tail = q.root
    //fmt.Println(q.root)
  } else {
    n := &NodeQ{nil, val}
    (*q.tail).next = n
    q.tail = n
  }
}

func (q *Queue) Remove() (NodeG) {
  if q.root == nil {
    return NodeG{"", nil} 
  } else {
    retNode := q.root
    q.root = (*(q.root)).next
    (*retNode).next = nil
    return (*retNode).val
  }
}

func (g *Graph) FindVertex(verStr string) (NodeG) {
  var v NodeG
  for i:=0 ; i<len(g.vertex) ; i++ {
    if g.vertex[i].name == verStr {
      v = g.vertex[i]
    }
  }
  return v
}

// don't use start and end as variable names...
// method name should be upper camel case.
func (g *Graph) HasPath (startS, endS string) (bool) {
  var q Queue

  // add start node
  // the biggie is that you can't create this.
  // you have to take it from the graph...
  //fmt.Println(startV)
  q.Add(g.FindVertex(startS))
  //fmt.Println(">>>",g.FindVertex(startS))

  curV := q.Remove()
  //fmt.Println(curV)
  for ; curV.name != "" ; curV = q.Remove() {
    // has this as val before.
    // this was wrong. should be the graph node.
    // and here too...
    if curV.name == endS {
      return true
    } 
    for i :=0 ; i<len(curV.children) ; i++ {
      v := g.FindVertex(curV.children[i].name)
      //fmt.Println(">", v)
      q.Add(v)
    }
  }

  // nothing found...
  return false
}

func main() {
  v1 := createVertex([]string{"v2"}, "v1")
  v2 := createVertex([]string{"v3"}, "v2")
  v3 := createVertex([]string{}, "v3")
  v4 := createVertex([]string{}, "v4")
  var g Graph
  g.vertex = []NodeG{v1, v2, v3, v4}
  fmt.Println(g.HasPath("v1", "v3"))
}


/*

08:53

one big note...
didn't do the main func on paper...
- not sure if that is good or bad..

08:59
lots of errors...
can't return a nil on a struct...
- can only do that on a pointer.
slices can't use non ints...
- this needs to be a hashmap...
if its a struct, you can't just return nil like a pointer...
- this i'm not sure on the return type...
in adjency list, you have to define all vertexes...

tons of errors...
9:19a

so i created the list with children [] as empty one.
not nil...

09:27
the big thing is that you can't create graph nodes...
thats wrong...
s

lots of errors...
on the queue..
when i created a struct. i didnt set the method ot use the struct poninters...



*/