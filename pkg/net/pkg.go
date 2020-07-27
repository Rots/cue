// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../gen/gen.go

package net

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("net", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name:   "SplitHostPort",
		Params: []adt.Kind{adt.StringKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret, c.Err = SplitHostPort(s)
			}
		},
	}, {
		Name:   "JoinHostPort",
		Params: []adt.Kind{adt.TopKind, adt.TopKind},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			host, port := c.Value(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = JoinHostPort(host, port)
			}
		},
	}, {
		Name:   "FQDN",
		Params: []adt.Kind{adt.StringKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = FQDN(s)
			}
		},
	}, {
		Name:  "IPv4len",
		Const: "4",
	}, {
		Name:  "IPv6len",
		Const: "16",
	}, {
		Name:   "ParseIP",
		Params: []adt.Kind{adt.StringKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret, c.Err = ParseIP(s)
			}
		},
	}, {
		Name:   "IPv4",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = IPv4(ip)
			}
		},
	}, {
		Name:   "IP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = IP(ip)
			}
		},
	}, {
		Name:   "LoopbackIP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = LoopbackIP(ip)
			}
		},
	}, {
		Name:   "MulticastIP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = MulticastIP(ip)
			}
		},
	}, {
		Name:   "InterfaceLocalMulticastIP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = InterfaceLocalMulticastIP(ip)
			}
		},
	}, {
		Name:   "LinkLocalMulticastIP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = LinkLocalMulticastIP(ip)
			}
		},
	}, {
		Name:   "LinkLocalUnicastIP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = LinkLocalUnicastIP(ip)
			}
		},
	}, {
		Name:   "GlobalUnicastIP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = GlobalUnicastIP(ip)
			}
		},
	}, {
		Name:   "UnspecifiedIP",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret = UnspecifiedIP(ip)
			}
		},
	}, {
		Name:   "ToIP4",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret, c.Err = ToIP4(ip)
			}
		},
	}, {
		Name:   "ToIP16",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.ListKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret, c.Err = ToIP16(ip)
			}
		},
	}, {
		Name:   "IPString",
		Params: []adt.Kind{adt.TopKind},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			ip := c.Value(0)
			if c.Do() {
				c.Ret, c.Err = IPString(ip)
			}
		},
	}},
}
