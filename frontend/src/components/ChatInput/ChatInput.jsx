import React, {Component} from 'react'
import './ChatInput.scss'

class ChatInput extends Component {
    render(){
        return (
            <div className="chatInput">
                <input onKeyDown={this.props.send} placeholder="Type a message here ..." />
            </div>
        )
    }
}
export default ChatInput