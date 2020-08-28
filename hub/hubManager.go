package hub

import (
	"errors"
)

func CreateRoomRequest(cli *Client, roomName string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	if getRoom(roomName) != nil {
		return errors.New("room name already taken")
	}

	newRoom := CreateRoom(cli, roomName)
	if err := cli.joinRoom(roomName); err != nil {
		return err
	}

	if err := newRoom.addClient(cli); err != nil {
		return err
	}

	rooms[roomName] = newRoom
	return newRoom.Room
}

func JoinRoomRequest(cli *Client, roomName string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	room := getRoom(roomName)
	if room == nil {
		return errors.New("room doesn't exist")
	}

	if err := cli.joinRoom(roomName); err != nil {
		return err
	}

	if err := room.addClient(cli); err != nil {
		return err
	}

	return room.Room
}

func SendMessageToRoomRequest(cli *Client, namespace string, roomName string, msg string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	room := getRoom(roomName)
	if room == nil {
		return errors.New("room doesn't exist")
	}

	if room.Clients[cli.Socket.ID()] == nil {
		return errors.New("client is not a part of this room")
	}

	newMsg := Message{
		From: cli.User.Username,
		Msg:  msg,
	}

	handler.BroadcastToRoom(namespace, room.Room.Name, MESSAGE_ROOM, newMsg)

	return newMsg
}
