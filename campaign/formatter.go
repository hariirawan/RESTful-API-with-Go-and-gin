package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ImageURL         string `json:"image_url"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

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
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	ImageURL         string                   `json:"image_url"`
	Description      string                   `json:"description"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	CampaignImages   []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary int    `json:"avatar_url"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	detail := CampaignDetailFormatter{}
	detail.ID = campaign.ID
	detail.UserID = campaign.UserID
	detail.Name = campaign.Name
	detail.ShortDescription = campaign.ShortDescription
	detail.ImageURL = ""
	detail.GoalAmount = campaign.GoalAmount
	detail.CurrentAmount = campaign.CurrentAmount
	detail.Description = campaign.Description

	if len(campaign.CampaignImages) > 0 {
		detail.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ", ") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	detail.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.AvatarUrl = user.AvatarFileName

	detail.User = campaignUserFormatter

	campaignImagesFormatter := []CampaignImageFormatter{}

	for _, images := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{
			ImageURL:  images.FileName,
			IsPrimary: images.IsPrimary,
		}
		campaignImagesFormatter = append(campaignImagesFormatter, campaignImageFormatter)
	}

	detail.CampaignImages = campaignImagesFormatter

	return detail
}
