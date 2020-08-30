package hub

import (
	"errors"
	"fmt"
)

func CreateRoomRequest(cli *Client, roomName string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	if getRoom(roomName) != nil {
		return fmt.Errorf("room name %v already taken", roomName)
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
		return fmt.Errorf("room %s doesn't exist", roomName)
	}

	if err := cli.joinRoom(roomName); err != nil {
		return err
	}

	if err := room.addClient(cli); err != nil {
		return err
	}

	return room.Room
}

func LeaveRoomRequest(cli *Client, roomName string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	room := getRoom(roomName)
	if room == nil {
		return fmt.Errorf("room %s doesn't exist", roomName)
	}

	if err := room.removeClient(cli); err != nil {
		return err
	}

	if err := cli.leaveRoom(roomName); err != nil {
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
		return fmt.Errorf("room %s doesn't exist", roomName)
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
