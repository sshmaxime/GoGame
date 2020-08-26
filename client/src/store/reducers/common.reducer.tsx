// import commonConstants from "../constants/common.constants";
import { Action } from "../types/types"

type commonState = {
};

const commonState = (
  state: commonState = {
  },
  action: Action
): commonState => {
  switch (action.type) {
    default:
      return {
        ...state
      };
  }
};

export { commonState };
