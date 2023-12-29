# Wanderer API
<div align="center">
  <h1>Welcome to Wanderer</h1>

<!-- PROJECT LOGO -->

<img src="https://res.cloudinary.com/dhxzinjxp/image/upload/v1703555074/asset-default/Wanderer_dg29rh.svg" alt="Logo" width="500" height="auto" />
</div>

## 📑 About the Project
Wanderer API is a RESTful API designed for booking travel packages with features for reservation and payment. Additionally, it includes an admin panel for efficient content management within the system.

## 🌐 Deployment
 - [VERCEL](https://wanderer-delta.vercel.app/)

## 🖼 Prototype
- [FIGMA](https://www.figma.com/file/QjFROWypWKpnjN3AZgYpWS/Wanderer-App?type=design&node-id=0-1&mode=design&t=96eQ9N3axnjnLPgv-0)

## 🤝 Collaboration
- [GitHub (Version Control System Platform)](https://github.com/Kelompok-2-Wanderer)
- [Discord](https://discord.com/)

### ⚙ Backend
- [Github Repository for the Backend team](https://github.com/Kelompok-2-Wanderer/Wanderer-BE)
- [Swagger OpenAPI](https://app.swaggerhub.com/apis-docs/GALIHP83/Wanderer/1.0.0#)
- [Postman Workspace](https://www.postman.com/herusetiawans/workspace/alta-wanderer) to see the API spesification.

## 🔮 Features
- Login
- Register
- Payment
- Import file
- Export file

### 🌟 As User

- See destination lists
- See the details of the tour package like the facility and the airline
- Book a tour packages
- Payement for a tour packages
- Give a review of the tour package
- See user profile
- Update profile

### ✨ As  Admin

- Add new tour destinations
- Update tour destinations
- Delete tour destinations
- Add new airlines
- Update airlines
- Delete airlines
- Add new locations
- Update locations
- Delete locations
- Add new facilities
- Update facilities
- Delete facilities
- See the report of bookings and contents from dashboard 


## 🗺️ ERD
![ERD](https://raw.githubusercontent.com/Kelompok-2-Wanderer/Wanderer-BE/docs/docs/erd.png)

## ✔️ Requirement
Things you need to run the Wanderer API
1. **Cloudinary Account**
2. **Midtrans Account**

## 🧰 Installation
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

## 🤖 Author

- Heru Setiawan
  - [Github](https://github.com/heru-setiawan)

- Galih Prayoga
  - [Github](https://github.com/galihpra)