package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"regexp"
)

func validateSlice(slice []string) bool {
	if len(slice) != 0 {
		for _, v := range slice {
			if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", v); !matched || err != nil {
				return false
			}
		}
	}
	return true
}

func contains(value string, slice []string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func getActiveSlugs(conn *pgxpool.Pool, c *gin.Context, userId string) ([]string, error) {
	var activeSlugs []string
	sql := "SELECT slug_name FROM Users_With_Slugs WHERE is_valid=True AND user_id=$1"
	rows, err := conn.Query(context.Background(), sql, userId)
	if err != nil && err.Error() != "no rows in result set" {
		return activeSlugs, err
	}
	for {
		if !rows.Next() {
			break
		}
		var slug string
		err := rows.Scan(&slug)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"rows_error": err.Error()})
		}
		activeSlugs = append(activeSlugs, slug)
	}
	defer rows.Close()
	return activeSlugs, nil
}

func (h *Handlers) UpdateUserSlugs(c *gin.Context) {
	var req UpdateUserSlugsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if matched, err := regexp.MatchString("^[\\d]+$", req.UserId); !matched || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id validation error"})
		return
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !validateSlice(req.InsertSlugs) || !validateSlice(req.DeleteSlugs) || (len(req.InsertSlugs)+len(req.DeleteSlugs) == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slugs validation error"})
		return
	}
	//db := postgres.NewPostgresConnection(h.url)
	//defer db.Close(context.Background())

	activeSlugs, err := getActiveSlugs(h.db, c, req.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(activeSlugs)

	for _, v := range req.InsertSlugs {
		if !contains(v, activeSlugs) {
			sql := "INSERT INTO Users_With_Slugs (user_id, slug_name) VALUES ($1, $2)"
			_, err := h.db.Query(context.Background(), sql, req.UserId, v)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusOK, gin.H{"info": v + " not inserted " + req.UserId})

			}
		}
	}

	activeSlugs, err = getActiveSlugs(h.db, c, req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println(activeSlugs)

	for _, v := range req.DeleteSlugs {
		if contains(v, activeSlugs) {
			sql := "UPDATE Users_With_Slugs SET is_valid=False WHERE is_valid=True AND user_id=$1 AND slug_name=$2"
			_, err := h.db.Query(context.Background(), sql, req.UserId, v)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusOK, gin.H{"info": v + " not updated for " + req.UserId, "delete": err.Error()})
			}
		}
	}

	//err, _ := db.QueryRow(context.Background(), "IN slug_name FROM Users_With_Slugs WHERE is_valid=True AND user_id=$1", req.UserId).Scan(slugs)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "select error"})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
