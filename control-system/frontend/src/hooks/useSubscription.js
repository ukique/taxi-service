import { useEffect, useCallback } from "react";
import ws from "../services/websocket";

export function useSubscription({ subscribeMsg, unsubscribeMsg, onMessage }) {
    const handler = useCallback(
        (data) => { onMessage(data); },
        [onMessage]
    );

    useEffect(() => {
        if (!subscribeMsg) {
            ws.off("message", handler);
            return;
        }

        const sendSubscribe = () => ws.send(subscribeMsg);

        if (ws.socket?.readyState === WebSocket.OPEN) {
            sendSubscribe();
        } else {
            ws.on("open", sendSubscribe);
        }

        ws.on("message", handler);

        return () => {
            if (unsubscribeMsg) ws.send(unsubscribeMsg);
            ws.off("message", handler);
            ws.off("open", sendSubscribe);
        };
    }, [subscribeMsg, handler]);
}