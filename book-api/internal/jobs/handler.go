package jobs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var jobIDCounter = 1

func SubmitJob(c *gin.Context) {
	var input struct {
		Payload string `json:"payload"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job data"})
		return
	}

	job := Job{
		ID:      jobIDCounter,
		Payload: input.Payload,
	}
	jobIDCounter++

	JobQueue <- job

	c.JSON(http.StatusAccepted, gin.H{"message": "Job submitted", "job_id": job.ID})
}
