package counter

type counter struct {
	value int
}

func (c *counter) Incr() {
	c.value += 1
}

func (c *counter) Value() int {
	return c.value
}

func New() counter {
	return counter{0}
}
