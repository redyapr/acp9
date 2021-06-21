# ACP9 - Redy & Gigih

## Version: v1

### /register

#### POST
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

### /register/confirm/{userOTP}

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| userOTP | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

### /login

#### POST
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

### /categories

#### GET
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

### /products

#### GET
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

### /products/{categorySlug}

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| categorySlug | path |  | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

### /cart

#### POST
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

#### GET
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

### /cart/{cartId}

#### PUT
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| cartId | path |  | Yes | integer |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

#### DELETE
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| cartId | path |  | Yes | integer |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

### /checkout

#### POST
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |

### /payment

#### POST
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Successful operation |

##### Security

| Security Schema | Scopes |
| --- | --- |
| Authorization | |
