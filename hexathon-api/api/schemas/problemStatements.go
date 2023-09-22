package schemas

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
)

// ProblemStatementSerializer for displaying problem statement data
func ProblemStatementSerializer(problemStatement models.ProblemStatement) map[string]interface{} {
	return map[string]interface{}{
		"id":          problemStatement.ID,
		"name":        problemStatement.Name,
		"one_liner":   problemStatement.OneLiner,
		"description": problemStatement.Description,
	}
}

// ProblemStatementListSerializer for displaying list of problem statements
func ProblemStatementListSerializer(problemStatements []models.ProblemStatement) []map[string]interface{} {
	var result []map[string]interface{}

	for _, problemStatement := range problemStatements {
		result = append(result, ProblemStatementSerializer(problemStatement))
	}

	return result
}

// ProblemStatementGenerationSerializer for displaying problem statement generation data
func ProblemStatementGenerationSerializer(problemStatement models.ProblemStatement, generations int, statementConfirmed bool) map[string]interface{} {
	return map[string]interface{}{
		"id":                  problemStatement.ID,
		"name":                problemStatement.Name,
		"one_liner":           problemStatement.OneLiner,
		"description":         problemStatement.Description,
		"generations_left":    generations,
		"statement_confirmed": statementConfirmed,
	}
}

// ProblemStatementOneLinerSerializer for displaying problem statement one liner data
func ProblemStatementOneLinerSerializer(problemStatement models.ProblemStatement) map[string]interface{} {
	return map[string]interface{}{
		"id":        problemStatement.ID,
		"name":      problemStatement.Name,
		"one_liner": problemStatement.OneLiner,
	}
}
