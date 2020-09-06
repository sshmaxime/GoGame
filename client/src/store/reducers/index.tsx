import { combineReducers, Reducer } from "redux";

import { commonState } from "./common.reducer";
import { websocketState } from "./websocket.reducer";

export interface IAppState {
  commonState: commonState;
  websocketState: websocketState;
}

export const reducers: Reducer<IAppState> = combineReducers<IAppState>({
  commonState,
  websocketState,
});
