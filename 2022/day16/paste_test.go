// https://topaz.github.io/paste/#XQAAAQDZEAAAAAAAAAA4GEiZzRd1JAgz+whYRQxSFI7XvmlfhtGDvX76cUN2pCKaM67D8VPjgDk/uZdVtNnI0yVRZ1WPjBk3sqiEcJ2/FpnNh9TPK1RRdntlhpAZTKlbXtFlMgts7Duw20X9DWTIjE+7tVbLzirxlrJy/qWskhqFnBH6tVULNyCA/FTjhTD4JJVZ8QqD8kFSPI4ydzjxV6mypzogamVmlemFkeREuE3fsz5WujqH5J8l/o57hGWCoB0CkfMM1vyZ4OxOA0G0e8GBuQhwkFbjsLqbi6B7NkyZyUkUOX4IqTjJhuFrhSd/zoDA64gw4NOc6ym/SkAYRofJXG+YlHNuIoBo5CKD9juiI4pk6VziiCAJuSfna//3T0Qiql3e00z4i5qRqg8tSFYJCuvifo/hw/jsaD5kEKj1XUpYKd7Eh6zFyXTevrfMEG7apSSflkMZSWXjHtJ6PjiaXHy2ohpskBVX8SdJ4CTvO/jWEigsAVJuhkv9+zZ5tusWcee5YCxq0HIxVr5PU5uZZ63CnFWRf90H2bP8Opk/ivUgduiTUsyO/8xXnY58B8jUz2DwJRqJ1YAP9JsRrbj4V0IT2IG5xmwCWzkyOt/FYmrNNziNjbqQF+rccsO9V1JlkAoklmTGfBcGIpbbD7xAlKu13VuXh/o8ggYaNFTqASlHURU45MWgEdgwJ88d5JbuCsyS9YCqQ1xv131JRjDKA6B777R6NHHxJlt7igvHXC+ZmX/M2Z9BNGy8y9alnSxIwB6R4TcD0ZsYfbCFmiTlLQJ/05nztm6y7AmLWkb0ILpx5L1X+tFV7KY1O9hQtYf0odsRpZt/eFsabeea03ofOxUZYihMe0fYMtZ/4v+5O1hwSHQG0BqJy1uweRNqueAgw1f3SQkquKnbfKy8RTopv9gs46Vjg5bt4dEPn835i29B+AAdbH9HdJvm5AEz+6DYjRRiHp4JXHibEZqxOaJNXOj6KkpJvqpNGS5v5xqpq219Ex/La88IgZgFXdJJekoHPNACU7NCAmBR8lNHfnfSxN98nXztFizvgWMyiKlxw4OG2kTXtGgszYakxIEfOAgc8NAMpt6LJaI6H1CRRJ67noHRuleJmxpMvDdICorMpHzbq54pux/+hso8rNnS/1NAqManKN8bofy2fDuwaUmgtq5yVqyFGYnaFABZ2XkSrlKFj/49M7b37kxcclqlqgWzGsVLP4gjxfJ2xB9eSytTc5uMbGS4oj8/lDTrW951Hsiz5kxCHusq+Pz9ohrf6sp9jQwyPSR3QIMVXCg+9mIsDiac7O2NGEgXfDFQ8YXgRHkZY9ABZ785EHQbPpEHFueAtBTthmtDJ+VkD+BkeP1Q2atm3rYrdINw3mSV80KnKsmsFAKEo1ZV9s87RfY1q1nnzG7VsD0YuKfsfraWSyoWQmf2zj2Xc+T97le4iUgARZbNWdgvhsdjdBsQQpKXiFIqCZF3sc2JxvIVFjCKdPGLIXHtWqPqx17lcMKmv+Tly4PHFWOCPABkmb1EObYXU7yZPGYsGFm1NawSH45UwC6G+nlZyWr90Eg8LkvB+PrNKtd9r5cBIWboxgvUNrQfd2+pIkQbCPCQlLDvINwCJhPo8CT2Ti2KCaQq8rfCny824S+rjm1HfBxzrLH49/IxIKUyTuYLjaVFgAEoEfhlSX/edWNwbT+CJza5tQeiskWv6rQcCBxGdGRB/7KpD78p8kMTnk3bvdNEwJXAD7plXVMYJgPLDPg6i7PAfb1abENSi5fzMyOzhUC9Gtyr4RdWq1ifhcfSucY3X1fnon/vDxFhkfWSFUH/Md+v13DjkqpiI0EgBv1cS/vY9xlAzMMUCmLpOX04CVZpbL8OikBBXEf/14OQgg==
package main

import (
	_ "embed"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{`Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
		Valve BB has flow rate=13; tunnels lead to valves CC, AA
		Valve CC has flow rate=2; tunnels lead to valves DD, BB
		Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
		Valve EE has flow rate=3; tunnels lead to valves FF, DD
		Valve FF has flow rate=0; tunnels lead to valves EE, GG
		Valve GG has flow rate=0; tunnels lead to valves FF, HH
		Valve HH has flow rate=22; tunnel leads to valve GG
		Valve II has flow rate=0; tunnels lead to valves AA, JJ
		Valve JJ has flow rate=21; tunnel leads to valve II`}, 1651},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
