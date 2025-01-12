const initialState = {
  user: {
    country: "",
    username: "",
    global_rank: "",
    country_rank: "",
    wins: "",
    losses: "",
  },
};

const UserReducer = (state = initialState, action) => {
  switch (action.type) {
    case "SET_USER":
      return {
        ...state,
        user: { ...state.user, ...action.payload },
      };
    case "RM_USER":
      return initialState;
    default:
      return state;
  }
};

const SetUser = (user) => ({
  type: "SET_USER",
  payload: user,
});

const RemoveUser = () => ({
  type: "RM_USER",
});

export default UserReducer;
export { SetUser, RemoveUser, initialState };
