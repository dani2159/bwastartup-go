package campaign

import "strings"

type CampaignFormater struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	Slug             string `json:"slug"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormaterCampaign(campaign Campaign) CampaignFormater {
	campaignFormater := CampaignFormater{}
	campaignFormater.ID = campaign.ID
	campaignFormater.UserID = campaign.UserID
	campaignFormater.Name = campaign.Name
	campaignFormater.ShortDescription = campaign.ShortDescription
	campaignFormater.ImageUrl = ""
	campaignFormater.GoalAmount = campaign.GoalAmount
	campaignFormater.Slug = campaign.Slug
	campaignFormater.CurrentAmount = campaign.CurrentAmount
	if len(campaign.CampaignImages) > 0 {
		campaignFormater.ImageUrl = campaign.CampaignImages[0].FileName
	}
	return campaignFormater
}

func FormaterCampaigns(campaigns []Campaign) []CampaignFormater {
	// if len(campaigns) == 0 {
	// 	return []CampaignFormater{}
	// }

	// var campaignFormaters []CampaignFormater
	// for _, campaign := range campaigns {
	// 	campaignFormaters = append(campaignFormaters, FormaterCampaign(campaign))
	// }
	// return campaignFormaters

	//bisa seperti ini juga
	campaignFormaters := []CampaignFormater{}
	for _, campaign := range campaigns {
		campaignFormaters = append(campaignFormaters, FormaterCampaign(campaign))
	}
	return campaignFormaters
}

type CampaignDetailFormater struct {
	ID               int                     `json:"id"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	ImageUrl         string                  `json:"image_url"`
	GoalAmount       int                     `json:"goal_amount"`
	UserID           int                     `json:"user_id"`
	Slug             string                  `json:"slug"`
	CurrentAmount    int                     `json:"current_amount"`
	Description      string                  `json:"description"`
	Perks            []string                `json:"perks"`
	User             CampaignUserFormater    `json:"user"`
	Images           []CampaignImageFormater `json:"images"`
}

type CampaignUserFormater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormater struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormater {
	CampaignDetailFormater := CampaignDetailFormater{}
	CampaignDetailFormater.ID = campaign.ID
	CampaignDetailFormater.Name = campaign.Name
	CampaignDetailFormater.ShortDescription = campaign.ShortDescription
	CampaignDetailFormater.ImageUrl = ""
	CampaignDetailFormater.GoalAmount = campaign.GoalAmount
	CampaignDetailFormater.UserID = campaign.UserID
	CampaignDetailFormater.Slug = campaign.Slug
	CampaignDetailFormater.CurrentAmount = campaign.CurrentAmount
	CampaignDetailFormater.Description = campaign.Description
	// CampaignDetailFormater.Perks = strings.TrimSpace(strings.Split(campaign.Perks, ","))
	if len(campaign.CampaignImages) > 0 {
		CampaignDetailFormater.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	CampaignDetailFormater.Perks = perks

	campaignUserFormater := CampaignUserFormater{}
	campaignUserFormater.Name = campaign.User.Name
	campaignUserFormater.ImageURL = campaign.User.Avatar
	CampaignDetailFormater.User = campaignUserFormater

	images := []CampaignImageFormater{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormater := CampaignImageFormater{}
		campaignImageFormater.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormater.IsPrimary = isPrimary
		images = append(images, campaignImageFormater)
	}

	CampaignDetailFormater.Perks = perks
	CampaignDetailFormater.Images = images

	return CampaignDetailFormater
}
