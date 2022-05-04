## Installation
 Clone the repository
````
git clone https://github.com/hsndmr/go-sanctum
````
Go to the repository folder and execute command
````
docker build --tag go-sanctum .
````

Run the go-sanctum docker image
````
docker run -p 3000:3000 --name go-sanctum go-sanctum
````