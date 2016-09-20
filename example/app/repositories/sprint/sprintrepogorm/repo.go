package app_repo_sprint_gorm

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
)

type SprintRepositoryGorm struct {
	sprintgw.SprintGateway

	db *gorm.DB
}

type issueModel struct {
	gorm.Model
	status string
}

type sprintModel struct {
	gorm.Model
	status            string
	expectedClosedAt  time.Time
	effectiveClosedAt time.Time
	issues            []issueModel
}

func New(db *gorm.DB) *SprintRepositoryGorm {
	// create table if needed
	db.AutoMigrate(&issueModel{})
	db.AutoMigrate(&sprintModel{})

	return &SprintRepositoryGorm{
		db: db,
	}
}

func (r *SprintRepositoryGorm) New() (*sprint.Sprint, error) {
	ret := sprint.New()
	entity := sprintModel{
		status: ret.GetStatus(),
	}
	if err := r.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	ret.SetID(int(entity.ID))
	ret.SetCreatedAt(entity.CreatedAt)
	return ret, nil
}

func (r *SprintRepositoryGorm) Add(sprint *sprint.Sprint) error {
	entity := sprintModel{}
	entity.status = sprint.GetStatus()
	entity.effectiveClosedAt = sprint.GetEffectiveClosedAt()
	entity.expectedClosedAt = sprint.GetExpectedClosedAt()
	entity.ID = uint(sprint.GetID())
	entity.CreatedAt = sprint.GetCreatedAt()
	// FIXME: populate issues

	return r.db.Create(&entity).Error
}

func (r SprintRepositoryGorm) Find(id int) (*sprint.Sprint, error) {
	obj := sprintModel{}
	if err := r.db.First(&obj, "id = ?", id).Error; err != nil {
		return nil, err
	}

	fmt.Println(obj)

	ret := sprint.New()
	ret.SetCreatedAt(obj.CreatedAt)
	ret.SetID(int(obj.ID))
	ret.SetStatus(obj.status)
	ret.SetEffectiveClosedAt(obj.effectiveClosedAt)
	ret.SetExpectedClosedAt(obj.expectedClosedAt)

	return ret, nil
}

func (r SprintRepositoryGorm) FindSprintToClose() (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r SprintRepositoryGorm) FindAverageClosedIssues() float64 {
	// Not Implemented
	return float64(-1)
}

func (r *SprintRepositoryGorm) Update(updated *sprint.Sprint) error {
	obj := sprintModel{}
	obj.ID = uint(updated.GetID())
	obj.status = updated.GetStatus()
	obj.CreatedAt = updated.GetCreatedAt()
	obj.expectedClosedAt = updated.GetExpectedClosedAt()
	obj.effectiveClosedAt = updated.GetEffectiveClosedAt()

	// FIXME: populate issues
	return r.db.Save(&obj).Error
}
