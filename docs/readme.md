# Wanderer API

Wanderer API is a RESTful API designed for booking travel packages with features for reservation and payment. Additionally, it includes an admin panel for efficient content management within the system.


## Features

- See destination lists
- See the details of the tour package like the facility and the airline
- Book a tour packages
- Payement for a tour packages
- Give a review of the tour package
- Admin dashboard to manage the content of the system


## ERD

![ERD](https://raw.githubusercontent.com/Kelompok-2-Wanderer/Wanderer-BE/docs/docs/erd.png)


## API Spec

Please visit [SwaggerHub](https://app.swaggerhub.com/apis-docs/GALIHP83/Wanderer/1.0.0#/) or [Postman Workspace](https://www.postman.com/herusetiawans/workspace/alta-wanderer) to see the API spesification.

## Requirement

Things you need to run the Wanderer API
1. **Cloudinary Account**
2. **Midtrans Account**

## Installation

Follow these steps to install and set up the Wanderer API:
1. **Clone the repository:**

   ```bash
   git clone https://github.com/Kelompok-2-Wanderer/Wanderer-BE.git
   
2. **Move to Cloned Repository Folder**

    ```bash
    cd Wanderer-BE
    
3. **Update dependecies**
    
    ```bash
    go mod tidy

4. **Create a database** 

5. **Copy `.env.example` to `.env`**

    ```bash
    cp .env.example .env
    
6. **Configure yout `.env` to configure JWT token, connect to your database, cloudinary, and Midtrans**
7. **Run Wanderer API** 
8. 
    ```bash
    go run .
