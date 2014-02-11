Tardis
========

A simple windows executable that forces a user logout after a specified time.


Setup
=======
#### Setup environment
For Linux: ``export GOPATH=`pwd` `` <br />
For Windows: `set GOPATH=%CD%`


#### Fetch dependencies
Now run <br />
  `go get github.com/lxn/walk` <br />
  `go get github.com/akavel/rsrc`


Build
=======

#### Compile
In the directory containing `build.bat` and `build.sh` run

For Linux: ``./build.sh`` <br />
For Windows: `build.bat`

#### Run app
`tardis.exe`