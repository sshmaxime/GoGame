import io from 'socket.io-client';

import { Dispatch } from "redux";
import { store } from "../index"
import { user, Error, message, room, RegisterRequest, state } from "../types/types"

import { LoginRequest, CreateRoomRequest, LeaveRoomRequest, JoinRoomRequest, MessageRoomRequest } from "../types/types"

import { LOGIN_SUCCESS, LEAVE_ROOM_SUCCESS, MESSAGE_ROOM, CREATE_ROOM_SUCCESS, JOIN_ROOM_SUCCESS, READY, UPDATE_STATE } from "../reducers/websocket.reducer";
import NotificationCenter from '../../global/notification';

const ENDPOINT = "https://gogame.sshsupreme.xyz";

const socket = io(ENDPOINT, {
  transports: ['websocket'],
  reconnection: false,
})

socket.on("connect", () => {

  socket.on("ERROR", (error: Error) => {
    NotificationCenter.getInstance().notificationErr(error)
  })
  socket.on("REGISTER_SUCCESS", (user: user) => {
    NotificationCenter.getInstance().notificationSuccess("User registered !")
  })

  socket.on(LOGIN_SUCCESS, (user: user) => {
    store.dispatch({ type: LOGIN_SUCCESS, payload: user });
  })

  socket.on(UPDATE_STATE, (state: state) => {
    console.log(state)
    store.dispatch({ type: UPDATE_STATE, payload: state });
  })

  socket.on(CREATE_ROOM_SUCCESS, (room: room) => {
    store.dispatch({ type: CREATE_ROOM_SUCCESS, payload: room });
    store.dispatch({ type: JOIN_ROOM_SUCCESS, payload: room });
  })

  socket.on(JOIN_ROOM_SUCCESS, (room: room) => {
    store.dispatch({ type: JOIN_ROOM_SUCCESS, payload: room });
  })

  socket.on(LEAVE_ROOM_SUCCESS, () => {
    store.dispatch({ type: LEAVE_ROOM_SUCCESS });
  })

  socket.on(MESSAGE_ROOM, (msg: message) => {
    store.dispatch({ type: MESSAGE_ROOM, payload: msg });
  })

  store.dispatch({ type: READY });

  // socket.emit("LOGIN_REQUEST", { username: "player1", password: "player1" })
  // socket.emit("JOIN_ROOM_REQUEST", { name: "demo" })
})


const Login = (username: string, password: string) => {
  const request: LoginRequest = { username: username, password: password }

  return async () => {
    socket.emit("LOGIN_REQUEST", request)
  };
}

const Register = (username: string, password: string) => {
  const request: RegisterRequest = { username: username, password: password }

  return async () => {
    socket.emit("REGISTER_REQUEST", request)
  };
}

const CreateRoom = (name: string) => {
  const request: CreateRoomRequest = { name: name }

  return async () => {
    socket.emit("CREATE_ROOM_REQUEST", request)
  };
}

const LeaveRoom = (name: string) => {
  const request: LeaveRoomRequest = { name: name }

  return async () => {
    socket.emit("LEAVE_ROOM_REQUEST", request)
  };
}

const JoinRoom = (name: string) => {
  const request: JoinRoomRequest = { name: name }

  return async () => {
    socket.emit("JOIN_ROOM_REQUEST", request)
  };
}

const MessageRoom = (roomName: string, msg: string) => {
  const request: MessageRoomRequest = { room_name: roomName, msg: msg }
  return async () => {
    socket.emit("MESSAGE_ROOM_REQUEST", request)
  };
}

export const websocketActions = {
  Login,
  Register,
  CreateRoom,
  JoinRoom,
  LeaveRoom,
  MessageRoom
};

