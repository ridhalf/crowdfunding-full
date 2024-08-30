package web

import (
	"crowdfunding/model/domain"
	"strings"
)

type CampaignDetailResponse struct {
	ID               int                     `json:"id"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	Description      string                  `json:"description"`
	ImageURL         string                  `json:"image_url"`
	GoalAmount       int                     `json:"goal_amount"`
	CurrentAmount    int                     `json:"current_amount"`
	BackerCount      int                     `json:"backer_count"`
	UserID           int                     `json:"user_id"`
	Slug             string                  `json:"slug"`
	Perks            []string                `json:"perks"`
	User             CampaignUserResponse    `json:"user"`
	Images           []CampaignImageResponse `json:"images"`
}
type CampaignUserResponse struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}
type CampaignImageResponse struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func ToCampaignDetailResponse(campaign domain.Campaign) CampaignDetailResponse {
	path := ""
	var perks []string
	if len(campaign.CampaignImages) > 0 {
		path = campaign.CampaignImages[0].FileName
	}
	if campaign.Perks != "" {
		perks = strings.Split(campaign.Perks, ",")
	}
	campaignUserResponse := CampaignUserResponse{
		Name:     campaign.User.Name,
		ImageURL: campaign.User.AvatarFileName,
	}
	var images []CampaignImageResponse
	for _, image := range campaign.CampaignImages {
		isPrimary := false
		if image.IsPrimary {
			isPrimary = true
		}
		campaignImageResponse := CampaignImageResponse{
			ImageURL:  image.FileName,
			IsPrimary: isPrimary,
		}
		images = append(images, campaignImageResponse)
	}
	return CampaignDetailResponse{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		BackerCount:      campaign.BackerCount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		Perks:            perks,
		ImageURL:         path,
		User:             campaignUserResponse,
		Images:           images,
	}
}
