# 项目介绍
小程序服务端示例代码

# 开始之前
## 说明
 1. 本示例基于golang实现，主要展示对接oauth2.0授权登录过程以及对接百度收银台支付服务时,小程序开发者需要自行实现的服务端接口；
 2. 本示例仅建议在开发机进行本地部署，配合相关[示例小程序](https://github.com/baidu-smart-app/auth-pay-demo-frontend)在开发者工具里进行调试；
 3. 如果需要在真机进行调试，建议开发者自行实现相关服务端功能，提供广域网接口，并参考[示例小程序](https://github.com/baidu-smart-app/auth-pay-demo-frontend)配置说明将服务域名替换为真实域名；  
 
## 免责声明
 本示例仅供调试参考，不具备真实的业务处理能力，具体业务逻辑请开发者根据实际业务需求自行实现。
 
## 本地部署说明
### 运行
1. 参考 github 指示，将代码拉取到本地；
2. 按照【参数替换】说明替换配置参数；
3. 在代码根目录运行 go mod tidy 安装相关依赖；（golang环境自行安装，建议golang版本>=1.13）
4. 在代码根目录运行 go run main.go
5. 第4步启动后使用浏览器直接访问 http://127.0.0.1:8080/, 如果返回404说明服务启动成功；
6. 按照[示例小程序文档](https://github.com/baidu-smart-app/auth-pay-demo-frontend)指示，将第5步中的服务域名(127.0.0.1:8080)配置到小程序示例中，即可在开发者工具体验登录、订单相关功能；

### 参数替换 

1. 测试登录相关服务，替换`conf/smart_app.json` 文件中的 `{app_key}`, `{secret_key}`；app_key为智能小程序的 AppKey,示例：4fecoAqgCIUtzIyA4FAPgoyrc4oUc25c;secret_key 为智能小程序的 AppSecret 从开发者平台中获取; 
2. 测试订单相关业务，替换 `conf/pay.json`文件中的`{deal_id}`, `{app_key}`, `{app_id}` 为百度收银台真实数据，参考[支付管理后台操作指引](https://smartprogram.baidu.com/docs/introduction/background-guide/)  
3. 测试订单相关业务，替换 `conf/platform_rsa_public_key.pem` 和 `conf/self_rsa_private_key.pem` 为[百度收银台支付开通指引](https://smartprogram.baidu.com/docs/introduction/pay/) 中生成的公钥和私钥。


# 项目模块

### 支付示例

#### 创建订单

##### Path
/pay/gen

##### Method
GET

#### Params
无

##### 正常返回
```
{
    "code": 0,
    "msg": "succ",
    "data": {
        "appKey": "MMMzUX",
        "dealId": "470193086",
        "dealTitle": "支付示例",
        "rsaSign": "sD9aQUrZJB/9+++YmQxAS7tZ0/905P+ql7q7YCcFdiFNw/KBJUcI/NhVlW/ov0dawggrHwdT6THA9gwCip21k8Mr23fuV0m2sLOLsWxJOgCfs2DaqfCu76TZC/qPzWvXEWX/7A2sKYHLxoslgYA/otcTfH7zRANwr6JM5Yknt24=",
        "totalAmount": "1",
        "tpOrderId": "16944676582",
        "bizInfo": "{\"tpData\":{\"appKey\":\"MMMzUX\",\"dealID\":\"470193086\",\"tpOrderID\":\"16944676582\",\"rsaSign\":\"sD9aQUrZJB/9+++YmQxAS7tZ0/905P+ql7q7YCcFdiFNw/KBJUcI/NhVlW/ov0dawggrHwdT6THA9gwCip21k8Mr23fuV0m2sLOLsWxJOgCfs2DaqfCu76TZC/qPzWvXEWX/7A2sKYHLxoslgYA/otcTfH7zRANwr6JM5Yknt24=\",\"totalAmount\":\"1\",\"payResultUrl\":\"\",\"returnData\":null,\"dealTitle\":\"\",\"detailSubTitle\":\"支付示例\",\"dealTumbView\":\"\",\"displayData\":{\"cashierTopBlock\":[[{\"leftCol\":\"订单名称\",\"rightCol\":\"智能小程序支付实例16944676582\"},{\"leftCol\":\"数量\",\"rightCol\":\"1\"},{\"leftCol\":\"订单金额\",\"rightCol\":\"0.01元\"}],[{\"leftCol\":\"服务地址\",\"rightCol\":\"北京市海淀区上地十街10号百度大厦\"}]]}},\"orderDetailData\":null}"
    }
}
```

#### 查询订单

##### Path
/pay/status

##### Method
GET

#### Params
- tp_order_id
    - type: string
    - must: true
    - explain: /pay/gen 下发的tpOrderID

##### 正常返回
```
{
    "code": 0,
    "msg": "succ",
    "data": {
        "payStatus": {
            "statusNum": 1, //-1未支付,1支付成功
            "statusDesc": "支付成功"
        },
        "refundStatus": {
            "statusNum": -1, //-1未退费,1退费中,2退费成功,9退费失败
            "statusDesc": "未退费"
        },
        "verification":{
           "statusNum": -1, //-1未核销,1已核销
           "statusDesc": "无核销数据"
       }
    }
}
```

#### 申请退款

##### Path
/pay/refund

##### Method
GET

#### Params
- tp_order_id
    - type: string
    - must: true
    - explain: /pay/gen 下发的tpOrderID

##### 正常返回
```
{
    "code": 0,
    "msg": "succ",
    "data": {
         "refundBatchId": "152713835",//平台退款批次号
        "refundPayMoney": "9800" //平台可退退款金额【分为单位】
    }
}
```

#### 支付回调（内部）



### 用户示例

####  用户登录
swan.login 回调记录openid和换session key

##### Path
/auth/login

##### Method
GET

#### Params
- code
    - type: string
    - must: true
    - explain: swan.login 得到的code

##### 正常返回
```
{
    "code": 0,
    "msg": "succ",
    "data": {
        "open_id": "用户标识"
    }
}
```

### 用户数据解密
swan.getUserInfo 回调 进行解密 + 账户管理示例

##### Path
/auth/userinfo

##### Method
POST

#### Params
- data
    - type: string
    - must: true
    - explain: swan.getUserInfo 得到的data
- iv
    - type: string
    - must: true
    - explain: swan.getUserInfo 得到的iv
- open_id
    - type: string
    - must: true
    - explain: /auth/login 下发的用户标识

##### 正常返回
```
{
    "code": 0,
    "msg": "succ",
    "data": {
        "openid":"open_id",
        "nickname":"baidu_user",
        "headimgurl":"url of image",
        "sex":1
    }
}
```

### 用户手机号解密
swan.getPhoneNumber 回调 进行解密 + 账户管理示例

##### Path
/auth/phone

##### Method
POST

#### Params
- data
    - type: string
    - must: true
    - explain: swan.getUserInfo 得到的data
- iv
    - type: string
    - must: true
    - explain: swan.getUserInfo 得到的iv
- open_id
    - type: string
    - must: true
    - explain: /auth/login 下发的用户标识

##### 正常返回
```
{
    "code": 0,
    "msg": "succ",
    "data": {
        "我也不知道key是什么":"185777777777",
    }
}
```