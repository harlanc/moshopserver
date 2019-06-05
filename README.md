# moshopserver小程序商城后台系统

## 介绍

- [nideshop](https://github.com/tumobi/nideshop)的golang实现
- 基于[beego](https://github.com/astaxie/beego)开发



本项目需要配合微信小程序端使用，GitHub: [https://github.com/tumobi/nideshop-mini-program](https://github.com/tumobi/nideshop-mini-program)

## 测试环境搭建

- 克隆源码到$GOPATH目录下
    
        go clone https://github.com/harlanc/moshopserver
  
- 下载所有依赖包

        go get -d ./...

- 创建数据库nideshop并导入项目根目录下的nideshop.sql
      
        CREATE SCHEMA `nideshop` DEFAULT CHARACTER SET utf8mb4 ;

- 配置好小程序相关字段
-  运行以下命令（默认为开启8080端口）

        go run main.go

- 小程序的配置[参考最后一节](https://www.nideshop.com/documents/nideshop-manual/deployment-centos)


## 微信小程序客户端截图

![首页](http://upload-images.jianshu.io/upload_images/3985656-c543b937ac6e79bb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/320)

![专题](http://upload-images.jianshu.io/upload_images/3985656-bd606aac3b5491c2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/320)

![分类](http://upload-images.jianshu.io/upload_images/3985656-fa9565158376d439.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/320)

![商品列表](http://upload-images.jianshu.io/upload_images/3985656-788b7fd2c4a558d0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/320)

![商品详情](http://upload-images.jianshu.io/upload_images/3985656-99a6e0a57778d85f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/320)

![购物车](http://upload-images.jianshu.io/upload_images/3985656-60ff2307d81f6bb2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/320)

![订单中心](http://upload-images.jianshu.io/upload_images/3985656-dff837e6b2ec87b3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/320)


## 功能列表
+ 首页
+ 分类首页、分类商品、新品首发、人气推荐商品页面
+ 商品详情页面，包含加入购物车、收藏商品、商品评论功能
+ 搜索功能
+ 专题功能
+ 品牌功能
+ 完整的购物流程，商品的加入、编辑、删除、批量选择，收货地址的选择，下单支付
+ 会员中心（订单、收藏、足迹、收货地址、意见反馈）
....


## 第三方依赖包

- [beego](https://github.com/astaxie/beego)
- [go-sql-driver](https://github.com/go-sql-driver/mysql)
- [go.uuid](https://github.com/satori/go.uuid)

## 联系方式

有疑问可以加我：

![](http:/qiniu.harlanc.vip/6.5.2019_8:00:27.png)


