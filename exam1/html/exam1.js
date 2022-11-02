let ws = new WebSocket("ws://localhost:8080/calcul"); // Création du socket
let input = document.getElementById("input");
ws.onopen = function(evt) {
    console.log("OPEN");
}
ws.onmessage = function(evt) {
    console.log("RESPONSE: " + evt.data); // Lorsqu'on reçoit un message
    input.value = evt.data;
}


document.getElementById("calcul").onclick = function(evt) {
    if (!ws) {
        return false;
    }
    console.log("SEND: " + input.value);
    ws.send(input.value); // Envoi d'un message
    return false;
};