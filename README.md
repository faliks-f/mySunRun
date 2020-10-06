<!--
 * @Author: your name
 * @Date: 2020-10-06 18:37:45
 * @LastEditTime: 2020-10-06 18:42:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \undefinede:\mySunRun\README.md
 -->
# 内部使用，不要外传    
account.txt里放姓名和IMEICode，中间用空格隔开   
使用需要有python解释器和golang编译器环境   
## 如何运行  
命令行进入文件夹目录，执行
```
python SunRun.py  
```  

# SunRun
本repo为阳光长跑的Go实现  
实话说 实现原理已经烂大街了
适用于阳光长跑  
阳光体育服务平台

# 抓包
## 安卓端
使用`packet capture`抓取登录包即可  
打开主界面之后的一般为第一个包  

### 一般长这样

```
{
  "Success": true,
  "Data": {
    "Token": "0d3bf9578beb4*******a9384272a488",
    "UserId": 619***,
    "IMEICode": "03df3a1af37848******7c70a085d55d",
    "AndroidVer": 2.1*,
    "AppleVer": 1.2*,
    "WinVer": 1.0
  }
}
```

## ios端
目前好像接口失效了  
参见魔盒运动的一个神奇方法(bundle id)  
自行搜索吧 以前用的抓包应用叫`stream`

# 使用方法
`git clone git@github.com:Licsber/SunRun.git`  
`cd SunRun`  
`vim main.go`  
在这里调用`run.go`的`justRun`方法即可  
参数`imeicode`为抓包获取  
参数`distance`为学校要求  
一般男生2400 女生2000即可

# 关于
适用的学校如下

```
【1】武汉大学
【2】武汉工程大学
【3】安徽师范大学
【4】安徽国防科技职业学院
【5】南京科技职业学院
【6】南京晓庄学院
【7】南通师范高等专科学校
【8】南昌大学
【9】南昌航空大学
【10】宁夏大学
【11】南京工程学院
【12】井冈山大学
【13】山西大学商务学院
【14】武汉纺织大学
【15】扬州大学
【16】江南大学
【17】南京农业大学
【18】成都工业学院
【19】保定电力职业技术学院
【20】包头医学院
【21】内蒙古医科大学
【22】安徽财贸职业学院
【23】南京工业大学
【24】辽宁工程技术大学
【25】南京旅游职业学院
【26】黎明职业大学
【27】丽水学院
【28】新华学院
【29】铜川职业技术学院
【30】太原学院
【31】厦门大学
【32】江西师范大学科学技术学院
```

# 免责声明
代码仅供研究与学习交流之用  
另外鄙视一下利用信息差牟利的同学
