import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { TouchableHighlight } from "react-native-gesture-handler";
import BurgerIcon from "../shared/assets/icons/burger.svg";
import text from "../shared/theme/text";
import { useAppStore } from "../store";

export interface NavbarProps {}

const Navbar = (props: NavbarProps) => {
  const store = useAppStore();

  return (
    <View style={style.wrapper}>
      <TouchableHighlight
        onPress={() => store.setOne("modal", "settings")}
        style={style.burger}
      >
        <BurgerIcon />
      </TouchableHighlight>
      <Text style={text["20-500"]}>Navbar</Text>
    </View>
  );
};

const style = StyleSheet.create({
  burger: {
    width: 24,
    aspectRatio: 1,
  },
  wrapper: {
    height: 50,
    flexDirection: "row",
    justifyContent: "space-between",
    paddingHorizontal: 20,
    alignItems: "center",
    backgroundColor: "#dbdbdbb3",
  },
});

export default Navbar;
