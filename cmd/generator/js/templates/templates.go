package templates

var Index_autojoin = `
const MatrixClient = require("matrix-bot-sdk").MatrixClient;
const AutojoinRoomsMixin = require("matrix-bot-sdk").AutojoinRoomsMixin;

const client = new MatrixClient("{{homeserver}}", "{{access_token}}");
AutojoinRoomsMixin.setupOnClient(client);

client.start().then(() => console.log("Client started!"));
`
