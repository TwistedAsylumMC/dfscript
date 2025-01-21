package dfscript

type cancelable interface {
	Cancel()
	Canceled() bool
}

type cancelableEvent struct {
	cancel bool
}

func (c *cancelableEvent) Cancel() {
	c.cancel = true
}

func (c *cancelableEvent) Canceled() bool {
	return c.cancel
}
