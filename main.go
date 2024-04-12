package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DiscordTime/ggdrive/src/repository"
	"github.com/DiscordTime/ggdrive/src/service"
	"github.com/DiscordTime/ggdrive/src/utils"
	"github.com/viniciusalbuquerque/gli/src/gli"
)

var logger utils.Logger

func initRepo() (repository.DriveRepository, error) {
    logger.LogD("initRepo", "Starting...")

    ctx := context.Background()
    gSvc := service.New(ctx, logger)


    if (gSvc == nil) {
	fmt.Println("Exiting")
	return nil, fmt.Errorf("Could not start service")
    }

    driveRepo := repository.New(gSvc, logger)

    return driveRepo, nil

}

func listFiles(params []string) error {
    driveRepo, err := initRepo()
    logger.LogD("listFiles", "Starting...")

    if err != nil {
	return err
    }

    return driveRepo.ListFiles()
}

func downloadFile(params []string) error {
    driveRepo, err := initRepo()
    logger.LogD("downloadFile", "Starting...")

    if err != nil {
	return err
    }

    if len(params) < 1 {
	return fmt.Errorf("Please inform id of the file to be downloaded")
    }

    return driveRepo.DownloadFile(params[0])
}

func uploadFile(params []string) error {
    driveRepo, err := initRepo()
    logger.LogD("uploadFile", "Starting...", params)

    if err != nil {
	return err
    }

    if len(params) < 1 {
	return fmt.Errorf("Please select file to upload")
    }

    return driveRepo.UploadFile(params[0])
}

func main() {
    logger = utils.DefaultLogger{}

    gli.RegisterCommand("list", "List files", listFiles)
    gli.RegisterCommand("download", "Download a file", downloadFile)
    gli.RegisterCommand("upload", "Upload a file", uploadFile)

    gli.ExecCommand(os.Args)
}
