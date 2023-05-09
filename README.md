<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- [![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url] -->
[![LinkedIn][linkedin-shield]][linkedin-url]
[![Gmail][gmail-shield]][gmail-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/luongquochai/ecommerce_golang">
    <img src="logo/logo.jpg" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Ecommerce Project - Golang</h3>

  <p align="center">
    Description
    <br />
    <a href="https://github.com/luongquochai/ecommerce_golang"><strong>Explore the docs »</strong></a>
    <br />
    <!-- <br />
    <a href="https://github.com/github_username/repo_name">View Demo</a>
    ·
    <a href="https://github.com/github_username/repo_name/issues">Report Bug</a>
    ·
    <a href="https://github.com/github_username/repo_name/issues">Request Feature</a> -->
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
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
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <!-- <li><a href="#roadmap">Roadmap</a></li> -->
    <li><a href="#contributing">Contributing</a></li>
    <!-- <li><a href="#license">License</a></li> -->
    <li><a href="#contact">Contact</a></li>
    <!-- <li><a href="#acknowledgments">Acknowledgments</a></li> -->
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://github.com/luongquochai/ecommerce_golang)

* Diagrams

[![Diagram Screen Shot][diagram-screenshot]](https://github.com/luongquochai/ecommerce_golang/diagram.drawio)

This is a ecommerce project that helps me practice and improve backend skills. Other than that, I can learn more about using Golang and frameworks.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

<!-- * [![Golang][Golang]][Go-url] -->
<!-- * [![React][React.js]][React-url]
* [![Vue][Vue.js]][Vue-url]
* [![Angular][Angular.io]][Angular-url]
* [![Svelte][Svelte.dev]][Svelte-url]
* [![Laravel][Laravel.com]][Laravel-url]
* [![Bootstrap][Bootstrap.com]][Bootstrap-url]
* [![JQuery][JQuery.com]][JQuery-url] -->
<div align="center">
  <a href="https://github.com/luongquochai/ecommerce_golang">
    <img src="logo/background.png" alt="Logo">
  </a>
</div>
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started
* **System**:
```
> NAME="Ubuntu"
> VERSION="20.04.6 LTS (Focal Fossa)"
> ID=ubuntu
> ID_LIKE=debian
> PRETTY_NAME="Ubuntu 20.04.6 LTS"
> VERSION_ID="20.04"
```
* **Structure**:
```
├── controllers
│   ├── address.go
│   ├── cart.go
│   └── controllers.go
├── database
│   ├── cart.go
│   └── dbsetup.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
├── middleware
│   └── middleware.go
├── models
│   └── models.go
├── routes
│   └── routes.go
└── tokens
    └── tokengen.go
```
### Prerequisites

* Golang, Docker, MongoDB

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/luongquochai/ecommerce_golang.git
   ```
2. Install go packages
   ```sh
   go mod tidy
   ```
3. Start mongo server
   ```
   docker-compose up -d
   ```
4. Start project
    ```
    go run main.go
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

* SIGNUP FUNCTION API CALL (POST REQUEST)

> http://localhost:8000/users/signup

```
{
    "first_name": "Hai",
    "last_name": "Luong",
    "email": "luongquochai@gmail.com",
    "password": "hailuong",
    "phone": "0123456789"
}
```

**Response** : "Successfully Signed Up!!"

* LOGIN FUNCTION API CALL (POST REQUEST)

> http://localhost:8000/users/login

```
{
    "email": "luongquochai@gmail.com",
    "password": "hailuong"
}
```

**Response will be like this**
```
{
    "_id": "6457631b6c87663cfc793657",
    "first_name": "Hai",
    "last_name": "Luong",
    "password": "$2a$14$Z3ZYr22QAlYKSGb2d90BYeZZDCM7rSKMFN1OkBWAQoIPtcKW2RpDC",
    "email": "luongquochai@gmail.com",
    "phone": "01234456789",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Imx1b25ncXVvY2hhaTA3MDdAZ21haWwuY29tIiwiRmlyc3RfTmFtZSI6IkhhaSIsIkxhc3RfTmFtZSI6Ikx1b25nIiwiVWlkIjoiNjQ1NzYzMWI2Yzg3NjYzY2ZjNzkzNjU3IiwiZXhwIjoxNjgzNTM1MDAzfQ.HP4eMyXAPI9x_Fv51QPClVx3wZkiR9t9xTZ7nPYOR-w",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE2ODQwNTM0MDN9.9VDAPZNppmelT_zKULP5or0KFlEXf777RHT3s0jMh7Y",
    "created_at": "2023-05-07T08:36:43Z",
    "updated_at": "2023-05-07T08:36:43Z",
    "user_id": "6457631b6c87663cfc793657",
    "usercart": [],
    "address": [],
    "orders": []
}
```

> ******
> You need to add ***id*** & ***userID*** into Params tab and ***token*** in Headers tab.

* Admin add Product Function (POST REQUEST)

> http://localhost:8000/admin/addproduct

```
{
  "product_name": "Alienware x15",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg"
}
```

> Response : "Successfully added our Product Admin!!"

* View all the Products in db GET REQUEST

> View all the Products in db GET REQUEST

```Respone
[
  {
    "Product_ID": "6153ff8edef2c3c0a02ae39a",
    "product_name": "alienwarex15",
    "price": 1500,
    "rating": 10,
    "image": "alienware.jpg"
  },
  {
    "Product_ID": "616152679f29be942bd9df8f",
    "product_name": "giner ale",
    "price": 900,
    "rating": 5,
    "image": "gin.jpg"
  },
  {
    "Product_ID": "616152ee9f29be942bd9df90",
    "product_name": "iphone 13",
    "price": 1700,
    "rating": 4,
    "image": "ipho.jpg"
  },
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "whiskey",
    "price": 100,
    "rating": 7,
    "image": "whis.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "acer predator",
    "price": 3000,
    "rating": 10,
    "image": "acer.jpg"
  }
]
```

* Search Product by regex function (GET REQUEST)

> defines the word search sorting http://localhost:8000/users/search?name=al

```
[
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "Alienware x15",
    "price": 1500,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "ginger Ale",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  }
]
```

* Adding the Products to the Cart (GET REQUEST)

> http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

Corresponding mongodb query: "Successfully added to the cart"

* Removing Item From the Cart (GET REQUEST)

> Removing Item From the Cart (GET REQUEST)

>   http://localhost:8000/addtocart?id=xxxxxxx&userID=xxxxxxxxxxxx

Respone: "Successfully removed item from the cart"

* Listing the item in the users cart (GET REQUEST) and total price

> http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx

* Addding the Address (POST REQUEST)

> http://localhost:8000/addadress?id=user_id**\*\***\***\*\***

The Address array is limited to two values home and work address more than two address is not acceptable

```
{
  "house_name": "white house",
  "street_name": "white street",
  "city_name": "washington",
  "pin_code": "332423432"
}
```

* Editing the Home Address(PUT REQUEST)

> http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

* Editing the Work Address(PUT REQUEST)

> http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

* Delete Addresses(GET REQUEST)

> http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx

delete both addresses

* Cart Checkout Function and placing the order(GET REQUEST)

After placing the order the items have to be deleted from cart functonality added

> http://localhost:8000?id=xxuser_idxxx

* Instantly Buying the Products(GET REQUEST)

>http://localhost:8000?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx


_For more examples, please refer to the [Documentation](https://github.com/luongquochai/ecommerce_golang)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
<!-- ## Roadmap

- [ ] Feature 1
- [ ] Feature 2
- [ ] Feature 3
    - [ ] Nested Feature

See the [open issues](https://github.com/github_username/repo_name/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- CONTRIBUTING -->
## Contributing

Open source: https://github.com/AkhilSharma90/go-ecommerce-project
This is an amazing place to learn, inspire and create.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
<!-- ## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- CONTACT -->
## Contact

Luong Quoc Hai - [@Linkedin](https://www.linkedin.com/in/luongquochai/) - luongquochai0707@gmail.com

Project Link: [https://github.com/luongquochai/ecommerce_golang](https://github.com/luongquochai/ecommerce_golang)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
<!-- ## Acknowledgments

* []()
* []()
* []()

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/luongquochai/

[gmail-shield]: https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white
[gmail-url]: luongquochai0707@gmail.com

[product-screenshot]: logo/Capture.PNG
[diagram-screenshot]: logo/diagrams.PNG
[Golang]: https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png
[Go-url]: https://go.dev/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 