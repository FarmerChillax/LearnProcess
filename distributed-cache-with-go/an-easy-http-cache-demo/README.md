# 基于HTTP服务的简单缓存

## 接口设计

通过`HTTP`协议的`GET/PUT/DELETE`等方法实现缓存的`Get/Set/Del`等基本操作

### GET方法

Get操作用于查询某个键并获取其值，对应`http`的`GET`方法

- endpoint: `/cache/<key>`
- methods: `GET`
- response body: `<value>`
- response encode type: `json`


### PUT方法



### DELETE方法

