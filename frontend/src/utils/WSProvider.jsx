import { useEffect, useRef, useState } from 'react';
import { WSContext } from './WSContext';

export function WSProvider({ children }) {
    const wsRef = useRef(null);
    const [isConnected, setIsConnected] = useState(false);
    const callbacksRef = useRef([]);

    useEffect(() => {
        const websocket = new WebSocket('ws://localhost:8080');
        wsRef.current = websocket;

        websocket.onopen = () => setIsConnected(true);
        websocket.onclose = () => setIsConnected(false);
        websocket.onmessage = (e) => {
            const data = JSON.parse(e.data);
            callbacksRef.current.forEach(cb => cb(data));
        };

        return () => websocket.close();
    }, []);

    const sendMessage = (data) => {
        wsRef.current?.send(JSON.stringify(data));
    };

    const onMessage = (cb) => {
        callbacksRef.current.push(cb);
    };

    return (
        <WSContext.Provider value={{ sendMessage, onMessage, isConnected }}>
            {children}
        </WSContext.Provider>
    );
}