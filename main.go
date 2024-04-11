package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DiscordTime/ggdrive/src/utils"
	gdrive_repository "github.com/DiscordTime/ggdrive/src/repository"
	gdrive_service "github.com/DiscordTime/ggdrive/src/service"
)

func main() {

    logger := utils.DefaultLogger{}

    for i := 0; i < len(os.Args); i++ {
        logger.LogD(os.Args[i])
    }


    ctx := context.Background()
    gSvc := gdrive_service.New(ctx, logger)

    if (gSvc == nil) {
	fmt.Println("Exiting")
	return
    }
    driveRepo := gdrive_repository.New(gSvc, logger)

    logger.LogD("Main", "Starting...")
    driveRepo.ListFiles()
    //driveRepo.UploadFile(ctx, "test.txt")
    //fmt.Println("Starting...")
    //err := DownloadFile(ctx, "1C71utWp3sOCx5yj-aLsXF3RR0dCaJ7ne")
    //if err != nil {
    //    log.Fatalf("Error while downloading file. %v", err)
    //}
}
