import { combineReducers, Reducer } from "redux";

import { commonState } from "./common.reducer";

export interface IAppState {
  commonState: commonState;
}

export const reducers: Reducer<IAppState> = combineReducers<IAppState>({
  commonState
});
