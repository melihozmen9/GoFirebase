package handlers

func HealtCheckHandler() func(c *gin.Context) {
	
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok":true})
	}
}