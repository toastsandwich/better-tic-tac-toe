import { applyMiddleware, createStore } from "redux";
import UserReducer, { initialState } from "./reducers";
import { thunk } from "redux-thunk";

const persistedUser = JSON.parse(localStorage.getItem("user"));
const isLoggedIn = !!persistedUser && persistedUser.token;
const initialStoreState = {
  ...initialState,
  user: persistedUser || initialState.user,
  status: isLoggedIn,
};
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
