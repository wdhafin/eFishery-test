package init

// StartAppInit is
func StartAppInit() {
	setupLogrus()
	setupMainConfig()

	setupAuthHelper()
}
