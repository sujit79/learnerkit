package learnerkit

import (
	. "gorgonia.org/gorgonia"
)

type LearnerKit struct {
	node    LearnerKitNode
	graph   LearnerKitGraph
	machine LearnerKitMachine
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

func Initialize_LearnerKit(graph *ExprGraph, machine VM, node_v **Node, size int) LearnerKit {
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

	learnerkit := LearnerKit{
		node:    learnerkitnode,
		graph:   learnerkitgraph,
		machine: learnerkitmachine,
	}

	return learnerkit
}
