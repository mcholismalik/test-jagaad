package main

import (
	"os"

	"github.com/test-jagaad/cmd"
	"github.com/test-jagaad/internal/domain"
	"github.com/test-jagaad/internal/usecase"
)

func main() {
	os.Setenv("TMP_DIR", "")
	mockyDom := domain.NewMockyDom()
	userUc := usecase.NewUserUc(mockyDom)
	command := cmd.NewCommandCmd(userUc)
	command.Init()
}
