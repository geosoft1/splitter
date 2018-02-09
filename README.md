http splitter
====

[![license](https://img.shields.io/badge/license-gpl-blue.svg)](https://github.com/geosoft1/splitter/blob/master/LICENSE)

Receive a http GET request and send it to multiple microservices. Useful when you have a production server and want to receive the same request on development server. Another application is database realtime replication by a microservice. Can work with [reverseproxy](https://github.com/geosoft1/reverseproxy) project.

### How it works?

[![6e464537980306edd9f0df0d2d1c30ef1dbedd62](https://user-images.githubusercontent.com/6298396/36018301-efb3a37a-0d83-11e8-84d0-43320e1e6212.png)]

Just complete the `conf.json` file and run the server. Example:

     {
         "ip":"",
         "port":"8080",
         "handler":"/update",
         "routes": [
              "192.168.88.160:8080",
              "192.168.88.164:8000"
         ]
     }

### Configuration details

     "ip":"",

No ip mean `localhost` on hosting server. Is no need to change this, in most cases can miss.

     "port":"8080",

The server listening on this port. Remeber to forward the port `80` to this port if your connection pass through a router. No root right are required if you run on big ports (eg. `8080`).

### Handler

Mean what splitter listen for. This request with query string will be forward to destination routes. Target always start with `/`.

### Routes

Routes has the folowing structure

     "http://target:port"

Isn't a limit to define routes but usually you need two or three routes (a basic scenario include main server, backup and dev). Is better to keep this limit reasonable low.