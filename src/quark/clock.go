package quark

type AtomicClock struct {
	ClockName Name
}

func (c *AtomicClock) Start() *ObjectPosition {
	return c.ClockName.Start()
}

func (c *AtomicClock) End() *ObjectPosition {
	return c.ClockName.End()
}

func (c *AtomicClock) clockExprNode() {}

func (c *AtomicClock) Accept(v Visitor) {
	if v.Visit(c) == nil {
		return
	}

	c.ClockName.Accept(v)
}
