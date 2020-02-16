import React from 'react';
import {Alert, StyleSheet, Text, View} from 'react-native';
import {PingClient} from "./src/apis/proto/ping/PingServiceClientPb";
import {PingRequest} from "./src/apis/proto/ping/ping_pb";

export default function App() {
  const client = new PingClient("http://localhost:8080");
  const req = new PingRequest();
  req.setMessage("hello sueken");
  client.ping(req, null, (err, response) => {
    return (Alert.alert(response.getMessage()))
  });
  return (
    <View style={styles.container}>
      <Text>Hello world</Text>
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
