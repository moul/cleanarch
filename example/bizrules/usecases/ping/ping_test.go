package ping

import (
	"testing"

	"github.com/moul/cleanarch/example/bizrules/usecases/ping/dto"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUseCase(t *testing.T) {
	Convey("Testing UseCase", t, FailureContinues, func() {
		// prepare
		uc := New(pingdto.ResponseAssembler{})
		req := pingdto.RequestBuilder{}.Create().Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.GetPong(), ShouldEqual, "pong")
	})
}
