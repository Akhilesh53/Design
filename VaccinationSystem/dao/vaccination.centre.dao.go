package dao

import "pattern/VaccinationSystem/entities"

type VcDao struct {
	vcDao map[string][]entities.IVC
}

func NewVCDao() *VcDao {
	return &VcDao{
		vcDao: make(map[string][]entities.IVC),
	}
}

func (vd *VcDao) AddVC(vc entities.IVC) {
	vd.vcDao[vc.GetDistrict()] = append(vd.vcDao[vc.GetDistrict()], vc)
}

func (vd *VcDao) GetVC(district string) []entities.IVC {
	return vd.vcDao[district]
}

func (vd *VcDao) RemoveVC(district string) {
	delete(vd.vcDao, district)
}

func (vd *VcDao) GetAllVCs() map[string][]entities.IVC {
	return vd.vcDao
}
