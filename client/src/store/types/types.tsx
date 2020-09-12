export interface user {
  username: string
}
export interface room {
  name: string
}
export interface game {
  id: string,
}
export interface message {
  from: string,
  msg: string,
  time: string,
}
export interface state {
  online_users: user[],
  online_rooms: room[]
}

// ERROR
export type Error = string

// REQUEST
export interface LoginRequest {
  username: string
  password: string
}
export interface RegisterRequest {
  username: string
  password: string
}

// Room
export interface CreateRoomRequest {
  name: string
}
export interface LeaveRoomRequest {
  name: string
}
export interface JoinRoomRequest {
  name: string
}
export interface MessageRoomRequest {
  room_name: string
  msg: string
}

// Game
export interface CreateGameRequest {
  room_name: string
  game_name: string
}
export interface LeaveGameRequest {
  room_name: string
}
export interface JoinGameRequest {
  room_name: string
}
export interface PlayGameRequest {
  room_name: string
  data: any
}