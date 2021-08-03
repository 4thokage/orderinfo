import React, {useEffect, useState} from 'react';
import {StyleSheet,  View} from 'react-native';
import {OrderWatcherClient} from "./_proto/orderWatcher_pb_service";
import {Request, Response} from "./_proto/orderWatcher_pb";
import { Card, ListItem, Button, Icon } from 'react-native-elements'

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
            {messages.map(order =>
                <Card key={order.order}>
                    <Card.Title>{order.eta}</Card.Title>
                    <Card.Divider/>
                    {order.itemsList.map((u, i) => {
                        return (
                            <View key={i}>
                                <p>{u}</p>
                            </View>
                        );
                    })
                    }
                </Card>


            )}
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
