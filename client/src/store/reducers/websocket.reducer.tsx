import { act } from "react-dom/test-utils"
import { user, room, message, state, game } from "../types/types"

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

export const CREATE_GAME_SUCCESS = "CREATE_GAME_SUCCESS"
export interface ICREATE_GAME_SUCCESS {
  type: typeof CREATE_GAME_SUCCESS
}

export const JOIN_GAME_SUCCESS = "JOIN_GAME_SUCCESS"
export interface IJOIN_GAME_SUCCESS {
  type: typeof JOIN_GAME_SUCCESS
}

export const LEAVE_GAME_SUCCESS = "LEAVE_GAME_SUCCESS"
export interface ILEAVE_GAME_SUCCESS {
  type: typeof LEAVE_GAME_SUCCESS
}

export const GAME_STATE = "GAME_STATE"
export interface IGAME_STATE {
  type: typeof GAME_STATE
  payload: game
}

type Actions = IREADY
  | ILOGIN_SUCCESS
  | ICREATE_ROOM_SUCCESS
  | IJOIN_ROOM_SUCCESS
  | IMESSAGE_ROOM
  | ILEAVE_ROOM_SUCCESS
  | IUPDATE_STATE
  | IGAME_STATE
  | IJOIN_GAME_SUCCESS
  | ICREATE_GAME_SUCCESS
  | ILEAVE_GAME_SUCCESS

interface websocketState {
  ready: boolean;
  online_users: user[];
  online_rooms: room[];

  connected: boolean;
  user: user;

  inRoom: boolean;
  messages: message[];
  room: room;

  inGame: boolean;
  gameStarted: boolean;
  game: game;
};

const websocketState = (
  state: websocketState = {
    ready: false,
    online_users: [],
    online_rooms: [],

    connected: false,
    user: { username: "" },

    inRoom: false,
    messages: [],
    room: { name: "" },

    inGame: false,
    gameStarted: false,
    game: { board: [], who_to_play: [], victory: -1, players: new Map<string, number>() }
  },
  action: Actions
): websocketState => {
  console.log(action)
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
        ...state, connected: true, user: action.payload
      }
    case CREATE_ROOM_SUCCESS:
      return {
        ...state
      }
    case JOIN_ROOM_SUCCESS:
      return {
        ...state, inRoom: true, room: action.payload
      }
    case LEAVE_ROOM_SUCCESS:
      return {
        ...state, inRoom: false, inGame: false, gameStarted: false, messages: []
      }
    case MESSAGE_ROOM:
      return {
        ...state, messages: [...state.messages, action.payload]
      }

    case CREATE_GAME_SUCCESS:
      return {
        ...state
      }
    case JOIN_GAME_SUCCESS:
      return {
        ...state, inGame: true,
      }
    case LEAVE_GAME_SUCCESS:
      return {
        ...state, inGame: false, gameStarted: false, game: { board: [], who_to_play: [], victory: -1, players: new Map<string, number>() }
      }
    case GAME_STATE:
      action.payload.players = new Map(Object.entries(action.payload.players));
      return {
        ...state, gameStarted: true, game: action.payload
      }
    default:
      return {
        ...state
      };
  }
};

export { websocketState };
