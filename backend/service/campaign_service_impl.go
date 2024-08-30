package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"errors"
	"github.com/gosimple/slug"
	"strconv"
)

type CampaignServiceImpl struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(campaignRepository repository.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{
		campaignRepository: campaignRepository,
	}
}

func (service CampaignServiceImpl) FindAll(userID int) ([]domain.Campaign, error) {
	if userID != 0 {
		campaigns, err := service.campaignRepository.FindByUserID(userID)
		return helper.ResultOrError(campaigns, err)
	}
	campaigns, err := service.campaignRepository.FindAll()
	return helper.ResultOrError(campaigns, err)
}

func (service CampaignServiceImpl) FindByID(request web.CampaignRequestByID) (domain.Campaign, error) {
	campaign, err := service.campaignRepository.FindByID(request.ID)
	return helper.ResultOrError(campaign, err)
}

func (service CampaignServiceImpl) Create(request web.CampaignRequestCreate) (domain.Campaign, error) {
	slug.Make(request.Name + strconv.Itoa(request.User.ID))
	campaign := domain.Campaign{
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		Description:      request.Description,
		GoalAmount:       request.GoalAmount,
		Perks:            request.Perks,
		UserID:           request.User.ID,
		Slug:             slug.Make(request.Name + " " + strconv.Itoa(request.User.ID)),
	}
	result, err := service.campaignRepository.Save(campaign)
	return helper.ResultOrError(result, err)

}

func (service CampaignServiceImpl) Update(campaignID web.CampaignRequestByID, request web.CampaignRequestCreate) (domain.Campaign, error) {
	campaign, err := service.campaignRepository.FindByID(campaignID.ID)
	if err != nil {
		return helper.ResultOrError(campaign, err)
	}

	if campaign.User.ID != request.User.ID {
		return helper.ResultOrError(campaign, errors.New("user is not the owner of the campaign"))
	}
	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.Perks = request.Perks
	campaign.GoalAmount = request.GoalAmount
	result, err := service.campaignRepository.Update(campaign)
	return helper.ResultOrError(result, err)
}

func (service CampaignServiceImpl) CreateCampaignImage(request web.CampaignImageCreate, fileLocation string) (domain.CampaignImage, error) {
	campaign, err := service.campaignRepository.FindByID(request.CampaignID)
	if err != nil {
		return domain.CampaignImage{}, err
	}

	if campaign.UserID != request.User.ID {
		return domain.CampaignImage{}, errors.New("user is not the owner of the campaign")
	}

	if request.IsPrimary {
		_, err := service.campaignRepository.MarkAllImageNonPrimary(request.CampaignID)
		if err != nil {
			return helper.ResultOrError(domain.CampaignImage{}, err)
		}
	}
	campaignImage := domain.CampaignImage{
		CampaignID: request.CampaignID,
		IsPrimary:  request.IsPrimary,
		FileName:   fileLocation,
	}
	image, err := service.campaignRepository.CreateImage(campaignImage)
	return helper.ResultOrError(image, err)
}
