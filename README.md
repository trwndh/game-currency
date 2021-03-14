# Game Currency API
Service for create currency, create conversion rate, get list of currencies, and get conversion rate result from 2 different currencies

## Quick Start
Let's run this service!

## Docker and docker-compose
* Docker ```v20.10.5``` 
* docker-compose ```1.28.5```

these versions are installed on my machine. make sure you install both

## Running service
* Clone this repo 
* Go to this repo folder 
* And just run 
``` docker-compose up```

```
example:
$ ~/workspace : git clone https://github.com/trwndh/game-currency.git
$ ~/workspace : cd game-currency
$ ~/workspace/game-currency : docker-compose up
```

## Documentation
- Please refer to openAPI 3 specs on ```api/v1/openapi/gamecurrency-http-api.yaml```
- or you can click [here](http://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/trwndh/game-currency/main/api/v1/openapi/gamecurreny-http-api.yaml) to easily read documentation for this service (using ReDoc)

---
**TLDR;**

__API server__ located on ```http://localhost:8081/game-currency/api/v1```

__Authorization key for header__ is ```Basic super-dupper-rahasia```

* Create new currency :  
```
POST /currencies
```
```
Body Payload (Json):

{
    "name": "Knut"
}
```
---
* Get All currencies
```
GET /currencies
```
---
* Create new conversion rate
```
POST /conversions
```
```
Body Payload (Json):

{
    "currency_id_from":1,
    "currency_id_to":2,
    "rate":29
}

```
---
* Calculate conversion amount between 2 currencies
```
GET /conversions?currency_id_from=2&currency_id_to=1&amount=580
```
---

## Tech used in this project
```
- Service : Golang 1.14
- Database: MySql 5.7 
```

## Database Documentation
```
here is some database documentation for this service
```
* ERD

<img src="https://raw.githubusercontent.com/trwndh/game-currency/main/docs/database/erd.png"/>

---

* Table currency 

| PK | Field      | Type      | Length | Nullable | Description                                  |
|----|------------|-----------|--------|----------|----------------------------------------------|
| Y  | id         | int       | 5      | N        | Identifier for each currency, auto increment |
|    | name       | varchar   | 25     | N        | Currency name, cannot be blank. UNIQUE       |
|    | created_at | timestamp |        | Y        | default CURRENT_TIMESTAMP                    |
|    | updated_at | timestamp |        | Y        | default CURRENT_TIMESTAMP, auto update       |

Indexes : ```id (PRIMARY), Name (UNIQUE)```

---

* Table conversion_rate

| PK | Field            | Type      | Length | Nullable | Description                                  |
|----|------------------|-----------|--------|----------|----------------------------------------------|
| Y  | id               | int       | 5      | N        | Identifier for each currency, auto increment |
|    | currency_id_from | int       | 5      | N        | source of ID currency when add new rate      |
|    | currency_id_to   | int       | 5      | N        | destination of ID currency when add new rate |
|    | rate             | int       | 5      | N        | currency rate from 2 different currencies    |
|    | created_at       | timetamp  |        | Y        | default CURRENT_TIMESTAMP                    |
|    | updated_at       | timestamp |        | Y        | default CURRENT_TIMESTAMP, auto update       |

Indexes : ``` id (PRIMARY), currency_id_from (index name: currency_id_from), currency_id_to (index name: currency_id_to)```

Foreign: ```currency_id_from > currency.id, currency_id_to > currency.id```

---

* SQL file for migrations available at ``` mysql-dump/init_tables.sql ```

---

## Testing
* use `make coverage` or `make test` for running unit test
```
example:
$ ~/workspace/game-currency : make coverage
```
