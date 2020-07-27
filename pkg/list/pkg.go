// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../gen/gen.go

package list

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("list", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name:   "Drop",
		Params: []adt.Kind{adt.ListKind, adt.IntKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			x, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Drop(x, n)
			}
		},
	}, {
		Name:   "FlattenN",
		Params: []adt.Kind{adt.TopKind, adt.IntKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			xs, depth := c.Value(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = FlattenN(xs, depth)
			}
		},
	}, {
		Name:   "Take",
		Params: []adt.Kind{adt.ListKind, adt.IntKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			x, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Take(x, n)
			}
		},
	}, {
		Name:   "Slice",
		Params: []adt.Kind{adt.ListKind, adt.IntKind, adt.IntKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			x, i, j := c.List(0), c.Int(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = Slice(x, i, j)
			}
		},
	}, {
		Name:   "MinItems",
		Params: []adt.Kind{adt.ListKind, adt.IntKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret = MinItems(a, n)
			}
		},
	}, {
		Name:   "MaxItems",
		Params: []adt.Kind{adt.ListKind, adt.IntKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret = MaxItems(a, n)
			}
		},
	}, {
		Name:   "UniqueItems",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a := c.List(0)
			if c.Do() {
				c.Ret = UniqueItems(a)
			}
		},
	}, {
		Name:   "Contains",
		Params: []adt.Kind{adt.ListKind, adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a, v := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret = Contains(a, v)
			}
		},
	}, {
		Name:   "Avg",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Avg(xs)
			}
		},
	}, {
		Name:   "Max",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Max(xs)
			}
		},
	}, {
		Name:   "Min",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Min(xs)
			}
		},
	}, {
		Name:   "Product",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Product(xs)
			}
		},
	}, {
		Name:   "Range",
		Params: []adt.Kind{adt.NumKind, adt.NumKind, adt.NumKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			start, limit, step := c.Decimal(0), c.Decimal(1), c.Decimal(2)
			if c.Do() {
				c.Ret, c.Err = Range(start, limit, step)
			}
		},
	}, {
		Name:   "Sum",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Sum(xs)
			}
		},
	}, {
		Name:   "Sort",
		Params: []adt.Kind{adt.ListKind, adt.TopKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = Sort(list, cmp)
			}
		},
	}, {
		Name:   "SortStable",
		Params: []adt.Kind{adt.ListKind, adt.TopKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = SortStable(list, cmp)
			}
		},
	}, {
		Name:   "SortStrings",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			a := c.StringList(0)
			if c.Do() {
				c.Ret = SortStrings(a)
			}
		},
	}, {
		Name:   "IsSorted",
		Params: []adt.Kind{adt.ListKind, adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret = IsSorted(list, cmp)
			}
		},
	}, {
		Name:   "IsSortedStrings",
		Params: []adt.Kind{adt.ListKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			a := c.StringList(0)
			if c.Do() {
				c.Ret = IsSortedStrings(a)
			}
		},
	}},
	CUE: `{
	Comparer: {
		T:    _
		x:    T
		y:    T
		less: bool
	}
	Ascending: {
		Comparer
		T:    number | string
		x:    T
		y:    T
		less: true && x < y
	}
	Descending: {
		Comparer
		T:    number | string
		x:    T
		y:    T
		less: x > y
	}
}`,
}
