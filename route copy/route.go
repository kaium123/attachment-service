package route

import (
	"newsfeed/db"
	"newsfeed/modules/attachment/controller"
	"newsfeed/modules/attachment/file_uploader"
	"newsfeed/modules/attachment/repository"
	"newsfeed/modules/attachment/service"

	"github.com/gin-gonic/gin"
)

func AttachmentSetup(api *gin.RouterGroup) {
	db := db.NewEntDb()
	FileInterface := file_uploader.NewFileUploaderFactory()
	repo := repository.NewAttachmentRepository(db)
	svc := service.NewAttachmentService(repo, FileInterface)
	ctlr := controller.NewAttachmentController(svc)

	file := api.Group("file")

	file.GET("", ctlr.GetSingleFile)
	file.DELETE("", ctlr.DeleteSingleFile)
	file.POST("/upload/single", ctlr.UploadSingleFile)
	file.GET("/single/:path", ctlr.GetSingleAttachmentFile)

}
