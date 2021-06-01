package controllers

import (
	"fmt"
	"go-trendy-wash-backend/models"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

func getBannerHomeList(c echo.Context) error {
	var (
		bannerModel     models.BannerModel
		bannerModelList []models.BannerModel
	)

	rows, err := conn.Query(`
		SELECT 
			banner_id, 
			CONCAT('` + urlImage + `',banner_url), 
			create_date,
			link_url
		FROM banner 
		WHERE banner_type='HOME'
		AND active = 1`)
	if err != nil {
		fmt.Println("err query : ", err)
	}

	for rows.Next() {
		err = rows.Scan(
			&bannerModel.BannerID,
			&bannerModel.BannerUrl,
			&bannerModel.CreateDate,
			&bannerModel.LinkUrl,
		)
		if err != nil {
			fmt.Println(err)
		} else {
			bannerModelList = append(bannerModelList, bannerModel)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": bannerModelList,
	})
}

func getBannerPromotionList(c echo.Context) error {
	var (
		bannerModel     models.BannerModel
		bannerModelList []models.BannerModel
	)

	rows, err := conn.Query(`
		SELECT 
			banner_id, 
			CONCAT('` + urlImage + `',banner_url), 
			create_date,
			link_url
		FROM banner 
		WHERE banner_type='PROMOTION'
		AND active = 1`)
	if err != nil {
		fmt.Println("err query : ", err)
	}

	for rows.Next() {
		err = rows.Scan(
			&bannerModel.BannerID,
			&bannerModel.BannerUrl,
			&bannerModel.CreateDate,
			&bannerModel.LinkUrl,
		)
		if err != nil {
			fmt.Println(err)
		} else {
			bannerModelList = append(bannerModelList, bannerModel)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"result": bannerModelList,
	})
}

func postBanner(c echo.Context) error {

	imagePath := "image/"

	// Source
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileName, err := GenerateRandomString(16)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  false,
			"message": "create file error : " + err.Error(),
		})
	}
	// Destination
	dst, err := os.Create(imagePath + fileName) // file.Filename
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	rows, err := conn.Query(`
		INSERT INTO banner(
			banner_url,
			banner_type,
			create_date,
			link_url) VALUES(?,?,?,?)`,
		"/api/v1/"+imagePath+fileName,
		c.FormValue("type"),
		time.Now().Add(timeInc),
		c.FormValue("link_url"),
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  false,
			"message": "insert db error : " + err.Error(),
		})
	}
	defer rows.Close()

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
	})
}

func delBanner(c echo.Context) error {

	bannerId := c.QueryParam("banner_id")

	rows, err := conn.Query(`
		UPDATE banner
		SET active = 0
		WHERE banner_id = ` + bannerId)

	if err != nil {
		fmt.Println(timeLog, " - err query : ", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": false,
		})
	}
	rows.Close()

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
	})
}

func BannerDock(_echo *echo.Group) {
	_echo.GET("/banner/home/list", getBannerHomeList)
	_echo.GET("/banner/promotion/list", getBannerPromotionList)
	_echo.DELETE("/banner", delBanner)

	_echo.POST("/banner", postBanner)
}
