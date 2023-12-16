# ZopSmart--task

A Simple Library management system with end-points for CRUD operations

Postman collection:
https://www.postman.com/joint-operations-observer-83521299/workspace/zopsmart-task/collection/31312207-c9341749-3396-4f39-8d3d-fdac5f5af199?action=share&creator=31312207

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Postman Screenshots](#postman-screenshots)
-  [System design](#System-Design-for-basic-Library-management-system)

## Getting Started

Welcome to LMS, your go-to application for managing and tracking your Books effortlessly.


### Prerequisites

- [Go](https://golang.org/) (Programming Language)
- [GoFr](https://gofr.dev/) (Framework)
- [MySQL](https://www.mysql.com/) (or any preferred database)
- [Git](https://git-scm.com/) (Version Control System)


### Installation

step-by-step instructions on how to install and set up.

- git clone [https://github.com/Ronakkumbhat/ZopSmart--task.git]
- go run main.go 

## Usage

### View All Books:

- Explore the "/getstatus" route to view a list of all your books.

### add book:

- Check the "/addbook" route to add a book to your database.

### Update a book:

- Visit the "/updatestatus" route to modify details of a specific book using its name.

### Delete a book:

- Visit the "/deletebook" route to delete details of a specific book using its name.

## Postman Screenshots

### Add book:

1. Open Postman and set up a `POST` request to the "addbook" endpoint.
2. Provide the {"Name": "ronak"} in the request body.
3. Send the request.

  ![image](https://github.com/Ronakkumbhat/ZopSmart--task/assets/91602958/43fb5de4-186d-45d4-83b1-9911e4f11975)


### Delete book:

1. Open Postman and set up a `DELETE` request to the "deletebook" endpoint.
2. Provide the {"Name": "ronak"} in the request body.
3. Send the request.

  ![image](https://github.com/Ronakkumbhat/ZopSmart--task/assets/91602958/2e17287b-e33b-439e-a628-af64d6a42a05)

### View  books:

1. Set up a `GET` request to your "/getstatus" endpoint.
2. 2. Provide the {"Name": "ronak"} in the request body.
3. Send the request.

  ![image](https://github.com/Ronakkumbhat/ZopSmart--task/assets/91602958/23e71e05-1f91-4a8d-b7cc-848765893f79)


### Update a status:

1. Configure a `PUT` request to your "/updatestatus" endpoint.
2. Adjust the request body with the updated details.
3. Send the request.

  ![image](https://github.com/Ronakkumbhat/ZopSmart--task/assets/91602958/b5448a0c-5a0a-471c-a2b3-b9cb741c8438)

### System Design for basic Library management system
![image](https://github.com/Ronakkumbhat/ZopSmart--task/assets/91602958/0ca590f2-c41f-4314-96d4-d1415e3a62e6)


