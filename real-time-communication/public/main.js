const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const send = document.querySelector('#send')

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

ws.onmessage = function (msg) {
    insertMessage(JSON.parse(msg.data))
};

send.onclick = () => {
    const message = {
		username: username.value,
		content: input.value,
	}

    ws.send(JSON.stringify(message));
    input.value = "";
};

/**
 * Insere a mensagem no UI
 * @param {Message that will be displayed in the UI} messageObj
 */
function insertMessage(messageObj) {
	// Cria um objeto div para a mensagem
	const message = document.createElement('div')

	// Seta o atributo div na mensagem
	message.setAttribute('class', 'chat-message')
	console.log("name: " +messageObj.username + " content: " + messageObj.content)
	message.textContent = `${messageObj.username}: ${messageObj.content}`

	// Faz o Append da mensagem na div do chat
	messages.appendChild(message)

	// Inseri como a primeira mensagem do chat
	messages.insertBefore(message, messages.firstChild)
}