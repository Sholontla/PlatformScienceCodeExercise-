<div id="top"></div>

<!-- Structure Demo Spring -->
<br />
<div align="center">
  <a href="https://github.com/Sholontla/Platform-Science-Code-Exercise-">
  </a>

<h3 align="center">Platform Science Code Exercise</h3>

  <p align="center">
   This Structure Demo is use for testing and demostration go to for running instructions: https://github.com/Sholontla/admin-store-manager/blob/main/arch-high-level-overview.pdf
    <br />
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->

## About The Project

  <p align="right">(<a href="#top">back to top</a>)</p>

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

- File Server Service (file_server/):

  Project OverView:
  . This service is used to created and store the dummy creation and files.
  . Send the data if needed to throw websocket client and use fleet_service to process some logic business.
  .

- Fleet Service (fleet_service/):

  Project OverView:
  . This service manage the "logic Platform Science Code Exercise" described in the "SDE Code Exercise".
  . The service have one websocket server that handle the incoming data from the client "file_server".
  . Then, process the data and apply the logic mention before, and sends the data through the websocket client to the analytics service that render the data into html / bootstrap frontEnd.

- Analytics Service (analytics/):

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

## License

For testing and demostrations purposes.

<!-- CONTACT -->

## Contact

Gerardo Ruiz Bustani - solbustani@gmail.com - 442 488 6193
