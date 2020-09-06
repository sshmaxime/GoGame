import { user, room, message, state } from "../types/types"

export const READY = "READY"
export interface IREADY {
  type: typeof READY
}

export const LOGIN_SUCCESS = "LOGIN_SUCCESS"
export interface ILOGIN_SUCCESS {
  type: typeof LOGIN_SUCCESS
  payload: user
}

export const UPDATE_STATE = "UPDATE_STATE"
export interface IUPDATE_STATE {
  type: typeof UPDATE_STATE
  payload: state
}

export const CREATE_ROOM_SUCCESS = "CREATE_ROOM_SUCCESS"
export interface ICREATE_ROOM_SUCCESS {
  type: typeof CREATE_ROOM_SUCCESS
  payload: room
}

export const JOIN_ROOM_SUCCESS = "JOIN_ROOM_SUCCESS"
export interface IJOIN_ROOM_SUCCESS {
  type: typeof JOIN_ROOM_SUCCESS
  payload: room
}

export const LEAVE_ROOM_SUCCESS = "LEAVE_ROOM_SUCCESS"
export interface ILEAVE_ROOM_SUCCESS {
  type: typeof LEAVE_ROOM_SUCCESS
}

export const MESSAGE_ROOM = "MESSAGE_ROOM"
export interface IMESSAGE_ROOM {
  type: typeof MESSAGE_ROOM
  payload: message
}

type Actions = IREADY
  | ILOGIN_SUCCESS
  | ICREATE_ROOM_SUCCESS
  | IJOIN_ROOM_SUCCESS
  | IMESSAGE_ROOM
  | ILEAVE_ROOM_SUCCESS
  | IUPDATE_STATE

interface websocketState {
  ready: boolean
  user: user | undefined;
  room: room | undefined;
  messages: message[];
  online_users: user[];
  online_rooms: room[];
};

const websocketState = (
  state: websocketState = {
    ready: false,
    user: undefined,
    room: undefined,
    messages: [],
    online_users: [],
    online_rooms: []
  },
  action: Actions
): websocketState => {
  switch (action.type) {
    case READY:
      return {
        ...state, ready: true
      }
    case UPDATE_STATE:
      return {
        ...state,
        online_rooms: action.payload.online_rooms || [],
        online_users: action.payload.online_users || []
      }
    case LOGIN_SUCCESS:
      return {
        ...state, user: action.payload
      }
    case CREATE_ROOM_SUCCESS:
      return {
        ...state
      }
    case JOIN_ROOM_SUCCESS:
      return {
        ...state, room: action.payload
      }
    case LEAVE_ROOM_SUCCESS:
      return {
        ...state, room: undefined, messages: []
      }
    case MESSAGE_ROOM:
      console.log(action.payload)
      return {
        ...state, messages: [...state.messages, action.payload]
      }
    default:
      return {
        ...state
      };
  }
};

export { websocketState };
