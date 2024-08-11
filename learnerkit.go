package learnerkit

import (
	. "gorgonia.org/gorgonia"
)

type LearnerKit struct {
	node    LearnerKitNode
	graph   LearnerKitGraph
	machine LearnerKitMachine
	op      LearnerKitOps
}

type LearnerKitNode struct {
	size int
	node **Node
}
type LearnerKitGraph struct {
	g *ExprGraph
}

type LearnerKitMachine struct {
	machine VM
}
type LearnerKitOps struct {
	ops Op
}

func ApplyOps_LearnerKit(op Op, children ...*Node) {
	ApplyOp(op, children...)
}

func ApplyOps_With_Name_LearnerKit(op Op, name string, children ...*Node) {
	ApplyOpWithName(op, name, children...)
}

func Initialize_LearnerKit(graph *ExprGraph, machine VM, node_v **Node, size int, op Op) LearnerKit {
	learnerkitnode := LearnerKitNode{
		size: size,
		node: node_v,
	}
	learnerkitgraph := LearnerKitGraph{
		g: graph,
	}
	learnerkitmachine := LearnerKitMachine{
		machine: machine,
	}
	learnerkitops := LearnerKitOps{
		ops: op,
	}
	learnerkit := LearnerKit{
		node:    learnerkitnode,
		graph:   learnerkitgraph,
		machine: learnerkitmachine,
		op:      learnerkitops,
	}

	return learnerkit
}
