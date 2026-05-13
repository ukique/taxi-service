class WebSocketService {
    constructor() {
        this.socket = null;
        this.listeners = {};
    }

    connect(url) {
        if (this.socket?.readyState === WebSocket.OPEN) return;

        this.socket = new WebSocket(url);
        this.socket.onopen    = () => this.emit("open");
        this.socket.onmessage = (e) => this.emit("message", JSON.parse(e.data));
        this.socket.onclose   = () => this.emit("close");
        this.socket.onerror   = (err) => this.emit("error", err);
    }

    on(event, cb)    { (this.listeners[event] ??= []).push(cb); }
    off(event, cb)   { this.listeners[event] = this.listeners[event]?.filter(fn => fn !== cb); }
    emit(event, data){ this.listeners[event]?.forEach(cb => cb(data)); }

    send(data) {
        if (this.socket?.readyState === WebSocket.OPEN)
            this.socket.send(JSON.stringify(data));
    }

    disconnect() {
        this.socket?.close();
        this.socket = null;
    }
}

export default new WebSocketService();