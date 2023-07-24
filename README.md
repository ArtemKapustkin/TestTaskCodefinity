# FizzBuzz Quiz Solver

## Overview

This project is a simple web application that runs a Fiber server, allowing clients to interact with a FizzBuzz Solver. 
The server listens on port 5000 and provides a single POST endpoint `/run`, which executes the FizzBuzz Solver and returns the resulting array of strings, which length is based on the provided input parameter 'N'.

## FizzBuzz Rules

The FizzBuzz quiz is a simple programming task that follows these rules:

* For numbers divisible by 3, replace the number with "Fizz".
* For numbers divisible by 5, replace the number with "Buzz".
* For numbers divisible by both 3 and 5, replace the number with "FizzBuzz".
* For other numbers, keep the number as is.

## Usage

You can run server locally using `go run main.go` command in your terminal.
Or you also can easily run this project using `docker compose up`, which simplifies the setup process.

## API Endpoint

### `POST /run`

This endpoint accepts a JSON request with a single parameter 'n', which represents the number of strings in the FizzBuzz solver result array. 
The application then calculates the FizzBuzz quiz for the numbers from 1 to 'n' and returns the resulting array of strings.

### Request

To test the FizzBuzz quiz solver API, you can use the curl command-line tool to make POST requests, like this:

`curl -X POST http://localhost:5000/run -d "n=15"`

This command sends a JSON request with the parameter n set to 15 to the server's /run endpoint, and it will return the FizzBuzz solver result for the numbers from 1 to 15.
