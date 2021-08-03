import React, {useEffect, useState} from 'react';
import {StyleSheet, Text, View} from 'react-native';
import {OrderWatcherClient} from "./_proto/orderWatcher_pb_service";
import {Request, Response} from "./_proto/orderWatcher_pb";

export default function App() {
    const [messages, setMessages] = useState<Array<Response.AsObject>>([]);

    const host = "http://localhost:8080";
    const grpcClient = new OrderWatcherClient(host);


    useEffect(() => {
        (() => {
            const customer = new Request()
            customer.setId(1)
            const orderStream = grpcClient.subscribe(customer);
            orderStream.on("data", (chunk) => {
                const msg = chunk.toObject();
                setMessages((prev) => [...prev, msg]);
            });
        })();


    }, []);

    return (
        <View style={styles.container}>
            <Text>{messages.map(txt => <p key={txt.order}>{txt.itemsList[0]}</p>)}</Text>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
});
