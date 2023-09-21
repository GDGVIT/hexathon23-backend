package schemas

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

// Participant Serializer for displaying participant data
func ParticipantSerializer(participant models.Participant) map[string]interface{} {
	return map[string]interface{}{
		"id":        participant.ID,
		"name":      participant.Name,
		"reg_no":    participant.RegNo,
		"email":     participant.Email,
		"team_id":   participant.TeamID,
		"team_name": participant.Team.Name,
	}
}

// ParticipantBlockSerializer for public participant data
func ParticipantBlockSerializer(participant models.Participant) map[string]interface{} {
	return map[string]interface{}{
		"id":     participant.ID,
		"name":   participant.Name,
		"reg_no": participant.RegNo,
	}
}

// ParticipantListSerializer for displaying list of participants
func ParticipantListSerializer(participants []models.Participant) []map[string]interface{} {
	var result []map[string]interface{}

	for _, participant := range participants {
		result = append(result, ParticipantBlockSerializer(participant))
	}

	return result
}
