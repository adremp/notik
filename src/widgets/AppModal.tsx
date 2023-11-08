import React from "react";
import { StyleSheet, Text, View } from "react-native";

export interface ModalProps {}

const AppModal = (props: ModalProps) => {
	console.log('modal print :>> ');
  return (
    <View style={style.wrapper}>
      <Text>Modal</Text>
    </View>
  );
};

const style = StyleSheet.create({
  wrapper: {
    width: "80%",
    height: "100%",
    position: "absolute",
    left: 0,
    top: 0,
    bottom: 0,
    backgroundColor: "#fff",
  },
});

export default AppModal;
