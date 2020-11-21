package routes

import (
	"net/http"

	"github.com/nocubicles/veloturg/models"
	"github.com/nocubicles/veloturg/utils"
)

type AdData struct {
	ID        uint
	Title     string
	Thumbnail string
}

func getAdsData(ads *[]models.Ad) []AdData {
	adsData := []AdData{}

	for _, ad := range *ads {
		adImageUrls := utils.GetAdImageUrls(ad.ID)
		thumbNail := ""
		if len(adImageUrls) > 0 {
			thumbNail = adImageUrls[0]
		}
		adData := AdData{
			ID:        ad.ID,
			Title:     ad.Title,
			Thumbnail: thumbNail,
		}
		adsData = append(adsData, adData)
	}
	return adsData
}

func RenderHome(w http.ResponseWriter, r *http.Request) {

	ads := &[]models.Ad{}

	db := utils.DbConnection()

	result := db.Limit(12).Select("title", "ID").Find(&ads)
	data := []AdData{}
	if result.RowsAffected > 0 {
		data = getAdsData(ads)
		utils.Render(w, "index.html", data)
	} else {
		utils.Render(w, "index.html", data)

	}
}
