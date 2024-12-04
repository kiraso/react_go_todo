import React,{Component} from 'react'
import axios from 'axios'
import {Card,Header,Form,Input,Icon,Button} from 'semantic-ui-react'

let endpoint = "http://localhost:9000"

//Old react write with class base
class ToDoList extends Component{
    constructor(props){
        super(props)

        this.state = {
            task:"",
            items:[]
        }
    }

    componentDidMount(){
        this.getTask();
        
    }

    onChange = (e) => {
        this.setState({
            [e.target.name]: e.target.value
        })
    }
    onSubmit = () =>{
        let {task }  =  this.state;

        if(task) {
            axios.post(endpoint+"/api/task",{task},{
                headers:{
                    "Content-Type": "application/x-www-form-urlencoded",
                }
            }).then((res)=>{
                this.getTask();
                this.setState({
                    task:"",
                })

                console.log(res)
            })
        }
    }

    getTask = () => {
        axios.get(endpoint + '/api/task').then((res)=>
        {
            if(res.data){
                this.setState({
                    items: res.data.map((item) => {
                        let color = "yellow";
                        let style = {
                            wordWrap: "break-word"
                        }

                        if(item.status){
                            color = "green";
                            style["textDecoration"] = "line-through"
                        }

                        return (
                            <Card key={item._id} color={color} fluid className="rough">
                                <Card.Content>
                                    
                                    <Card.Header textAlign="left">
                                        <div style={style}>{item.task}</div>
                                    </Card.Header>

                                    <Card.Meta textAligh="right">
                                        <Icon
                                        name="check circle"
                                        color="blue"
                                        onClick={()=>this.updateTask(item.id)}
                                        />
                                        <span style={{paddingRight:"10px"}}>Undo</span>

                                        <Icon
                                        name="delete"
                                        color="red"
                                        onClick={()=> this.deleteTask(item.id)}
                                        />
                                        <span style={{paddingRight:"10px"}}>Delete</span>
                                    </Card.Meta>
                                </Card.Content>
                            </Card>
                        )
                    })
                })
            }else {
                this.setState({items: []})
            }
        })
    }

    updateTask = (id) => {
        axios.put(endpoint + "/api/tasks" + id,{
            headers:{
                "Content-Type": "application/x-www-form-urlencoded",

            }
        }).then((res)=>{
            console.log("res",res)
            this.getTask()
        })
    }

    undoTask = (id) => {
        axios.put(endpoint + "/api/undoTask" + id,{
            headers:{
                "Content-Type": "application/x-www-form-urlencoded",
            }
        }).then((res)=>{
            console.log("res",res)
            this.getTask()
        })
    }

    deleteTask = (id) => {
        axios.delete(endpoint + "/api/deleteTask" + id,{
            headers:{
                "Content-Type": "application/x-www-form-urlencoded",
            }
        }).then((res)=>{
            console.log("res",res)
            this.getTask()
        })
    }
    render(){
        
        return(
            <div>
                <div className="row">
                    <Header className="header" as="h1">To Do List</Header>
                </div>
                <div className='row'>
                    <Form onSubmit={this.onSubmit}>
                        <Input
                        type="text"
                        name="task"
                        onChange={this.onChange}
                        value={this.state.task}
                        fluid
                        placeholder="Create Task"
                        />
                        <Button>Create task</Button>
                    </Form>
                </div>
                <div className="row">
                    <Card.Group>{this.state.items}
                        {/* {this.state.items.map(item => (
                            <Card key={item._id}> */}
                                {/* <Card.Content>
                                    <Card.Header>{item.task}</Card.Header>
                                </Card.Content> */}
                            {/* </Card>
                        ))} */}
                    </Card.Group>
                </div>
            </div>
        )
    } 
}

export default ToDoList


