package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodeMap := map[int]*Node{}
	visited := map[Record]bool{}
	nums := []int{}
	var root *Node
	for _, record := range records {
		if visited[record] {
			return nil, errors.New("Record visited")
		} else {
			visited[record] = true
		}
		if record.Parent > record.ID {
			return nil, errors.New("Parent ID is bigger than ID")
		}
		nums = append(nums, record.ID)
		if _, ok := nodeMap[record.ID]; !ok {
			nodeMap[record.ID] = &Node{ID: record.ID}
		}
		if _, ok := nodeMap[record.Parent]; !ok {
			nodeMap[record.Parent] = &Node{ID: record.Parent}
		}
		if record.Parent != record.ID {
			nodeMap[record.Parent].Children = append(nodeMap[record.Parent].Children, nodeMap[record.ID])
			sort.Slice(nodeMap[record.Parent].Children, func(i, j int) bool {
				return nodeMap[record.Parent].Children[i].ID < nodeMap[record.Parent].Children[j].ID
			})
		} else {
			if root != nil {
				return nil, errors.New("There is already a root")
			} else {
				root = nodeMap[record.ID]
			}
		}
	}
	sort.Ints(nums)
	for i, num := range nums {
		if i != num {
			return nil, errors.New("Not continue")
		}
	}
	if root == nil {
		return nil, errors.New("No root is found")
	} else {
		return root, nil
	}

}
