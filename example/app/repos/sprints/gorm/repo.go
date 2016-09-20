package sprintsgorm

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/moul/cleanarch/example/bizrules/entities"
	"github.com/moul/cleanarch/example/bizrules/gateways"
)

type Repo struct {
	gateways.Sprints

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

func New(db *gorm.DB) *Repo {
	// create table if needed
	db.AutoMigrate(&issueModel{})
	db.AutoMigrate(&sprintModel{})

	return &Repo{
		db: db,
	}
}

func (r *Repo) New() (*entities.Sprint, error) {
	ret := entities.NewSprint()
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

func (r *Repo) Add(sprint *entities.Sprint) error {
	entity := sprintModel{}
	entity.status = sprint.GetStatus()
	entity.effectiveClosedAt = sprint.GetEffectiveClosedAt()
	entity.expectedClosedAt = sprint.GetExpectedClosedAt()
	entity.ID = uint(sprint.GetID())
	entity.CreatedAt = sprint.GetCreatedAt()
	// FIXME: populate issues

	return r.db.Create(&entity).Error
}

func (r Repo) Find(id int) (*entities.Sprint, error) {
	obj := sprintModel{}
	if err := r.db.First(&obj, "id = ?", id).Error; err != nil {
		return nil, err
	}

	ret := entities.NewSprint()
	ret.SetCreatedAt(obj.CreatedAt)
	ret.SetID(int(obj.ID))
	ret.SetStatus(obj.status)
	ret.SetEffectiveClosedAt(obj.effectiveClosedAt)
	ret.SetExpectedClosedAt(obj.expectedClosedAt)

	return ret, nil
}

func (r Repo) FindSprintToClose() (*entities.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r Repo) FindAverageClosedIssues() float64 {
	// Not Implemented
	return float64(-1)
}

func (r *Repo) Update(updated *entities.Sprint) error {
	obj := sprintModel{}
	obj.ID = uint(updated.GetID())
	obj.status = updated.GetStatus()
	obj.CreatedAt = updated.GetCreatedAt()
	obj.expectedClosedAt = updated.GetExpectedClosedAt()
	obj.effectiveClosedAt = updated.GetEffectiveClosedAt()

	// FIXME: populate issues
	return r.db.Save(&obj).Error
}
