package main

import (
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCarList(w http.ResponseWriter, r *http.Request) {
	type Car struct {
		DealerName    string    `json:"dealer_name"`
		DealerProfile string    `json:"dealer_profile"`
		PostedDate    time.Time `json:"posted_date"`
		CarImg        string    `json:"car_img"`
		CarModel      string    `json:"car_model"`
		CarYear       string    `json:"car_year"`
		Price         int       `json:"price"`
		LikeCount     string    `json:"like_count"`
		CommentCount  string    `json:"comment_count"`
		ViewCount     string    `json:"view_count"`
	}

	car := []Car{
		{
			DealerName:    "Ko Htet Aung Khant",
			DealerProfile: "https://s3-alpha-sig.figma.com/img/2714/df7a/1e9d94a74c92e29d5fdf1c828724f623?Expires=1744588800&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=MExi4KyWbv79dnOk820xl7UjfRQsyLmMYLdzkYgI-R2HwqCLGaCMAIaxh8cDrpJOvbhYs9chNUaEqze1FqPFncJcIoAOd557VH8GRMqd9Mz5aGimaYXq9zUQdZlSrfD5hgi6L~ijb16BCkUWuFUrCK9n~HjQtBX3qy~JaR0ECY2dMil~CdCRbww65M4nKCC5Yg9AvyPSsj-oW1O6zRN4ciG1OH-TbaeNoxXnIPoRSV1WQDL3GdJA6ZPry4Y2o0j6B0AxnFP-XYJ0ndnNdtnChdY2c35ezbTSV6MnTNf6LgaKFrhnuYp4tie1xnXDBKWYjhAsTk~jDS~Dtp2EkRWX3A__",
			PostedDate:    time.Now().UTC(),
			CarImg:        "https://media.drive.com.au/obj/tx_q:50,rs:auto:1920:1080:1/caradvice/private/c5274f409066e881b1273f2ebbe1f8cd",
			CarModel:      "SRT Dodge Challenger",
			CarYear:       "2024",
			Price:         999999,
			LikeCount:     "988M",
			CommentCount:  "100k",
			ViewCount:     "931.8M",
		},
		{
			DealerName:    "Ko Naing Linn Phyo",
			DealerProfile: "https://s3-alpha-sig.figma.com/img/2714/df7a/1e9d94a74c92e29d5fdf1c828724f623?Expires=1744588800&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=MExi4KyWbv79dnOk820xl7UjfRQsyLmMYLdzkYgI-R2HwqCLGaCMAIaxh8cDrpJOvbhYs9chNUaEqze1FqPFncJcIoAOd557VH8GRMqd9Mz5aGimaYXq9zUQdZlSrfD5hgi6L~ijb16BCkUWuFUrCK9n~HjQtBX3qy~JaR0ECY2dMil~CdCRbww65M4nKCC5Yg9AvyPSsj-oW1O6zRN4ciG1OH-TbaeNoxXnIPoRSV1WQDL3GdJA6ZPry4Y2o0j6B0AxnFP-XYJ0ndnNdtnChdY2c35ezbTSV6MnTNf6LgaKFrhnuYp4tie1xnXDBKWYjhAsTk~jDS~Dtp2EkRWX3A__",
			PostedDate:    time.Now().UTC(),
			CarImg:        "https://masterpiecer-images.s3.yandex.net/7f677dcc773b11eeaead5696910b1137:upscaled",
			CarModel:      "TOYOTA MARK II",
			CarYear:       "2024",
			Price:         998888,
			LikeCount:     "98M",
			CommentCount:  "10k",
			ViewCount:     "93.8M",
		},
		{
			DealerName:    "Ko Win Thein",
			DealerProfile: "https://s3-alpha-sig.figma.com/img/2714/df7a/1e9d94a74c92e29d5fdf1c828724f623?Expires=1744588800&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=MExi4KyWbv79dnOk820xl7UjfRQsyLmMYLdzkYgI-R2HwqCLGaCMAIaxh8cDrpJOvbhYs9chNUaEqze1FqPFncJcIoAOd557VH8GRMqd9Mz5aGimaYXq9zUQdZlSrfD5hgi6L~ijb16BCkUWuFUrCK9n~HjQtBX3qy~JaR0ECY2dMil~CdCRbww65M4nKCC5Yg9AvyPSsj-oW1O6zRN4ciG1OH-TbaeNoxXnIPoRSV1WQDL3GdJA6ZPry4Y2o0j6B0AxnFP-XYJ0ndnNdtnChdY2c35ezbTSV6MnTNf6LgaKFrhnuYp4tie1xnXDBKWYjhAsTk~jDS~Dtp2EkRWX3A__",
			PostedDate:    time.Now().UTC(),
			CarImg:        "https://www-asia.nissan-cdn.net/content/dam/Nissan/global/vehicles/gt-r/r35/jprhd/2_minor_change/nismo/17TDIjprhd_GTRHelios146.jpg.ximg.l_full_m.smart.jpg",
			CarModel:      "GTR r35",
			CarYear:       "2020",
			Price:         9009,
			LikeCount:     "189",
			CommentCount:  "39",
			ViewCount:     "10.5M",
		},
		{
			DealerName:    "Ko Ame",
			DealerProfile: "https://s3-alpha-sig.figma.com/img/2714/df7a/1e9d94a74c92e29d5fdf1c828724f623?Expires=1744588800&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=MExi4KyWbv79dnOk820xl7UjfRQsyLmMYLdzkYgI-R2HwqCLGaCMAIaxh8cDrpJOvbhYs9chNUaEqze1FqPFncJcIoAOd557VH8GRMqd9Mz5aGimaYXq9zUQdZlSrfD5hgi6L~ijb16BCkUWuFUrCK9n~HjQtBX3qy~JaR0ECY2dMil~CdCRbww65M4nKCC5Yg9AvyPSsj-oW1O6zRN4ciG1OH-TbaeNoxXnIPoRSV1WQDL3GdJA6ZPry4Y2o0j6B0AxnFP-XYJ0ndnNdtnChdY2c35ezbTSV6MnTNf6LgaKFrhnuYp4tie1xnXDBKWYjhAsTk~jDS~Dtp2EkRWX3A__",
			PostedDate:    time.Now().UTC(),
			CarImg:        "https://carsguide-res.cloudinary.com/image/upload/f_auto,fl_lossy,q_auto,t_default/v1/editorial/story/hero_image/2024-Suzuki-Swift-EV-Hatchback-Teal-Render-Thanos-Pappas-1001x565-(1).jpg",
			CarModel:      "TOYOTA Swift",
			CarYear:       "2021",
			Price:         981,
			LikeCount:     "19",
			CommentCount:  "9",
			ViewCount:     "492",
		},
		{
			DealerName:    "Ko Zin Thant Win",
			DealerProfile: "https://s3-alpha-sig.figma.com/img/2714/df7a/1e9d94a74c92e29d5fdf1c828724f623?Expires=1744588800&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=MExi4KyWbv79dnOk820xl7UjfRQsyLmMYLdzkYgI-R2HwqCLGaCMAIaxh8cDrpJOvbhYs9chNUaEqze1FqPFncJcIoAOd557VH8GRMqd9Mz5aGimaYXq9zUQdZlSrfD5hgi6L~ijb16BCkUWuFUrCK9n~HjQtBX3qy~JaR0ECY2dMil~CdCRbww65M4nKCC5Yg9AvyPSsj-oW1O6zRN4ciG1OH-TbaeNoxXnIPoRSV1WQDL3GdJA6ZPry4Y2o0j6B0AxnFP-XYJ0ndnNdtnChdY2c35ezbTSV6MnTNf6LgaKFrhnuYp4tie1xnXDBKWYjhAsTk~jDS~Dtp2EkRWX3A__",
			PostedDate:    time.Now().UTC(),
			CarImg:        "https://carsguide-res.cloudinary.com/image/upload/f_auto,fl_lossy,q_auto,t_default/v1/editorial/story/hero_image/2024-Suzuki-Swift-EV-Hatchback-Teal-Render-Thanos-Pappas-1001x565-(1).jpg",
			CarModel:      "TOYOTA Swift",
			CarYear:       "2021",
			Price:         981,
			LikeCount:     "19",
			CommentCount:  "9",
			ViewCount:     "492",
		},
		{
			DealerName:    "Ko Arkar Lin",
			DealerProfile: "https://s3-alpha-sig.figma.com/img/2714/df7a/1e9d94a74c92e29d5fdf1c828724f623?Expires=1744588800&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=MExi4KyWbv79dnOk820xl7UjfRQsyLmMYLdzkYgI-R2HwqCLGaCMAIaxh8cDrpJOvbhYs9chNUaEqze1FqPFncJcIoAOd557VH8GRMqd9Mz5aGimaYXq9zUQdZlSrfD5hgi6L~ijb16BCkUWuFUrCK9n~HjQtBX3qy~JaR0ECY2dMil~CdCRbww65M4nKCC5Yg9AvyPSsj-oW1O6zRN4ciG1OH-TbaeNoxXnIPoRSV1WQDL3GdJA6ZPry4Y2o0j6B0AxnFP-XYJ0ndnNdtnChdY2c35ezbTSV6MnTNf6LgaKFrhnuYp4tie1xnXDBKWYjhAsTk~jDS~Dtp2EkRWX3A__",
			PostedDate:    time.Now().UTC(),
			CarImg:        "https://carsguide-res.cloudinary.com/image/upload/f_auto,fl_lossy,q_auto,t_default/v1/editorial/story/hero_image/2024-Suzuki-Swift-EV-Hatchback-Teal-Render-Thanos-Pappas-1001x565-(1).jpg",
			CarModel:      "TOYOTA Swift",
			CarYear:       "2021",
			Price:         981,
			LikeCount:     "19",
			CommentCount:  "9",
			ViewCount:     "492",
		},
	}

	responseWithJSON(w, 200, struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
		Data       []Car  `json:"car"`
	}{
		Message:    "car list successfully fetched",
		StatusCode: 200,
		Data:       car,
	})
}
