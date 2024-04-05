package gdrive_repository

import (
	"context"
	"fmt"

	"github.com/DiscordTime/ggdrive/src/utils"
	gdrive_service "github.com/DiscordTime/ggdrive/src/service"
)

type DriveRepository interface {
    ListFiles(context.Context) error
    DownloadFile(context.Context, string) error
    UploadFile(context.Context, string) error
}

type GDriveRepository struct {
    srv *gdrive_service.GdriveService
    logger utils.Logger
}

func New(gSrv *gdrive_service.GdriveService, logger utils.Logger) DriveRepository {
    return GDriveRepository{
        srv: gSrv,
        logger: logger,
    }
}

func (drv GDriveRepository) ListFiles(ctx context.Context) error {
    drv.logger.LogD("GDriveRepository", "List files called")
    r, err := drv.srv.ListFiles(10)
    if err != nil {
        drv.logger.LogD("GDriveRepository", "Unable to retrieve files")
        return err
    }
    if len(r.Files) == 0 {
        fmt.Println("No files found.")
    } else {
        fmt.Println("Files:")
        for _, i := range r.Files {
            fmt.Printf("%s (%s)\n", i.Name, i.Id)
        }
    }
    return nil
}

func (drv GDriveRepository) DownloadFile(ctx context.Context, fileId string) error {
    drv.logger.LogD("GDriveRepository", "DownloadFile called")
    return drv.srv.DownloadFile(ctx, fileId)
}

func (drv GDriveRepository) UploadFile(ctx context.Context, filename string) error {
    drv.logger.LogD("GDriveRepository", "UploadFile called")
    return drv.UploadFile(ctx, filename)

}

