const gk = new Gusk(window.location.host + '/chat_ws', false);
gk.Event('sms-del-servidor', (dato) => {
    addText(dato);
});


function sendTexto() {
    addText(`Yo: ${texto.value}`);
    gk.Send('sms-del-cliente', texto.value);
    texto.value = '';
}

function addText(text) {
    mensajes.innerHTML += `<p>${text}</p>`;
}

gk.Connect();