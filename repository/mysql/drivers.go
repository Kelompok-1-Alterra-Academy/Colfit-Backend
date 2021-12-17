package mysql

import (
	"CalFit/repository/mysql/addresses"
	"CalFit/repository/mysql/booking_details"
	"CalFit/repository/mysql/classes"
	"CalFit/repository/mysql/gyms"
	"CalFit/repository/mysql/membership_types"
	"CalFit/repository/mysql/newsletters"
	"CalFit/repository/mysql/operational_admins"
	"CalFit/repository/mysql/payments"
	"CalFit/repository/mysql/schedules"
	"CalFit/repository/mysql/sessions"
	"CalFit/repository/mysql/super_admins"
	"CalFit/repository/mysql/users"
	"CalFit/repository/mysql/video_contents"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_USERNAME": viper.GetString("database.username"),
		"DB_PASSWORD": viper.GetString("database.password"),
		"DB_HOST":     viper.GetString("database.host"),
		"DB_PORT":     viper.GetString("database.port"),
		"DB_NAME":     viper.GetString("database.name"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_USERNAME"], config["DB_PASSWORD"], config["DB_HOST"], config["DB_PORT"], config["DB_NAME"])

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&super_admins.Super_admin{},
		&operational_admins.Operational_admin{},
		&addresses.Address{},
		&membership_types.Membership_type{},
		&users.User{},
		&newsletters.Newsletter{},
		&gyms.Gym{},
		&payments.Payment{},
		&sessions.Session{},
		&schedules.Schedule{},
		&classes.Class{},
		&video_contents.Video_content{},
		&booking_details.Booking_detail{},
	)
	return DB
}
