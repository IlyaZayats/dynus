package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func (h *Handlers) ActiveUserSlugs(c *gin.Context) {
	var data ActiveUserSlugsRequest
	if err := c.BindUri(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if matched, err := regexp.MatchString("^[\\d]+$", data.UserId); !matched || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}
	activeSlugs, err := getActiveSlugs(h.db, c, data.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//rows, err := h.db.Query(context.Background(), "SELECT slug_name FROM Users_With_Slugs WHERE is_valid=True AND user_id=$1", data.UserId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "select error"})
	//	return
	//}
	//var activeSlugs []string
	//for {
	//	if !rows.Next() {
	//		break
	//	}
	//	var slug string
	//	err := rows.Scan(&slug)
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{"rows_error": err.Error()})
	//	}
	//	activeSlugs = append(activeSlugs, slug)
	//}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user_id": data.UserId, "slugs": activeSlugs})
}
