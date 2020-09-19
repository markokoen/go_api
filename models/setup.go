package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

// SetupModels howsit
func SetupModels() *gorm.DB {

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// viperUser := viper.Get("POSTGRES_USER")
	// viperPassword := viper.Get("POSTGRES_PASSWORD")
	// viperDb := viper.Get("POSTGRES_DB")
	// viperHost := viper.Get("POSTGRES_HOST")
	// viperPort := viper.Get("POSTGRES_PORT")
	// postgresConName := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viperHost, viperPort, viperUser, viperDb, viperPassword)
	viperURL := viper.Get("POSTGRES_URL")

	// fmt.Println("conname is\t\t", postgresConName)

	// cmd := exec.Command("createdb", "-p", fmt.Sprintf("%v", viperPort), "-h", fmt.Sprintf("%v", viperHost), "-U", fmt.Sprintf("%v", "superuser"), "-e", fmt.Sprintf("%v", viperDb))
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// if err := cmd.Run(); err != nil {
	// 	log.Printf("Error: %v", err)
	// }
	// log.Printf("Output: %q\n", out.String())
	fmt.Println("DB connection: " + fmt.Sprintf("%v", viperURL))
	db, err := gorm.Open("postgres", fmt.Sprintf("%v", viperURL))

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Post{})

	m := Post{Name: "Post1", Description: "My post description", ImageApp: "my image url", ImageWeb: "web img url"}
	db.Create(&m)
	fmt.Println("Post added to DB")
	return db
}
