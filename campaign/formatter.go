package campaign

import "strings"

type CampaignFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL string `json:"image_url"`
	GoalAmount int`json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	Slug string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID: campaign.ID,
		UserID: campaign.UserID,
		Name: campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Slug: campaign.Slug,
		ImageURL: "",
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
	ImageURL string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_ammount"`
	UserID int `json:"user_id"`
	Slug string `json:"slug"`
	Perks []string `json:"perks"`    
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{
		ID: campaign.ID,
		Name: campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description: campaign.Description,
		ImageURL: "",
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Slug: campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	return campaignDetailFormatter
}