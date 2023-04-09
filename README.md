<div id="top"></div>

<!-- Structure Demo  -->
<div align="center">
  <a href="https://github.com/Sholontla/Platform-Science-Code-Exercise-">
  </a>
<h3 align="center">Platform Science Code Exercise</h3>
  <p align="center">
<br />
  </p>
</div>

<!-- RUN THE PROJECT -->

## Run Project

1.  Download the repository

2.  In the root directory in your local filesystem where the project been downloaded, run the Makefile command "make" to build the project and run the docker contaners.

3.  Use any http client:
    . POST method: - http://localhost:1001/service/run/process/5/second

    - common route for file_server: http://localhost:1001/service/run/- the process will run 5 seconds /5/second and can be change by any second/minute/hour/day need for the tests.

4.  go to:
    localhost:1000
    And see the data result from the process in the forntEnd.

5.  go to:
    localhost:1000/profile
    Dev Profile.

This Structure Demo is use for testing and demostration go to see few more details in the process with using diagrams:
diagram 1. https://github.com/Sholontla/PlatformScienceCodeExercise-/blob/master/Diagrama%20en%20blanco%20-%20P%C3%A1gina%201%20(3).jpeg

diagram 2. https://github.com/Sholontla/PlatformScienceCodeExercise-/blob/master/Diagrama%20en%20blanco%20-%20P%C3%A1gina%201%20(4).jpeg

<br />

<!-- ABOUT THE PROJECT -->

## About The Project

Project structure by:

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

### Project Overview

- File Server Service (file_server/):

  PORT: ws/1001

  Project OverView:
  . This service is used to created and store the dummy creation and files.
  . Send the data if needed to throw websocket client and use fleet_service to process some logic business.
  .

- Fleet Service (fleet_service/):

  PORT: ws/2001

  Project OverView:
  . This service manage the "logic Platform Science Code Exercise" described in the "SDE Code Exercise".
  . The service have one websocket server that handle the incoming data from the client "file_server".
  . Then, process the data and apply the logic mention before, and sends the data through the websocket client to the analytics service that render the data into html / bootstrap frontEnd.

<br />

- Analytics Service (analytics/):

### Built With

- Python
- FastAPI (http framework)
- websockets (sockets)

Virtualization / Containers

- Docker
- Docker - Compose

  PORT: 1000

  Project OverView:
  . This service manage the front end analytics service rendered the data coming from the the "file_server" and "fleet_service"

## About in general the improvements.

This project is the high level implementation.
many improvments need it.

1. some new frontEnd frameworks (react, vue, etc)
2. if is robust change all the data analysis to python.
3. Manage better concurrency, performance and channels.
4. gandle better context and different concepts to improve performance of servers time responses, wait until server / client send/recieve data.
5. the file_server can be improve create multiples data dummy producers to multiple fleet_services and combine with improvment number 3 (manage better concurrency) to test performance services and stream, test escability, load balancers monitoring and more.

And thanks to the team many other improvements can be implemente.

<br />

## License

For testing and demostrations purposes.

<!-- CONTACT -->

<br />

## Contact

Gerardo Ruiz Bustani - solbustani@gmail.com - 442 488 6193

<p align="right">(<a href="#top">back to top</a>)</p>
