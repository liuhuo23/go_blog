package request

import "github.com/gin-gonic/gin"

func GetQueryParams(c * gin.Context) map[string]any {
	query := c.Request.URL.Query()
	var querMap = make(map[string]any, len(query))
	for k := range query {
		querMap[k] = c.Query(k)
	}
	return querMap
}

func GetPostFormParams