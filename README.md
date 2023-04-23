<div id="top"></div>

<!-- Structure Demo  -->
<div align="center">
  <a href="https://github.com/Sholontla/PlatformScienceCodeExercise-">
  </a>
<h3 align="center">Platform Science Code Exercise</h3>
  <p align="center">
<br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->

## About the project:

<!-- RUN THE PROJECT -->

## Run Project

1.  clone the repository

2.  In the root directory in your local file system where the project been cloned, run the Makefile command "make" to build the project and run the docker containers.
    \*\*\* Note: This commmand will work in wsl/2 and Linux systems

3.           CREATE PUBLISHER/USER:
        POST http://localhost:1004/service/access/register/user
        {
        "user_user_name": "user1",
        "user_user_last_name": "user1",
        "user_user_email": "user1",
        "password": "user1",
        "permissions": {
        "create_topic": true,
        "read_topic": false
        },
        "role": true
        }

4.           USER LOGIN:
        POST http://localhost:1004/service/access/login/user
        body:
        {
        "user_user_email": "user1",
        "password": "user1"
        }

5.            CREATE PUBLISH/ORDER
          POST http://localhost:1004/service/create/order
          body:
          {
          "store": {
          "region": "test1",
          "sub_region": "test1",
          "sale": {
          "product": "test1",
          "price": 45.25,
          "cost": 8.25,
          "unit_sold": 5,
          "region": "test1",
          "sub_region": "test1"
          }
          }
          }

6.            VISUALIZE THE DATA
          localhost:3000/

<br />

Project structure by:

## Publisher Service (publisher_service/):

### Built With

    - Golang (Go)
    - Fiber (http framework)
    - gorilla/websocket (sockets)

    Virtualization / Containers

    - Docker
    - Docker - Compose

    O/I

    - Windows(WSL2)
    - Linux

    Project OverView:
      * Create topics by http request
      * kafka Broker to create and send topic throw a service to be consumed by financa_service/.


    publisher_service have 4 end-point:

    1.  This endpoint register a service/access/register/user
        \*\*\* Noticed that we have permissions and in this service, the next enpoints will only work with
        "create_topic": true and "role": true

        POST http://localhost:1004/service/access/register/user
        body:
        {
        "user_user_name": "user1",
        "user_user_last_name": "user1",
        "user_user_email": "user1",
        "password": "user1",
        "permissions": {
        "create_topic": true,
        "read_topic": false
        },
        "role": true
        }

    2.  This service/access/login/user
        will Login the publisher/user into the service.

        POST http://localhost:1004/service/access/login/user
        body:
        {
        "user_user_email": "user1",
        "password": "user1"
        }

    3.  the endpoint service/create/order will set a topic into the service.
        The service will have a Kfka Broker service where the finaince_service/ will consume
        the message from the publisher.
        Benefits of using Kafka Broker:
        - Ensure high concurrency, low concurrency and scalibilty
        - If the consumer goes down. we ensure the messages will not be lost.
        - Amoung others ...
          POST http://localhost:1004/service/create/order
          body:
          {
          "store": {
          "region": "test1",
          "sub_region": "test1",
          "sale": {
          "product": "test1",
          "price": 45.25,
          "cost": 8.25,
          "unit_sold": 5,
          "region": "test1",
          "sub_region": "test1"
          }
          }
          }

<br />

## Finance Service (finance_service/):

### Built With

    - Golang (Go)
    - Fiber (http framework)
    - gorilla/websocket (sockets)

    Virtualization / Containers

    - Docker
    - Docker - Compose

    O/I

    - Windows(WSL2)
    - Linux

Project OverView:
. This service use some basic logic to apply "Daily Revenue" for every topic/order consume by Kafka queues and
rednder the data into a timeline chart using CORS and Next.js as a primary forntEnd framework
. Also use Redis as Caching system to presist the data from the cosumer side.

Post
http://localhost:1003/finance/:param
param:
revenue
avarage
avarage_product
top_selling
profit_margin_all
daily_cost

    with this parameters listed yo can get some logic apply to the data.

Get
http://localhost:1003/daily/revenue
This endpoint list all the data store un Cache redis abd this same endpoint is used in the front End to list all revenue
process by the service.
For the moment only one chart is render into the frontend this only for demo purposes.

The service have implmented the next finance operations:

- CalculateDailyRevenueService (this function is the only working for now)
- CalculateAverageRevenueService
- CalculateAverageRevenuePerProductService
- IdentifyTopSellingProductsService
- CalculateAllProfitMarginService
- CalculateDailyCostService
- CalculateGrossProfitService
- CalculateGrossProfitMarginService
- AnalyzeSalesTrendsService
- CalculateAverageDailyRevenueService
- CalculateStoreRevenueService
- AnalyzeProfitabilityByRegionService
- IdentifyUnderperformingProductsService
- AnalyzePricingStrategyService
- ForecastFutureSales

<br />

## Front End Finance Service (front-end-finance):

### Built With

    - javaScript
    - Next.js

    Virtualization / Containers

    - Docker
    - Docker - Compose

    O/I

    - Windows(WSL2)
    - Linux

Project OverView:
. This is the front end to render the data and chart comign from the backend.

http://localhost:3000/

<br />

## License

For testing and demostrations purposes.

<!-- CONTACT -->

<br />

## Contact

Gerardo Ruiz Bustani - solbustani@gmail.com - +52 442 488 6193

<p align="right">(<a href="#top">back to top</a>)</p>
