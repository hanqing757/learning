package gopl

import (
	"learning/util"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures", "formal languages", "computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	//"linear algebra":        {"calculus"},  // test hasring
}

//TopSort
/**
获取学习课程的顺序
学习某个课程时需要先学完它的前置课程
课程依赖不能有环
 */
func TopSort() []string {
	seen := make(map[string]bool)
	var order []string
	var visitAll func([]string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(prereqs[item])
				order = append(order, item)
			}
		}
	}

	var course []string
	for k := range prereqs {
		course = append(course, k)
	}
	sort.Strings(course)
	visitAll(course)

	return order
}

// HasRing  判断课程是否有相互依赖
func HasRing() bool {
	seen := make(map[string]bool)
	var order []string
	var visitAll func([]string, []string) bool
	visitAll = func(items []string, l []string) bool {
		tmp := make([]string, len(l))
		copy(tmp, l)
		for _, item := range items {
			if util.InStringSlice(tmp, item) {
				return false
			}
			if !seen[item] {
				seen[item] = true
				if !visitAll(prereqs[item], append(tmp, item)){
					return false
				}
				order = append(order, item)
			}
		}
		return true
	}

	var course []string
	for k := range prereqs {
		course = append(course, k)
	}
	sort.Strings(course)
	return visitAll(course, []string{})
}

