export interface user {
  username: string
}
export interface room {
  name: string
}
export interface message {
  from: string,
  msg: string,
  time: string,
}
export interface game {
  board: number[][],
  who_to_play: number[],
  victory: number,
  players: Map<string, number>,
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
  name: string
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