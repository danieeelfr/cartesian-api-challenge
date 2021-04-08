CARTESIAN API - Tech challenge - 2021, April 
##### https://github.com/tobym/cartesian
==========================================
## Challenge Description
Create an **API server in go**. It will deal with a series of points represented as (x,y) coordinates on a simple 2-dimensional plane.
![map](https://www.101computing.net/wp/wp-content/uploads/taxicab-manhattan-grid-layout-paths.gif)

*The Manhattan distance is measured "block-wise", as the distance in blocks between any two points in the plane (e.g. 2 blocks down and 3 blocks over for a total of 5 blocks).* It is defined as the sum of the horizontal and vertical distances between points on a grid. 

Formally, where p1 = (x1, y1) and p2 = (x2, y2), distance(p1,p2) = |x1-x2| + |y1-y2|.

Take a look 
[here](http://en.wikipedia.org/wiki/Cartesian_coordinate_system)  if you need a refresher on this concept.

## Challenge Requirements
- It must have an api route at **/api/points** that accepts a **GET** request with the following parameters 
- returns a JSON list of points that are within distance from x,y, **using the Manhattan distance method.** 
- The points should be returned in **order of increasing distance from the search origin**.
- **On startup**, the API server should read a list of points from data/points.json.

### Input Params
#### **x integer** (required). 
This represents the x coordinate of the search origin.
#### **y integer** (required). 
This represents the y coordinate of the search origin.
#### **distance integer** (required). 
This represents the Manhattan distance; points within distance from x and y are returned, points outside are filtered out.

__________
# Solution
### Prerequisites
What things you need to install the software and how to install them
#### Requires:
- Golang >= 1.15

### Development environment
#### Setting environment variables
```bash
. ./env.sh;
```

#### Building
```bash
go build -o ./bin/cartesianapi ./cmd/cartesianapi
```
#### Running the source code tests
```bash
go test -v -cover -covermode=atomic ./...
```
#### Running the application
```bash
./bin/cartesianapi
```

#### Directory
The directory structure follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
```
.
├── cmd
│   └── main.go
├── internal
│   ├── config
│   │   ├── config.go
│   │   └── domain.go
│   ├── controller
│   │   ├── httpserver
│   │   │   ├── handler.go
│   │   │   ├── router.go
│   │   │   ├── interactor.go
│   │   │   └── server.go
│   │   ├── points
│   │   │   ├── points.go
│   │   │   └── interactor.go
│   ├── entity
│   │   ├── errors.go
│   │   ├── distance.go
│   │   └── distance_test.go
│   └── usecase
│       └── distance_calculator
│           ├── domain.go
│           ├── domain_test.go
│           ├── interactor.go
│           ├── distance.go
│           ├── distance_test.go
```
