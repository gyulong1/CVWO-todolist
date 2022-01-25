# CVWO todolist
 
To run the project locally, first set up the postgre database with a table login.tasks which contains columns desc (text), done (bool), id (serial). 

Then, navigate to the server directory and run 
```
go build 
```
followed by 
```
./server 
```
to start the backend server. Then cd to the client directory and run the react server
```
npm install && npm start
```
The application should open on your browser (http://localhost:3000).

To add a task, type in the task name and hit enter.
To delete a task, click the delete sign.
To check/uncheck a task, click the check mark.

References
https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6
https://reactjs.org/docs/components-and-props.html

Guo Yulong, A0180089Y, NUS
