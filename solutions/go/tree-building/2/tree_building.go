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
	nodes := map[int]*Node{}
	visited := map[Record]bool{}
	maxNum := -1
	sort.Slice(records, func(i, j int) bool {
		if records[i].Parent != records[j].Parent {
			return records[i].Parent < records[j].Parent
		} else {
			return records[i].ID < records[j].ID
		}
	})
	for i, record := range records {
		if visited[record] {
			return nil, errors.New("Duplicated Record")
		}
		visited[record] = true
		maxNum = max(maxNum, record.ID)
		if i == 0 {
			if record.ID != record.Parent || record.ID != 0 {
				return nil, errors.New("root should be 0")
			} else {
				nodes[0] = &Node{ID: 0}
			}
		} else {
			if record.ID == record.Parent {
				return nil, errors.New("root number wrong")
			} else {
				if record.Parent > record.ID {
					return nil, errors.New("Parent is larger than child")
				}
				if _, ok := nodes[record.ID]; !ok {
					nodes[record.ID] = &Node{ID: record.ID}
				}
				if _, ok := nodes[record.Parent]; !ok {
					nodes[record.Parent] = &Node{ID: record.Parent}
				}
				nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
			}
		}
	}
	if maxNum+1 != len(nodes) {
		return nil, errors.New("No continous")
	} else {
		return nodes[0], nil
	}
}
