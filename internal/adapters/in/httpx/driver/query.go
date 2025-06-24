package driver

import (
	"fmt"
	"go-hex-temp/internal/core/domain"
	"go-hex-temp/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func ClaimQuery(c *gin.Context) *domain.Query {
	params := c.Request.URL.Query()

	query := &domain.Query{
		Filter:     make(domain.Filter),
		Pagination: new(domain.QPagination),
	}

	for rawKey, qVals := range params {
		for _, valStr := range qVals {
			for _, val := range utils.SplitCSV(valStr) {

				switch rawKey {
				case "limit":
					var limit uint = 0
					fmt.Sscanf(val, "%d", &limit)
					if limit == 0 {
						limit = 10
					}
					query.Pagination.Limit = limit
				case "offset":
					fmt.Sscanf(val, "%d", &query.Pagination.Offset)
				case "sort":
					direction := domain.Asc
					key := val
					if val[0] == '-' {
						direction = domain.Desc
						key = val[1:]
					}
					query.Pagination.Sorts = append(query.Pagination.Sorts, domain.QSort{
						Key:       key,
						Direction: direction,
					})
				default:

					var field, opRaw string
					if idx := strings.Index(rawKey, "["); idx != -1 && strings.HasSuffix(rawKey, "]") {
						field = rawKey[:idx]
						opRaw = rawKey[idx+1 : len(rawKey)-1]
					} else {
						field = rawKey
						opRaw = "eq" // default operator
					}

					op := domain.Operator(opRaw)

					if _, exists := query.Filter[field]; !exists {
						query.Filter[field] = make(domain.QCondition)
					}

					query.Filter[field][op] = append(query.Filter[field][op], val)
				}
			}
		}
	}

	return query
}
