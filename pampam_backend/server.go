package main

import (
	"net/http"
	"pampam/backend/tuqa/controller"
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/middlewares"
	"pampam/backend/tuqa/repository"
	"pampam/backend/tuqa/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	devicerepo         repository.DeviceRepo       = repository.NewDeviceRepo()
	deviceservice      service.DeviceService       = service.New(devicerepo)
	devicecontroller   controller.DeviceController = controller.New(deviceservice)
	device             entity.Device
	userrepo           repository.UserRepo           = repository.NewUserRepo()
	userservice        service.UserService           = service.NewUserService(userrepo)
	usercontroller     controller.UserController     = controller.NewUserController(userservice)
	pointerrepo        repository.PointerRepo        = repository.NewPointerRepo()
	pointerservice     service.PointerService        = service.NewPointerService(pointerrepo)
	pointercontroller  controller.PointerController  = controller.NewPointerController(pointerservice)
	volumerepo         repository.VolumeRepo         = repository.NewVolumeRepo()
	volumeservice      service.VolumeService         = service.NewVolumeService(volumerepo)
	volumecontroller   controller.VolumeController   = controller.NewVolumeController(volumeservice)
	hardwareservice    service.HardwareService       = service.NewHardwareService(devicerepo, volumerepo)
	hardwarecontroller controller.HardwareController = controller.NewHardwareController(hardwareservice)
	alamatrepo         repository.AlamatRepo         = repository.NewAlamatRepo()
	alamatservice      service.AlamatService         = service.NewAlamatService(alamatrepo)
	alamatcontroller   controller.AlamatController   = controller.NewAlamatController(alamatservice)
)

func sipa(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"messages": "sipa cantik",
	})
}

func basedevice(ctx *gin.Context) {
	ctx.JSON(200, devicecontroller.FindAll())
}

func devicenew(ctx *gin.Context) {

	err := devicecontroller.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}

func deviceupdate(ctx *gin.Context) {
	err := devicecontroller.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}

func devicedelete(ctx *gin.Context) {
	err := devicecontroller.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}
func alamatshow(ctx *gin.Context) {
	ctx.JSON(200, alamatcontroller.FindAlamat(ctx))
}
func alamatupdate(ctx *gin.Context) {
	alamat := alamatcontroller.Update(ctx)
	if alamat != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wingding"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Hamdabi Jaya Jaya Jaya"})
	}
}
func alamatsave(ctx *gin.Context) {
	alamat := alamatcontroller.Save(ctx)
	if alamat != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wingding"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Hamdabi Jaya Jaya Jaya"})
	}
}
func alamatdelete(ctx *gin.Context) {
	alamat := alamatcontroller.Delete(ctx)
	if alamat == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wingding"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Hamdabi Jaya Jaya Jaya"})
	}
}

func baseuser(ctx *gin.Context) {
	ctx.JSON(200, usercontroller.FindAll())
}

func usernew(ctx *gin.Context) {

	err := usercontroller.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}

func userupdate(ctx *gin.Context) {
	err := usercontroller.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}

func userdelete(ctx *gin.Context) {
	err := usercontroller.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}
func basepointer(ctx *gin.Context) {
	ctx.JSON(200, pointercontroller.FindMerge(ctx))
}
func login(ctx *gin.Context) {
	var userDetail entity.User
	userDetail = usercontroller.Login(ctx)
	if userDetail.Id == 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": "GAADA ANJING"})
	} else {
		ctx.JSON(200, userDetail)
	}
}

func basevolume(ctx *gin.Context) {
	ctx.JSON(200, volumecontroller.FindAll())
}

func volumenew(ctx *gin.Context) {

	err := volumecontroller.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}
