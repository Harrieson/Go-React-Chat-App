import { Component } from "react";

class Message extends Component{
    constructor(props){
        super(props)
        let temp = JSON.parse(this.props.message)
        this.state={
            message: temp
        }
    }
    render(){
        return(
            <div className="Message">
                {this.state.Message.body}
            </div>
        )
    }
}

export default Message