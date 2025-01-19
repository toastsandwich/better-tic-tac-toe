import net from "net";

class Client {
  constructor(host, port) {
    this.host = host;
    this.port = port;
    this.socket = new net.Socket();
  }

  connect() {
    this.socket.connect(this.port, this.host, () => {
      console.log(`connected to server ${this.host}:${this.port}`);
    });

    this.socket.on("data", (data) => {
      console.log(`server: ${data.toString()}`);
    });

    this.socket.on("error", (err) => {
      console.error(`error: ${err.message}`);
    });

    this.socket.on("close", () => {
      console.log("connection closed.");
      this.socket.destroy(); // Cleanup
    });
  }

  sendMove(x, y) {
    if (!this.socket.connecting && !this.socket.destroyed) {
      this.socket.write(`${x},${y}`);
      console.log(`Sent: ${x},${y}`);
    } else {
      console.log("Cannot send message, socket is not connected.");
    }
  }

  disconnect() {
    if (!this.socket.destroyed) {
      this.socket.end(() => console.log("disconnected from server"));
    } else {
      console.log("Socket already closed.");
    }
  }

  reconnect() {
    console.log("Reconnecting...");
    setTimeout(() => this.connect(), 1000); // Retry after 1 second
  }
}

export default Client;
