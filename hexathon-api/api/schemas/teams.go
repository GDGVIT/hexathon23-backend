package schemas

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

// Team Serializer for displaying team data
func TeamSerializer(team models.Team) map[string]interface{} {
	cart, _ := team.GetCart()
	return map[string]interface{}{
		"id":                team.ID,
		"name":              team.Name,
		"logo":              team.Logo,
		"members":           ParticipantListSerializer(team.Members),
		"role":              team.Role,
		"amount":            team.Amount,
		"checked_out":       cart.CheckedOut,
		"items":             ItemListSerializer(team.ItemsPurchased),
		"items_count":       len(team.ItemsPurchased),
		"problem_statement": ProblemStatementSerializer(team.ProblemStatement),
		"submitted":         team.Submitted,
	}
}

// TeamBlockSerializer for public team data
func TeamBlockSerializer(team models.Team) map[string]interface{} {
	return map[string]interface{}{
		"id":   team.ID,
		"name": team.Name,
		"logo": team.Logo,
	}
}

// TeamListSerializer for displaying list of teams
func TeamListSerializer(teams []models.Team) []map[string]interface{} {
	var result []map[string]interface{}

	for _, team := range teams {
		result = append(result, TeamSerializer(team))
	}

	return result
}

// TeamLoginSerializer for displaying team data after login
func TeamLoginSerializer(team models.Team, token string) map[string]interface{} {
	return map[string]interface{}{
		"id":                team.ID,
		"name":              team.Name,
		"logo":              team.Logo,
		"members":           ParticipantListSerializer(team.Members),
		"role":              team.Role,
		"amount":            team.Amount,
		"items":             ItemListSerializer(team.ItemsPurchased),
		"items_count":       len(team.ItemsPurchased),
		"token":             token,
		"problem_statement": ProblemStatementSerializer(team.ProblemStatement),
	}
}

// TeamCredentialsSerializer for displaying team data after generating credentials
func TeamCredentialsSerializer(team models.Team, password string) map[string]interface{} {
	return map[string]interface{}{
		"id":       team.ID,
		"name":     team.Name,
		"logo":     team.Logo,
		"members":  ParticipantListSerializer(team.Members),
		"role":     team.Role,
		"amount":   team.Amount,
		"password": password,
	}
}

// TeamCheckoutSerializer for displaying team data after checkout
func TeamCheckoutSerializer(team models.Team) map[string]interface{} {
	cart, _ := team.GetCart()

	return map[string]interface{}{
		"id":          team.ID,
		"name":        team.Name,
		"logo":        team.Logo,
		"members":     ParticipantListSerializer(team.Members),
		"role":        team.Role,
		"amount":      team.Amount,
		"checked_out": cart.CheckedOut,
		"items":       CartItemsSerializer(team.ItemsPurchased),
		"items_count": len(team.ItemsPurchased),
	}
}
