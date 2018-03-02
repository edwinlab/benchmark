package benchmark

type Controller struct {
	Service Service
}

func (c *Controller) Get() {
	c.Service.UpdateRandom()
}
