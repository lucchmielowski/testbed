package runtime

// Running is used to test if a container is running or not
func (c *Container) Running() bool {
	return c.Task.Pid > 0
}
