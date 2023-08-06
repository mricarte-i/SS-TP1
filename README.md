# SS - TP1: Cell index method
#### (or how to make a quadtree simpler by keeping it all in the root node and actually not use a tree but a map)

## How to run:
- ```go mod tidy```
- ```go run main.go 100 2 4 1```

- terminal will print particles' positions, and then a list of the neighbors' ids per each particle
- get those lists in two separate .json files
- run the web app (live server in vscode?)
- set the name numbers, upload the first array json w/ the UPLOAD INITIAL CONDITIONS button
- then upload the second json file w/ the UPLOAD NEIGHBORS DATA button
- use the select dropdown to pick which particle to get to know about (paint in red, paint neighbors in green)
