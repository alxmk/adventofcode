package main

import "sort"

// This is all massively plagiarised from some dude on Reddit

func bronKerbosch(g map[int][]int) [][]int {
	p := make(intSet)
	for v := range g {
		p.add(v)
	}
	state := &bkState{g: g}
	bronKerbosch1(state, make(intSet), p, make(intSet))
	return state.maxCliques
}

type bkState struct {
	g          map[int][]int
	maxCliques [][]int
}

func bronKerbosch1(state *bkState, r, p, x intSet) {
	if len(p) == 0 && len(x) == 0 {
		if len(state.maxCliques) > 0 && len(state.maxCliques[0]) > len(r) {
			return
		}
		if len(state.maxCliques) > 0 && len(state.maxCliques[0]) < len(r) {
			// Found a longer clique.
			state.maxCliques = nil
		}
		clique := make([]int, 0, len(r))
		for v := range r {
			clique = append(clique, v)
		}
		sort.Ints(clique)
		state.maxCliques = append(state.maxCliques, clique)
		return
	}
	u := -1
	if len(p) > 0 {
		for v := range p {
			u = v
			break
		}
	} else {
		for v := range x {
			u = v
			break
		}
	}
	nu := state.g[u]
	nuSet := make(intSet, len(nu))
	for _, uu := range nu {
		nuSet.add(uu)
	}
	for v := range p {
		if nuSet.contains(v) {
			continue
		}
		ns := state.g[v]
		p1 := make(intSet, len(ns))
		x1 := make(intSet, len(ns))
		for _, n := range ns {
			if p.contains(n) {
				p1.add(n)
			}
			if x.contains(n) {
				x1.add(n)
			}
		}
		r.add(v)
		bronKerbosch1(state, r, p1, x1)
		r.remove(v)
		p.remove(v)
		x.add(v)
	}
}

func (s intSet) contains(v int) bool {
	_, ok := s[v]
	return ok
}

type intSet map[int]struct{}

func (s intSet) add(v int)    { s[v] = struct{}{} }
func (s intSet) remove(v int) { delete(s, v) }
func (s intSet) copy() intSet {
	s1 := make(intSet, len(s))
	for v := range s {
		s1[v] = struct{}{}
	}
	return s1
}
