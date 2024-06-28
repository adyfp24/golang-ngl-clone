package handlers

import "github.com/gin-gonic/gin"

func PrivacyRender(c *gin.Context){
	c.HTML(200, "privacy.html", nil)
}
func SoonRender(c *gin.Context){
	c.HTML(200, "soon.html", nil)
}
func TermsRender(c *gin.Context){
	c.HTML(200, "terms.html", nil)
}