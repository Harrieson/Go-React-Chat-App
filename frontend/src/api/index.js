var socket = new Webscoket("ws://localhost:9000/ws")
let connect = (cb) => {
    console.log("Connecting ...")

    socket.onopen = () => {
        console.log("Connected Successfully")
    }

    socket.onmessage = (msg) => {
        console.log("Message from websocket:", msg)
    }

    socket.onclose = (event) => {
        console.log("socket closed connection", event)
    }
    socket.onerror = (error) => {
        console.log("Socket encountered an error", error)
    }
}

let sendMsg = (msg) => {
    console.log("Sending message", msg)
    socket.send(msg)
}

export { connect, sendMsg }