import { StyleSheet, TextInput, TextInputProps } from "react-native";

// type FormatingTypeValues = {
//   link: string; // url
//   comment: number; // commentId
// };

// interface FormatingTypeBase {
//   startIdx: number;
//   endIdx: number;
// }

// type FormatingTypes = {
//   [Key in keyof FormatingTypeValues]: FormatingTypeBase & {
//     type: Key;
//     value: FormatingTypeValues[Key];
//   };
// }[keyof FormatingTypeValues];

// type BaseVariantOptions = {
//   formatings: FormatingTypes[];
// };

type VariantsValues = {
  heading1: string;
  heading2: string;
  heading3: string;
  text: string;
  // bulletList: string[];
  // numberList: string[];
};

export type VariantTypes = keyof VariantsValues;

export type Variants = {
  [Key in keyof VariantsValues]: {
    type: Key;
    value: VariantsValues[Key];
  };
}[keyof VariantsValues];

interface EditableTextProps
  extends Omit<TextInputProps, "style" | "onChangeText"> {
  onChangeValue: (variant: Variants) => void;
}
const EditableText = ({ type, ...props }: EditableTextProps & Variants) => {
  return (
    <TextInput
      {...props}
      onChangeText={(value) => props.onChangeValue({ type, value })}
      style={editableStyle[type]}
    />
  );
};

const editableStyle = StyleSheet.create<Record<VariantTypes, any>>({
  heading1: {
    fontSize: 40,
    fontWeight: "600",
  },
  heading2: {
    fontSize: 32,
    fontWeight: "600",
  },
  heading3: {
    fontSize: 24,
    fontWeight: "600",
  },
  text: {
    fontSize: 16,
  },
});

export default EditableText;
