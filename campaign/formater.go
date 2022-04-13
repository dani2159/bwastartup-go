package campaign

type CampaignFormater struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
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
