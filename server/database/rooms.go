package database

import (
	"fmt"
	"github.com/GoGame/types"
)

type RoomDatabase struct {
	rooms map[string]*types.Room
}

var roomDatabase RoomDatabase

func InitRoomDatabase() error {
	roomDatabase.rooms = map[string]*types.Room{}
	return nil
}

func CreateRoom(ID string) (room *types.Room, err error) {
	errorMsg := "error while creating room"

	if room, err = GetRoomByID(ID); err != nil {
		return nil, fmt.Errorf("%v: room [%v] already exist", errorMsg, ID)
	}

	room = &types.Room{
		ID: ID,
	}
	roomDatabase.rooms[ID] = room
	return room, nil
}

func GetRoomByID(ID string) (room *types.Room, err error) {
	errorMsg := "error while getting room"

	room, _ = roomDatabase.rooms[ID]
	if room == nil {
		return nil, fmt.Errorf("%v: game [%v] doesn't exist", errorMsg, ID)
	}
	return room, nil
}

func GetAllRooms() (map[string]*types.Room, error) {
	return roomDatabase.rooms, nil
}
