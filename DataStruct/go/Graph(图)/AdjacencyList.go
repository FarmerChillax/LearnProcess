package main

const MaxVertexNum = 100

type Vertex int
type WeightType int
type DataType string

// 边的结构
type ENode struct {
	V1     Vertex
	Weight WeightType
}

// 邻接点的结构
type AdjVNode struct {
	Adjv   Vertex
	Weight WeightType
	Next   *AdjVNode
}

// 顶点表头结点定义
type Vnode struct {
	FirstEdge *AdjVNode
	Data      DataType
}

var AdjList []Vnode

func main() {

}
