package dto

type CommunityResponse struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Members     int      `json:"members"`
}

type CommunityMembers struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
	JobPosition  string `json:"job_position"`
}
