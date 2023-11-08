import { FieldSchema, FieldType } from "@/shared/types";
import EditableText from "@/shared/ui/EditableText";
import Editor from "@/widgets/Editor";
import { observer } from "mobx-react-lite";
import { useRef, useState } from "react";
import {
  GestureResponderEvent,
  Keyboard,
  Pressable,
  StyleSheet,
} from "react-native";

export interface HomePageProps {}

const HomePage = observer((props: HomePageProps) => {
  const [fields, setFields] = useState<FieldSchema[]>([
    { type: "title", body: "Untilted" },
  ]);
  const lastField = useRef<FieldSchema>();

  const onPressOutside = (e: GestureResponderEvent) => {
    if (lastField.current) {
      const { type, body } = lastField.current;
      const newFields = [...fields];
      const targetIndex = newFields.findIndex((el) => el.type === type);
      if (targetIndex === -1) {
        newFields.push({ type, body });
      } else {
        newFields[targetIndex].body = body;
      }
      setFields(newFields);
    }

    Keyboard.dismiss();
  };
  const onEditField = (type: FieldType, body: string) => {
    lastField.current = { type, body };
  };

  return (
		<Editor />
  );
});

const style = StyleSheet.create({
  wrapper: {
    flex: 1,
    padding: 40,
  },
});

export default HomePage;
