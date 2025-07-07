# API文档

## 概述
提供用户管理功能的REST API

## 认证
使用Bearer Token认证：
`Authorization: Bearer YOUR_API_KEY`

## 账号管理

### 获取用户
`GET /users/{id}`

**请求参数**

| 参数 | 类型 | 必填 | 描述 |
|------|------|------|------|
| id   | int  | 是   | 用户ID |

**响应示例**
```json
{
  "id": 123,
  "name": "John Doe"
}
```

**错误代码**

| 状态码 | 描述 |
|------|------|
|401	|未授权|
|404	|用户不存在|
