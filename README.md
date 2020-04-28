# 项目介绍
小程序服务端示例代码

# 开始之前

## 本地部署说明

### 参数替换
如果要测试真实的收银台服务端，请替换如下参数:
1. 替换 `conf/pay.json`文件中的`deal_id`, `app_key`, `app_id` 为百度收银台真实数据，参考[支付管理后台操作指引
](https://smartprogram.baidu.com/docs/introduction/background-guide/) 。
2. 替换 `conf/platform_rsa_public_key.pem` 和 `conf/self_rsa_private_key.pem` 为[百度收银台支付开通指引](https://smartprogram.baidu.com/docs/introduction/pay/) 中生成的公钥和私钥。

### 运行
请将代码库拉取下来，放入 go_path 下运行


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
