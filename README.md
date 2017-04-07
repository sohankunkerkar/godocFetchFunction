# GO program to extract fucntions from godoc package
Install latest golang version into your systems(go 1.8) and set the approriate $GOPATH in order to run the golang program properly.
Create one folder and keep all the files mentioned above including godeps folder to install package dependencies.<br>
For some reason it doesn't work then use following command to install the packages:

    go get -u <package name>
          
On Machine use following command to run the package and extract the functions:

    go run <file name> <package name> 
for e.g. if we want to extract functions from <b>OS</b> package then use following commad:

    go run printFunction.go os
       
for third party package use entire path:

    go run printFunction.go golang.org/x/tools/go/ast/astutil
