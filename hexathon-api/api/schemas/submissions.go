package schemas

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

// Category Serializer for displaying category data
func SubmissionSerializer(submission models.Submission) map[string]interface{} {
	return map[string]interface{}{
		"id":                submission.ID,
		"team":              TeamBlockSerializer(submission.Team),
		"problem_statement": ProblemStatementSerializer(submission.ProblemStatement),
		"figma_url":         submission.FigmaURL,
		"doc_url":           submission.DocURL,
		"created_at":        submission.CreatedAt,
	}
}

// CategoryListSerializer for displaying list of categories
func SubmissionListSerializer(submissions []models.Submission) []map[string]interface{} {
	var result []map[string]interface{}

	for _, submission := range submissions {
		result = append(result, SubmissionSerializer(submission))
	}

	return result
}
