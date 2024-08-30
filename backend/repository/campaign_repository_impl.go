package repository

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &CampaignRepositoryImpl{
		db: db,
	}
}

func (repository CampaignRepositoryImpl) FindAll() ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := repository.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	return helper.ResultOrError(campaigns, err)
}

func (repository CampaignRepositoryImpl) FindByUserID(userID int) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := repository.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns, "user_id = ?", userID).Error
	return helper.ResultOrError(campaigns, err)

}

func (repository CampaignRepositoryImpl) FindByID(ID int) (domain.Campaign, error) {
	var campaign domain.Campaign
	err := repository.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	return helper.ResultOrError(campaign, err)
}

func (repository CampaignRepositoryImpl) Save(campaign domain.Campaign) (domain.Campaign, error) {
	err := repository.db.Create(&campaign).Error
	return helper.ResultOrError(campaign, err)
}

func (repository CampaignRepositoryImpl) Update(campaign domain.Campaign) (domain.Campaign, error) {
	err := repository.db.Save(&campaign).Error
	return helper.ResultOrError(campaign, err)
}

func (repository CampaignRepositoryImpl) CreateImage(image domain.CampaignImage) (domain.CampaignImage, error) {
	err := repository.db.Create(&image).Error
	return helper.ResultOrError(image, err)
}

func (repository CampaignRepositoryImpl) MarkAllImageNonPrimary(ID int) (bool, error) {
	err := repository.db.Model(&domain.CampaignImage{}).Where("id = ?", ID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
