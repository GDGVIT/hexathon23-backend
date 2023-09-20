package schemas

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

// Team Serializer for displaying team data
func TeamSerializer(team models.Team) map[string]interface{} {
	return map[string]interface{}{
		"id":          team.ID,
		"name":        team.Name,
		"logo":        team.Logo,
		"members":     team.Members,
		"role":        team.Role,
		"amount":      team.Amount,
		"items":       ItemListSerializer(team.ItemsPurchased),
		"items_count": len(team.ItemsPurchased),
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
		result = append(result, TeamBlockSerializer(team))
	}

	return result
}

// TeamLoginSerializer for displaying team data after login
func TeamLoginSerializer(team models.Team, token string) map[string]interface{} {
	return map[string]interface{}{
		"id":          team.ID,
		"name":        team.Name,
		"logo":        team.Logo,
		"members":     team.Members,
		"role":        team.Role,
		"amount":      team.Amount,
		"items":       ItemListSerializer(team.ItemsPurchased),
		"items_count": len(team.ItemsPurchased),
		"token":       token,
	}
}
