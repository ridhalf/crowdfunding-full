package web

import "crowdfunding/model/domain"

type CampaignResponse struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func ToCampaignResponse(campaign domain.Campaign) CampaignResponse {
	path := ""
	if len(campaign.CampaignImages) > 0 {
		path = campaign.CampaignImages[0].FileName
	}
	return CampaignResponse{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageURL:         path,
	}
}
func ToCampaignsResponse(campaigns []domain.Campaign) []CampaignResponse {

	if len(campaigns) == 0 {
		return []CampaignResponse{}
	}
	var campaignResponses []CampaignResponse
	for _, campaign := range campaigns {
		campaignResponse := ToCampaignResponse(campaign)
		campaignResponses = append(campaignResponses, campaignResponse)
	}
	return campaignResponses
}
