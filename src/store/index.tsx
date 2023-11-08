import fields from "@/mocks/fields";
import { FieldSchema } from "@/shared/types";
import { makeAutoObservable } from "mobx";
import { PropsWithChildren, createContext, useContext } from "react";

export interface AppState {}

class AppStore implements AppState {
  constructor() {
    makeAutoObservable(this);
  }

  setOne = <K extends keyof this>(key: K, value: this[K]) => {
    this[key] = value;
  };

  fetchPage = () => {
    return fields;
  };

  updatePart = (part: FieldSchema & { id: number }) => {
    return new Promise((res) => setTimeout(res, 2000));
  };
}

export const store = new AppStore();

const StoreContext = createContext<AppStore>(store);

export const StoreProvider = ({ children }: PropsWithChildren) => {
  return (
    <StoreContext.Provider value={store}>{children}</StoreContext.Provider>
  );
};

export const useAppStore = () => useContext(StoreContext);
