# Game Currency API
Service for create currency, create conversion rate, get list of currencies, and get conversion rate result from 2 different currencies

## Documentation
- Please refer to openAPI 3 specs on ```api/v1/openapi/gamecurrency-http-api.yaml```
- or you can click [here](http://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/trwndh/game-currency/main/api/v1/openapi/gamecurreny-http-api.yaml) to easily read documentation for this service (using ReDoc)

<a href="http://google.com" target="_blank">new tab</a>

---
**TLDR;**

__Api server__ located on ```http://localhost:8081/game-currency/api/v1```

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

## Tech used in this project
```
- Service : Golang 1.14
- Database: MySql 5.7 
```

## Quick Start
Let's run this service!

## Docker and docker-compose
* Docker ```v20.10.5``` 
* docker-compose ```1.28.5```

these versions are installed on my machine.

## Running service
* After clone this repo, go to this repo folder and run 
``` docker-compose up```

```
example:
$ ~/workspace : git clone https://github.com/trwndh/game-currency.git
$ ~/workspace : cd game-currency
$ ~/workspace/game-currency : docker-compose up
```

## Testing
* use `make coverage` or `make test` for running unit test
```
example:
$ ~/workspace : make coverage
```
