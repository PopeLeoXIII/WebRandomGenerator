let submitBut = document.getElementById("submitBut");
let resetBut = document.getElementById("resetBut");

let nInput = document.getElementById("nInput");
let maxInput = document.getElementById("maxInput");
let genNInput = document.getElementById("genNInput");

let ul = document.getElementById("ul");

function createWS(msg){
    let socket = new WebSocket("ws://localhost:8080/wsocket" + msg);
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!")
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

    socket.onmessage = msg => {
        console.log(msg);
        addEl("<strong>Новое сообщение!</strong> " + msg.data)
    }
    return socket
}


function gen() {
    ul.innerHTML = "";
    let msg = "?n=" + nInput.value + "&max=" + maxInput.value  + "&genn=" +  genNInput.value;
    createWS(msg);
}

function myreset() {
    document.getElementById("nInput").setAttribute("value", "") ;
    maxInput.setAttribute("value", "") ;
    genNInput.setAttribute("value", "") ;

    alert("Данные очищены");
}

submitBut.onclick = function(){
    gen();
}

resetBut.onclick = function(){
    myreset();
}

function addEl(msg) {
    ul.insertAdjacentHTML('beforeend',"<li>" + msg + "</li>");
}
