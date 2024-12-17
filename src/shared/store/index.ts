import { observer } from "mobx-react";
import { userStore } from './UserStore';

const store = {
  userStore,
};

export * from "./UserStore";
export default store;
