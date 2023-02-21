# busdes-kic-api(Japanese)
立命館大学衣笠キャンパスと駅を結ぶバス情報をリアルタイムに表示するアプリケーションです。
# 概要
## ライブラリ
* [echo](https://echo.labstack.com/)
* [colly](http://go-colly.org/)
* [go-cache](https://github.com/patrickmn/go-cache)
* [sqlite](https://github.com/mattn/go-sqlite3)
* [gorm](https://github.com/go-gorm/gorm)

## アーキテクチャ
* Clean Archtechture
    * Controller
    * Usecase
    * Repository
    * Domain

# インフラ
## 使用環境
busdes-kic-apiは、GCPのCloudRun(Google Cloud Plateform)上で動作しています。また、DNSにはCloudFlareを使用しています。

* GCP
    * Cloud Run
* Cloud Flare
    * DNS

## CI/CD
このシステムでは、GithubActionsを使用して継続的なインティグレーションと継続的なデリバリーを実現しています。

![構成図](.img/busdes-kic-api.png)



# busdes-kic-api(English)
This application displays real-time bus information between Ritsumeikan University Kinugasa Campus and the station.

# overview
## Library
* [echo](https://echo.labstack.com/)
* [colly](http://go-colly.org/)
* [go-cache](https://github.com/patrickmn/go-cache)
* [sqlite](https://github.com/mattn/go-sqlite3)
* [gorm](https://github.com/go-gorm/gorm)

# architecture
* Clean Archtechture
    * Controller
    * Usecase
    * Repository
    * Domain

# Infra
## environment
This System is running on GCP's CloudRun(Google Cloud Plateform). And we use CloudFlare for DNS.

* GCP
    * Cloud Run
* Cloud Flare
    * DNS

## CI/CD
The system uses GithubActions for continuous intigration and continuous delivery.

![構成図](.img/busdes-kic-api.png)
