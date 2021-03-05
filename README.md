# fileviewer

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#usage">Usage</a></li>  
        <li><a href="#example">Example</a></li> 
      </ul>
    </li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

This service is used to view the static files in server.

## getting-started

### Built With
 golang
 
### installation
 
 step 1. download from <a href=https://github.com/smutil/fileviewer/releases>releases</a>. 
 
 step 2. start the service as shown below
 
 ```
 ./fileviewer -dest /tmp/configs
 ```
 
### usage

``` 
  ./fileviewer -h
    --version
          returns application version
    --dest string
          (required) destination/working directory, should not be root /
    --port string
          overwrite default port (default "3000")
    --tls-crt string
          certificate path, only needed for ssl service
    --tls-key string
          key path, only needed for ssl service
    --log-file
          writes all the log messages to a file
 ```
 
 ### example

  1. view the static file in given working directory
  ```
  curl  http://localhost:3000/
  ```

  2. returns application health check
  ```
  curl http://localhost:3000/health
  ```

  3. Metrics for Prometheus
  ```
  curl http://localhost:3000/metrics
  ```
