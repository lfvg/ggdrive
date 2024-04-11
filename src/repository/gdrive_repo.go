package gdrive_repository

import (
	"fmt"

	"github.com/DiscordTime/ggdrive/src/utils"
	gdrive_service "github.com/DiscordTime/ggdrive/src/service"
)

type DriveRepository interface {
    ListFiles() error
    DownloadFile(string) error
    UploadFile(string) error
}

type GDriveRepository struct {
    srv gdrive_service.GSvc
    logger utils.Logger
}

func New(gSrv gdrive_service.GSvc, logger utils.Logger) DriveRepository {
    return GDriveRepository{
        srv: gSrv,
        logger: logger,
    }
}

func (drv GDriveRepository) ListFiles() error {
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

func (drv GDriveRepository) DownloadFile(fileId string) error {
    drv.logger.LogD("GDriveRepository", "DownloadFile called")
    return drv.srv.DownloadFile(fileId)
}

func (drv GDriveRepository) UploadFile(filename string) error {
    drv.logger.LogD("GDriveRepository", "UploadFile called")
    return drv.UploadFile(filename)

}

