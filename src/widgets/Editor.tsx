import fields from "@/mocks/fields";
import { FieldSchema, FieldType } from "@/shared/types";
import EditableText from "@/shared/ui/EditableText";
import { store } from "@/store";
import { useMutation, useQuery } from "@tanstack/react-query";
import React, { useRef } from "react";
import {
  GestureResponderEvent,
  Keyboard,
  Pressable,
  StyleSheet,
  Text,
  View,
} from "react-native";

export interface EditorProps {}

const Editor = (props: EditorProps) => {
  const { data, isLoading, isError } = useQuery({
    queryKey: ["page"],
    queryFn: store.fetchPage,
  });
  const {mutate: updatePart} = useMutation({
    mutationFn: store.updatePart,
    mutationKey: ["updatePart"],
  });
  const lastField = useRef<FieldSchema>();

  if (isError) return <Text>Error</Text>;
  if (isLoading) return <Text>Loading...</Text>;

  return (
    <View>
      <Pressable style={s.wrapper}>
        {fields.map((el) => (
          <EditableText
            key={el.body}
            type={el.type}
            onChangeValue={() => {}}
          />
        ))}
      </Pressable>
    </View>
  );
};

const s = StyleSheet.create({
  wrapper: {},
});

export default Editor;
