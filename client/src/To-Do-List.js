import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";
import ToggleButton from 'react-bootstrap/ToggleButton'


let endpoint = "http://localhost:8080";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: []
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    if (task) {
      axios
        .post(
          endpoint + "/api/task", "task=".concat(task),
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        )
        .then(res => {
          this.getTask();
          this.setState({
            task: ""
          });
          console.log(res);
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/task").then(res => {
      console.log(res);
      if (res.data) {
        console.log(res.data)

        this.setState({
          items: res.data.map(item => {
            return (
              <Card key={item.Id} fluid>
                <Card.Content>
                  <Card.Meta>
                    <ToggleButton
                      className="checkmark"
                      type="checkbox"
                      checked={item.Done}
                      style= {{float: 'left'}}
                      onChange={(e) => this.updateTask(item.Id)}
                    >
                    </ToggleButton>

                    <Icon
                      name="delete"
                      color="red"
                      style= {{float: 'right'}}
                      onClick={() => this.deleteTask(item.Id)}
                    />
                  </Card.Meta>

                  <Card.Header textAlign="left" left-margin="100px">
                    <div style={{ marginLeft: '30px', marginTop: '-20px', wordWrap: "break-word" }}>{item.Desc}</div>
                  </Card.Header>

                </Card.Content>
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
    });
  };

  updateTask = (id) => {
    axios
      .post(endpoint + "/api/toggleTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };

  deleteTask = id => {
    axios
      .delete(endpoint + "/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };
  render() {
    return (
      <div>
        <div className="row">
          <Header className="App-header">
            To Do List
          </Header>
        </div>
        <div className="rowcreate">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Insert new task here"
              autoComplete="off"
            />
          </Form>
        </div>

        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;
