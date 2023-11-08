export type FieldType = "title" | "paragraph";

export interface FieldSchema {
  type: FieldType;
  body: string;
}
