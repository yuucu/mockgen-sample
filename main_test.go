package main

import (
	"fmt"
	"testing"

	"github.com/yuucu/mockgen-sample/random"
	"go.uber.org/mock/gomock"
)

func TestExec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRandom := random.NewMockIFRandom(ctrl)

	mockRandom.EXPECT().Intn(10).Return(5)

	res := exec(mockRandom, 10)

	fmt.Println(res)
}
