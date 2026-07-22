func (c *Pipeline) exec() error {
	for {
		// Reset errors for all commands before attempting execution
		for _, cmd := range c.cmds {
			cmd.SetErr(nil)
		}

		err := c.processPipeline(c.cmds)
		if err == nil || !c.shouldRetry(err) {
			return err
		}
	}
}