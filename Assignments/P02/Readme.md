## P02 - Image Manipulator
### Bryce Koch
### Description:

This program utilizes the github.com/fogleman/gg package to load images into go. 
Using this package, the program takes the image mustangs.jpg and outlines the 
mustangs in the photo with a black line. 


### Files

|   #   | File             | Description                                        |
| :---: | ---------------- | -------------------------------------------------- |
|   1   | [Main.go](./main.go)      | Main file for my program.      |
|   2   | [imageManipulator](./imageManipulator/)  | File that contains image manipulating instructions.   |
|   3   | [imageManipulator.go](./imageManipulator.go) | File that gets and manipulates the mustangs image. |
| 4 | [go.mod](./go.mod) | Initialized module file. |
| 5 | [mustangs.jpg]() | Mustangs picture being manipulated. |
| 6 | [mustangs.png]() | Mustangs picture after being manipulated. |


### Instructions

- The mustangs.jpg needs to be in the main directory.
- The github.com/fogleman/gg package must be installed using the command line go get.


- Example Command:
  - go run main.go
  - go get -u github.com/fogleman/gg
