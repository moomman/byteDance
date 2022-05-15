package cmd

import (
	"awesomeProject/day6/test1"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Name string

func init() {
	modelToDataCmd.Flags().StringVarP(&Name, "DbName", "d", "", "转移的指定位置")
}

var modelToDataCmd = &cobra.Command{
	Use:   "move",
	Short: "转移表",
	Long:  "转移表",
	Run: func(cmd *cobra.Command, args []string) {
		err := moveStruct(Name)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func moveStruct(DbName string) error {
	dsn := fmt.Sprintf("root:123456@tcp(127.0.0.1:3307)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&test1.User{})
	if err != nil {
		return err
	}
	return nil
}
