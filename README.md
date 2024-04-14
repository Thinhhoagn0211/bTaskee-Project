# bTaskee-Project

A **bTaskee-Project** is a project which consists of a collection of small, autonomous **services** about order service, calculate costs and send job requests to company employees immediately. Each **service** is self-contained and should implement a single business capability. 

There are four microservices in this project:
* [**Booking Service**](./src/booking-microservice/)
* [**Pricing Service**](./src/pricing-microservice/)
* [**Product Service**](./src/product-microservice/)
* [**Sender Service**](./src/sending-microservice/)

## Services

No. | Service | URI | Swagger URI
--- | --- | --- | ---
1 | product service | [http://localhost:5000](http://localhost:5000) | [http://localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html)
2 | booking service | [http://localhost:5001](http://localhost:5001) | [http://localhost:5001/swagger/index.html](http://localhost:5001/swagger/index.html)
3 | pricing service | [http://localhost:5002](http://localhost:5002) | [http://localhost:5002/swagger/index.html](http://localhost:5002/swagger/index.html)


## Architecture of project
![Microservices](./diagram.png)
## Starting project
We have two ways to run these microservice (by run local or by docker-compose)

##### - Change configuration file
<em>Update values of file **./src/{microserice_name}/config/config.json**</em>
```sh
{
    "port": "127.0.0.1:8808",
    "mgAddrs": "127.0.0.1:27017",
    "mgDbName": "companyDatabase",
    "mgDbUsername": "",
    "mgDbPassword": "",

    "rabbitMQAddr": "amqp://guest:guest@127.0.0.1:5672/",
    "queueName": "companyQueue"
}
```
**NOTE:** We should replace address 127.0.0.1 into your ip address in this file of each microservice before you build with docker and we must import file json to init data to mongoDB [database](./metadata/)

- Running with docker-compose: Jump into [`Makefile`](./Makefile), then

```bash
> make build
> make run
```



* Run the <strong>Product Service, Pricing Service, Booking Service, Sending Service</strong> 
```sh
$ cd ./src/{microservice}
$ go run main.go
>> [GIN-debug] Listening and serving HTTP on :port
```
