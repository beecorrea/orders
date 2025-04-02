package order

import (
	"github.com/beecorrea/orders/pkg/fake"
)

type poset struct {
	members []int
	order   Relation
}

type Poset interface {
	Members() []int
	Order() Relation
	IsPartiallyOrdered() bool
	Sort(s Sort) Poset
}

func New(xs []int, order Relation) Poset {
	return poset{
		members: xs,
		order:   order,
	}
}

func Random(order Relation) Poset {
	xs := fake.RandomInts(10)
	return New(xs, order)
}

func (ps poset) Members() []int {
	return ps.members
}

func (ps poset) Order() Relation {
	return ps.order
}

func initMatrix(n int) [][][]bool {
	rels := make([][][]bool, n)
	for range n {
		for j := range n {
			u := make([]bool, n)
			rels[j] = append(rels[j], u)
		}
		u := make([][]bool, n)
		rels = append(rels, u)
	}
	return rels
}

func (ps poset) IsPartiallyOrdered() bool {
	n := len(ps.members)
	rels := initMatrix(n)

	for i := range ps.members {
		rels[i][0][0] = ps.order.Reflexivity(ps.members[i])
		for j := range ps.members {
			rels[i][j][0] = ps.order.Antisymmetry(ps.members[i], ps.members[j])
			for k := range ps.members {
				rels[i][j][k] = ps.order.Transitivity(ps.members[i], ps.members[j], ps.members[k])
				if !(rels[i][0][0] && rels[i][j][0] && rels[i][j][k]) {
					return false
				}
			}
		}
	}

	return true
}

func (ps poset) Sort(s Sort) Poset {
	return s.Run(ps)
}