func arraydaily(ctx *gin.Context) {
	lalala := volumecontroller.ArrayDaily(ctx)
	if lalala != nil {
		ctx.JSON(http.StatusOK, lalala)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "GAADA GAES"})
	}
}
func arrayhourly(ctx *gin.Context) {
	lalala := volumecontroller.ArrayHourly(ctx)
	if lalala != nil {
		ctx.JSON(http.StatusOK, lalala)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "GAADA GAES"})
	}
}
func hardwareupdate(ctx *gin.Context) {
	hardware := hardwarecontroller.Updatebabang(ctx)
	if hardware == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wingding"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Hamdabi Jaya Jaya Jaya"})
	}
}

func volumeupdate(ctx *gin.Context) {
	err := volumecontroller.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Succesfully Posted"})
	}

}
func dataterakhir(ctx *gin.Context) {
	ctx.JSON(200, volumecontroller.DataTerakhir(ctx))
}
func volumedelete(ctx *gin.Context) {
	volume := volumecontroller.Delete(ctx)
	if volume == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wingding"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"messages": "Hamdabi Jaya Jaya Jaya"})
	}

}
func getvalvestatus(ctx *gin.Context) {
	ctx.JSON(200, devicecontroller.Getvalvestatus(ctx))
}
func volumedaily(ctx *gin.Context) {
	lalala := volumecontroller.FindID(ctx)
	if lalala.Volume != 0 {
		ctx.JSON(http.StatusOK, lalala)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada pemakaian air untuk hari ini"})
	}
}
func volumemonthly(ctx *gin.Context) {
	lalala := volumecontroller.FindMonthly(ctx)
	if lalala.Volume != 0 {
		ctx.JSON(http.StatusOK, lalala)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada pemakaian air untuk hari ini"})
	}
}
func getbaterai(ctx *gin.Context) {
	hehe := devicecontroller.GetBaterai(ctx)
	if hehe != 0 {
		ctx.JSON(http.StatusOK, hehe)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Gadapet Baterai Gaes"})
	}
}

func main() {
	server := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Access-Control-Allow-Origin")
	config.AddAllowHeaders("Authorization")
	server.Use(cors.New(config))
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/device", basedevice)
		apiRoutes.POST("/device/new", devicenew)
		apiRoutes.PUT("/device/:merge_id", deviceupdate)
		apiRoutes.DELETE("/device/:id", devicedelete)
		apiRoutes.GET("/device/baterai/:merge_id", getbaterai)

	}
	userRoutes := server.Group("/user")
	{
		userRoutes.GET("/", baseuser)
		userRoutes.POST("/new", usernew)
		userRoutes.PUT("/update/:id", userupdate)
		userRoutes.DELETE("/delete/:id", userupdate)
		userRoutes.POST("/login", login)
		userRoutes.GET("/pointer/:id", basepointer)
	}
	volumeRoutes := server.Group("/volume")
	{
		volumeRoutes.GET("/", basevolume)
		volumeRoutes.POST("/new", volumenew)
		volumeRoutes.PUT("/update/:id", volumeupdate)
		volumeRoutes.DELETE("/delete/:id", volumedelete)
		volumeRoutes.GET("/daily/:merge_id", volumedaily)
		volumeRoutes.GET("/monthly/:merge_id", volumemonthly)
		volumeRoutes.GET("/last/:device_index", dataterakhir)
		volumeRoutes.GET("/day/:merge_id", arraydaily)
		volumeRoutes.GET("/hour/:merge_id", arrayhourly)
	}
	hardwareRoutes := server.Group("/hardware")
	{
		hardwareRoutes.POST("/update/:merge_id/:volume/:baterai", hardwareupdate)
		hardwareRoutes.GET("/valve/:merge_id", getvalvestatus)
	}
	alamatRoutes := server.Group("/alamat")
	{
		alamatRoutes.GET("/:merge_id", alamatshow)
		alamatRoutes.POST("/new", alamatsave)
		alamatRoutes.PUT("/update/:merge_id", alamatupdate)
		alamatRoutes.DELETE("/delete/:merge_id", alamatdelete)
	}

	server.GET("/", devicecontroller.Show)
	server.GET("/sipa", sipa)
	server.Run(":8000")
}
