the new new NEW new NEW japura website

short history of japura.net:
first version was joomla
second version was newer joomla
third version was octopress
forth version was nodejs/bootstrap
this version is go/react (LAST VERSION i swear)

go backend, react frontend

## getting started
- install go, and setup your gopath
- install nodejs
- install make
- check out the github project into $GOPATH/src/github.com/monofuel/webapp-template
- `make setup` to prepare the project initially
- `make watch -j2` to start running both the frontend and backend with hot-reload

## make targets
- `run` build and run the server
- `build` build everything
- `buildServer` build only the server
- `buildClient` build only the frontend
- `watch` runs the server with hot reload for backend and frontend (make sure to run with -j2)
- `watchServer` runs the server with hot reload
- `watchClient` builds the frontend with hot reload
- `setup` runs go get and npm install to set things up
- `test` check/test frontend and backend

## TODO
- add authors to articles
- add new screenshot to readme
- add separate blogs for users
- add RSS feed
- show minecraft server status
- add image upload when making posts
- add paging system on frontend (already setup for backend)
- seed initial 10 posts and their authors in the index.html request
