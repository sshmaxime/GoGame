import io from "socket.io-client";

import { store } from "../index";
import {
  user,
  Error,
  message,
  room,
  RegisterRequest,
  state,
  CreateGameRequest,
  LeaveGameRequest,
  JoinGameRequest,
  game,
} from "../types/types";

import {
  LoginRequest,
  CreateRoomRequest,
  PlayGameRequest,
  LeaveRoomRequest,
  JoinRoomRequest,
  MessageRoomRequest,
} from "../types/types";

import {
  LOGIN_SUCCESS,
  LEAVE_ROOM_SUCCESS,
  MESSAGE_ROOM,
  CREATE_ROOM_SUCCESS,
  JOIN_ROOM_SUCCESS,
  READY,
  UPDATE_STATE,
  GAME_STATE,
  CREATE_GAME_SUCCESS,
  JOIN_GAME_SUCCESS,
} from "../reducers/websocket.reducer";
import NotificationCenter from "../../global/notification";

const ENDPOINT =
  process.env.NODE_ENV === "development"
    ? "http://localhost:3001"
    : "https://server.gogame.sshsupreme.xyz";

const socket = io(ENDPOINT, {
  transports: ["websocket"],
  reconnection: false,
});

socket.on("connect", () => {
  socket.on("ERROR", (error: Error) => {
    NotificationCenter.getInstance().notificationErr(error);
  });
  socket.on("REGISTER_SUCCESS", (user: user) => {
    NotificationCenter.getInstance().notificationSuccess("User registered !");
  });

  socket.on(LOGIN_SUCCESS, (user: user) => {
    store.dispatch({ type: LOGIN_SUCCESS, payload: user });
  });

  socket.on(UPDATE_STATE, (state: state) => {
    console.log(state);
    store.dispatch({ type: UPDATE_STATE, payload: state });
  });

  socket.on(CREATE_ROOM_SUCCESS, (room: room) => {
    store.dispatch({ type: CREATE_ROOM_SUCCESS, payload: room });
    store.dispatch({ type: JOIN_ROOM_SUCCESS, payload: room });
  });

  socket.on(JOIN_ROOM_SUCCESS, (room: room) => {
    store.dispatch({ type: JOIN_ROOM_SUCCESS, payload: room });
  });

  socket.on(LEAVE_ROOM_SUCCESS, () => {
    store.dispatch({ type: LEAVE_ROOM_SUCCESS });
  });

  socket.on(MESSAGE_ROOM, (msg: message) => {
    store.dispatch({ type: MESSAGE_ROOM, payload: msg });
  });

  socket.on(CREATE_GAME_SUCCESS, () => {
    store.dispatch({ type: CREATE_GAME_SUCCESS });
    store.dispatch({ type: JOIN_GAME_SUCCESS });
  });
  socket.on(JOIN_GAME_SUCCESS, () => {
    store.dispatch({ type: JOIN_GAME_SUCCESS });
  });
  socket.on(GAME_STATE, (msg: game) => {
    store.dispatch({ type: GAME_STATE, payload: msg });
  });

  store.dispatch({ type: READY });
});

const Login = (username: string, password: string) => {
  const request: LoginRequest = { username: username, password: password };

  return async () => {
    socket.emit("LOGIN_REQUEST", request);
  };
};

const Register = (username: string, password: string) => {
  const request: RegisterRequest = { username: username, password: password };

  return async () => {
    socket.emit("REGISTER_REQUEST", request);
  };
};

// Room
const CreateRoom = (name: string) => {
  const request: CreateRoomRequest = { name: name };

  return async () => {
    socket.emit("CREATE_ROOM_REQUEST", request);
  };
};
const LeaveRoom = (name: string) => {
  const request: LeaveRoomRequest = { name: name };

  return async () => {
    socket.emit("LEAVE_ROOM_REQUEST", request);
  };
};
const JoinRoom = (name: string) => {
  const request: JoinRoomRequest = { name: name };

  return async () => {
    socket.emit("JOIN_ROOM_REQUEST", request);
  };
};
const MessageRoom = (roomName: string, msg: string) => {
  const request: MessageRoomRequest = { room_name: roomName, msg: msg };
  return async () => {
    socket.emit("MESSAGE_ROOM_REQUEST", request);
  };
};

// Game
const CreateGame = (roomName: string, gameName: string) => {
  const request: CreateGameRequest = { room_name: roomName, name: gameName };

  return async () => {
    socket.emit("CREATE_GAME_REQUEST", request);
  };
};
const LeaveGame = (roomName: string, gameId: string) => {
  const request: LeaveGameRequest = { room_name: roomName };

  return async () => {
    socket.emit("LEAVE_GAME_REQUEST", request);
  };
};
const JoinGame = (roomName: string) => {
  const request: JoinGameRequest = { room_name: roomName };

  return async () => {
    socket.emit("JOIN_GAME_REQUEST", request);
  };
};
const PlayGame = (roomName: string, data: Object) => {
  const request: PlayGameRequest = {
    room_name: roomName,
    data: JSON.stringify(data),
  };
  return async () => {
    socket.emit("PLAY_GAME_REQUEST", request);
  };
};

export const websocketActions = {
  Login,
  Register,
  CreateRoom,
  JoinRoom,
  LeaveRoom,
  MessageRoom,
  CreateGame,
  LeaveGame,
  JoinGame,
  PlayGame,
};
