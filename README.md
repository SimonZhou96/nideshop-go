# nideshop-go

 [![Build Status][1]][2] [![Go Report Card][3]][4] [![MIT licensed][5]][6] 

[1]: https://travis-ci.org/harlanc/moshopserver.svg?branch=master
[2]: https://travis-ci.org/harlanc/moshopserver
[3]: https://goreportcard.com/badge/github.com/harlanc/moshopserver
[4]: https://goreportcard.com/report/github.com/harlanc/moshopserver
[5]: https://img.shields.io/badge/license-MIT-blue.svg
[6]: LICENSE

## Introduction

- [nideshop](https://github.com/tumobi/nideshop) backend service implemented by Golang
- Based on [Go-kit](https://github.com/go-kit/kit) to modulized the functionalities.



UI GitHub: [https://github.com/tumobi/nideshop-mini-program](https://github.com/tumobi/nideshop-mini-program)

## How to start

- Using `git clone` to clone the repository into the PC
    
        go clone https://github.com/harlanc/moshopserver
  
- Download the dependency

        go get -d ./...

- Create necessary database
      
        CREATE SCHEMA `nideshop` DEFAULT CHARACTER SET utf8mb4 ;

- Configure the information of the lightweight program
   
        [default]
        default_module='api'
        [weixin] 
        #小程序 appid
        appid=""
        #小程序密钥
        secret="" 
        #商户帐号ID
        mch_id='3' 
        #微信支付密钥
        apikey='4'
        #微信异步通知，例：https://www.nideshop.com/api/pay/notify 
        notify_url='5' 
        
-  Run the following command(Default port will be 8360)

        go run main.go

- The Complete Configuration(https://www.nideshop.com/documents/nideshop-manual/deployment-centos)


## Wechat lightweight program screenshot(Updating..)
<p align="center">
 <img src="http://qiniu.harlanc.vip/6.9.2019_5:41:56.png">
</p>
## Functionalities

- [x] Write data API services
- [x] Build the architecture of go-kit
- [x] Display the main page
- [x] Display good category
- [ ] Display Category
- [ ] Implement Login function
- [ ] Implement Add to Cart functionality

## Third-party dependency

- [go-kit](https://github.com/go-kit/kit)
- [wxpay](https://github.com/objcoding/wxpay)




