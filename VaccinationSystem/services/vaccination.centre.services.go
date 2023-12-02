package services

import (
	"pattern/VaccinationSystem/dao"
	"pattern/VaccinationSystem/entities"
)

type vaccinationCentreService struct {
	vcDao dao.VcDao
}

func NewVCService() *vaccinationCentreService {
	return &vaccinationCentreService{
		vcDao: *dao.NewVCDao(),
	}
}

func (vcs *vaccinationCentreService) AddVC(vc entities.IVC) {
	vcs.vcDao.AddVC(vc)
}

func (vcs *vaccinationCentreService) GetVC(district string) []entities.IVC {
	return vcs.vcDao.GetVC(district)
}

func (vcs *vaccinationCentreService) RemoveVC(district string) {
	vcs.vcDao.RemoveVC(district)
}

func (vcs *vaccinationCentreService) GetAllVCs() map[string][]entities.IVC {
	return vcs.vcDao.GetAllVCs()
}


