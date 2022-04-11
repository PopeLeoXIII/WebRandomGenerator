let submitBut = document.getElementById("submitBut");
let resetBut = document.getElementById("resetBut");

let nInput = document.getElementById("nInput");
let maxInput = document.getElementById("maxInput");
let genNInput = document.getElementById("genNInput");

let ul = document.getElementById("ul");

function createWS(){
    let socket = new WebSocket("ws://localhost:8080/wsocket");
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
    createWS();
    let msg = "?n=" + nInput.value + "&max=" + maxInput.value  + "&genn=" +  genNInput.value;
    ul.innerHTML = "";
    let xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://localhost:8080/ws'+msg, true);
    xhr.send();
}

function myreset() {
    document.getElementById("nInput").value = "";
    document.getElementById("maxInput").value = "";
    document.getElementById("genNInput").value = "";

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


function addElHard(msg) {
    let div = document.createElement('div');
    div.className = "alert";
    div.innerHTML = msg;

    document.body.append(div);
}