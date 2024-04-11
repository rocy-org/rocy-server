package conn

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rocy-org/rocy-server/database/album"
	"github.com/rocy-org/rocy-server/database/artist"
	"github.com/rocy-org/rocy-server/database/music"
	"github.com/rocy-org/rocy-server/database/playlist"
	"github.com/rocy-org/rocy-server/database/resource"
	"github.com/rocy-org/rocy-server/database/role"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnInit() {
	username := os.Getenv("ROCY_SERVER_USERNAME")
	password := os.Getenv("ROCY_SERVER_PASSWORD")
	if username == "" || password == "" {
		panic("failed to connect database")
	}
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/rocy_server_db?charset=utf8mb4&parseTime=True&loc=Local", username, password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&resource.Resource{}, &artist.Artist{}, &album.Album{}, &music.Music{}, &playlist.Playlist{}, &role.Role{}, &user.User{})
}
