package main

import (
	"context"
	"fmt"
	"os"

	gdrive_repository "github.com/DiscordTime/ggdrive/src/repository"
	gdrive_service "github.com/DiscordTime/ggdrive/src/service"
	"github.com/DiscordTime/ggdrive/src/utils"
	"github.com/viniciusalbuquerque/gli/src/gli"
)

var logger utils.Logger

func initRepo() (gdrive_repository.DriveRepository, error) {
    logger.LogD("initRepo", "Starting...")

    ctx := context.Background()
    gSvc := gdrive_service.New(ctx, logger)

    if (gSvc == nil) {
	fmt.Println("Exiting")
	return nil, fmt.Errorf("Could not start service")
    }

    driveRepo := gdrive_repository.New(gSvc, logger)

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

    if len(params) < 3 {
	return fmt.Errorf("Please inform id of the file to be downloaded")
    }

    return driveRepo.DownloadFile(params[3])
}

func uploadFile(params []string) error {
    driveRepo, err := initRepo()
    logger.LogD("uploadFile", "Starting...")

    if err != nil {
	return err
    }

    if len(params) < 3 {
	return fmt.Errorf("Please select file to upload")
    }

    return driveRepo.UploadFile(params[3])
}

func main() {
    logger = utils.DefaultLogger{}

    gli.RegisterCommand("list", "List files", listFiles)
    gli.RegisterCommand("download", "Download a file", downloadFile)
    gli.RegisterCommand("upload", "Upload a file", uploadFile)

    gli.ExecCommand(os.Args)
}
