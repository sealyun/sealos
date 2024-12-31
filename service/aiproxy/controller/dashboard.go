package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labring/sealos/service/aiproxy/middleware"
	"github.com/labring/sealos/service/aiproxy/model"
)

func getDashboardTime(t string) (time.Time, time.Time, time.Duration) {
	end := time.Now()
	var start time.Time
	var timeSpan time.Duration
	switch t {
	case "month":
		start = end.AddDate(0, 0, -30)
		timeSpan = time.Hour * 24
	case "two_week":
		start = end.AddDate(0, 0, -15)
		timeSpan = time.Hour * 12
	case "week":
		start = end.AddDate(0, 0, -7)
		timeSpan = time.Hour * 6
	case "day":
		fallthrough
	default:
		start = end.AddDate(0, 0, -1)
		timeSpan = time.Hour * 1
	}
	return start, end, timeSpan
}

func fillGaps(data []*model.HourlyChartData, timeSpan time.Duration) []*model.HourlyChartData {
	if len(data) <= 1 {
		return data
	}

	result := make([]*model.HourlyChartData, 0, len(data))
	result = append(result, data[0])

	for i := 1; i < len(data); i++ {
		curr := data[i]
		prev := data[i-1]
		hourDiff := (curr.Timestamp - prev.Timestamp) / int64(timeSpan.Seconds())

		// If gap is 1 hour or less, continue
		if hourDiff <= 1 {
			result = append(result, curr)
			continue
		}

		// If gap is more than 3 hours, only add boundary points
		if hourDiff > 3 {
			// Add point for hour after prev
			result = append(result, &model.HourlyChartData{
				Timestamp: prev.Timestamp + int64(timeSpan.Seconds()),
			})
			// Add point for hour before curr
			result = append(result, &model.HourlyChartData{
				Timestamp: curr.Timestamp - int64(timeSpan.Seconds()),
			})
			result = append(result, curr)
			continue
		}

		// Fill gaps of 2-3 hours with zero points
		for j := prev.Timestamp + int64(timeSpan.Seconds()); j < curr.Timestamp; j += int64(timeSpan.Seconds()) {
			result = append(result, &model.HourlyChartData{
				Timestamp: j,
			})
		}
		result = append(result, curr)
	}

	return result
}

func getTimeSpanWithDefault(c *gin.Context, defaultTimeSpan time.Duration) time.Duration {
	spanStr := c.Query("span")
	if spanStr == "" {
		return defaultTimeSpan
	}
	span, err := strconv.Atoi(spanStr)
	if err != nil {
		return defaultTimeSpan
	}
	if span < 1 || span > 48 {
		return defaultTimeSpan
	}
	return time.Duration(span) * time.Hour
}

func GetDashboard(c *gin.Context) {
	start, end, timeSpan := getDashboardTime(c.Query("type"))
	modelName := c.Query("model")
	timeSpan = getTimeSpanWithDefault(c, timeSpan)

	dashboards, err := model.GetDashboardData(start, end, modelName, timeSpan)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	dashboards.ChartData = fillGaps(dashboards.ChartData, timeSpan)
	middleware.SuccessResponse(c, dashboards)
}

func GetGroupDashboard(c *gin.Context) {
	group := c.Param("group")
	if group == "" {
		middleware.ErrorResponse(c, http.StatusOK, "invalid parameter")
		return
	}

	start, end, timeSpan := getDashboardTime(c.Query("type"))
	tokenName := c.Query("token_name")
	modelName := c.Query("model")
	timeSpan = getTimeSpanWithDefault(c, timeSpan)

	dashboards, err := model.GetGroupDashboardData(group, start, end, tokenName, modelName, timeSpan)
	if err != nil {
		middleware.ErrorResponse(c, http.StatusOK, "failed to get statistics")
		return
	}

	dashboards.ChartData = fillGaps(dashboards.ChartData, timeSpan)
	middleware.SuccessResponse(c, dashboards)
}
