### WEbsockets

Adding basic echo server websockets was successfull but cannot send echo messages to all the websocket clients. This way of websockets did only send echo message to client which sent the message in the first place. 

To check the websocket echo server working we do:

```js
    const socket = new WebSocket("ws://localhost:5000/handle-websocket-echo")
    socket.onmessage = console.log
    socket.send("hello")
```

And see if the sent message is logged to the console.