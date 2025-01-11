import { applyMiddleware, createStore } from "redux";
import UserReducer, { initialState } from "./reducers";
import { thunk } from "redux-thunk";

const persistedUser = JSON.parse(localStorage.getItem("user"));
const initialStoreState = persistedUser
  ? { ...initialState, user: persistedUser }
  : initialState;

const store = createStore(
  UserReducer,
  initialStoreState,
  applyMiddleware(thunk)
);

store.subscribe(() => {
  const state = store.getState();
  if (state.user) {
    localStorage.setItem("user", JSON.stringify(state.user));
  }
});

export default store;
