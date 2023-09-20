## P02 - Image Manipulator
### Bryce Koch
### Description:

This program utilizes the github.com/fogleman/gg package to load images into go. 
Using this package, the program takes the image mustangs.jpg and outlines the 
mustangs in the photo with a black line. 


### Files

|   #   | File             | Description                                        |
| :---: | ---------------- | -------------------------------------------------- |
|   1   |  [imagemod](https://github.com/BKoch74/4143-PLC/tree/main/Assignments/P02/imagemod) | Folder that contains the entire program. |
|   1   | [Main.go](https://github.com/BKoch74/4143-PLC/blob/main/Assignments/P02/imagemod/main.go)      | Main file for my program.      |
|   2   | [imageManipulator](https://github.com/BKoch74/4143-PLC/tree/main/Assignments/P02/imagemod/imageManipulator)  | File that contains image manipulating instructions.   |
|   3   | [imageManipulator.go](https://github.com/BKoch74/4143-PLC/blob/main/Assignments/P02/imagemod/imageManipulator/imageManipulator.go) | File that gets and manipulates the mustangs image. |
| 4 | [go.mod](https://github.com/BKoch74/4143-PLC/blob/main/Assignments/P02/imagemod/go.mod) | Initialized module file. |
| 5 | [mustangs.jpg](https://github.com/BKoch74/4143-PLC/blob/main/Assignments/P02/imagemod/mustangs.jpg) | Mustangs picture being manipulated. |
| 6 | [mustangs.png](https://github.com/BKoch74/4143-PLC/blob/main/Assignments/P02/imagemod/mustangs.png) | Mustangs picture after being manipulated. |


### Instructions

- The mustangs.jpg needs to be in the main directory.
- The github.com/fogleman/gg package must be installed using the command line go get.


- Example Command:
  - go run main.go
  - go get -u github.com/fogleman/gg
