package getsprint

import (
	"testing"

	"github.com/moul/cleanarch/example/app/repos/sprints/mem"
	"github.com/moul/cleanarch/example/bizrules/entities"
	"github.com/moul/cleanarch/example/bizrules/usecases"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/dto"
	. "github.com/smartystreets/goconvey/convey"
)

func dummyUseCase() UseCase {
	// prepare sprint repo
	repo := sprintsmem.New()
	repo.Add(usecases.SprintStub1)
	repo.Add(usecases.SprintStub2)

	resp := getsprintdto.ResponseAssembler{}

	// prepare usecase
	uc := New(repo, resp)
	return uc
}

func TestUseCase(t *testing.T) {
	Convey("Testing UseCase", t, func() {
		// prepare
		uc := dummyUseCase()
		req := getsprintdto.RequestBuilder{}.Create().WithSprintID(42).Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.GetID(), ShouldEqual, usecases.SprintStub1.GetID())
		So(resp.GetStatus(), ShouldEqual, usecases.SprintStub1.GetStatus())
		So(resp.GetCreatedAt(), ShouldNotEqual, usecases.SprintStub1.GetCreatedAt())
		So(resp.GetCreatedAt(), ShouldResemble, usecases.SprintStub1.GetCreatedAt())
		So(resp.GetEffectiveClosedAt(), ShouldNotEqual, usecases.SprintStub1.GetEffectiveClosedAt())
		So(resp.GetEffectiveClosedAt(), ShouldResemble, usecases.SprintStub1.GetEffectiveClosedAt())
		So(resp.GetExpectedClosedAt(), ShouldNotEqual, usecases.SprintStub1.GetExpectedClosedAt())
		So(resp.GetExpectedClosedAt(), ShouldResemble, usecases.SprintStub1.GetExpectedClosedAt())
	})

	Convey("Testing NotFound", t, func() {
		// prepare
		uc := dummyUseCase()
		req := getsprintdto.RequestBuilder{}.Create().WithSprintID(123456789).Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldNotBeNil)
		So(resp, ShouldBeNil)
		So(err, ShouldResemble, entities.SprintNotFoundError{})
	})
}
