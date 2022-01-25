# CVWO todolist
 
Through this project I learnt the basics of React.js, Golang and also PostgreSQL. 
There were numerous hurdles along the way. Initially I wanted to implement an authentication system but found it to be more complex and intricate than I could handle in the time frame. I was also unable to complete the tagging and search function due to underestimating the time it would take. I will continue to work on it in the future. I also had a great learning experience developing a slightly larger project, having only wrote programs that existed in a single file or function. This experience taught me how important it is to plan out the components beforehand and to integrate tools that can solve the problem that you face.

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

