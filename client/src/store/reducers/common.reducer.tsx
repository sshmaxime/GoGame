
interface commonState {

};

const commonState = (
  state: commonState = {

  },
  action: any
): commonState => {
  switch (action.type) {
    default:
      return {
        ...state
      };
  }
};

export { commonState };
