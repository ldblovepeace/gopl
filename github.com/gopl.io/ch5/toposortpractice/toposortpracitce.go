package main

import (
	"sort"
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string{
	var keys []string
		for key,_ := range m{
			//range map 结果每次不唯一
		keys = append(keys, key)
	}

	var visitAll func(items []string)
	var seen = make(map[string]bool)
	var order []string
	visitAll = func(items []string){
		for _,item := range items{
			if !seen[item]{
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	sort.Strings(keys)//因为range map 结果不唯一所以排序后再进行toposort
	visitAll(keys)
	return order
}	

func main(){
	for i, course := range topoSort(prereqs){
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}