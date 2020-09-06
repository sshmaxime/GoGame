import React, { FunctionComponent, useState } from 'react'
import Text from '../components/text';
import Card from '../components/card';
import { Dispatch, AnyAction, bindActionCreators } from "redux";
import { connect } from "react-redux";
import { IAppState } from "../store/reducers";

import Input from '../components/input';
import Message from '../components/message';

import { Row, Col } from 'antd';
import ChatComponent from './chat';
import { websocketState } from '../store/reducers/websocket.reducer';
import { websocketActions } from "../store/actions/websocket.actions";
import { CloseCircleFilled } from '@ant-design/icons';

const mapStateToProps = (state: IAppState): IAppState => {
  return {
    ...state,
    commonState: state.commonState,
  };
};

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => bindActionCreators(
  {
    createRoom: (roomName: string) => websocketActions.CreateRoom(roomName),
    joinRoom: (roomName: string) => websocketActions.JoinRoom(roomName),
    leaveRoom: (roomName: string) => websocketActions.LeaveRoom(roomName)
  },
  dispatch
);

type props = {
  websocketState: websocketState,
  createRoom: (roomName: string) => void
  joinRoom: (roomName: string) => void
  leaveRoom: (roomName: string) => void
}

const defaultState = (): {
  createRoomField: string
  joinRoomField: string
} => {
  return {
    createRoomField: "",
    joinRoomField: ""
  }
}

const AppComponent: FunctionComponent<props> = ({ children, websocketState, createRoom, joinRoom, leaveRoom }) => {
  const [state, setState] = useState(defaultState())


  const handlerCreateRoom = (event: any) => {
    setState({ ...state, createRoomField: event.target.value });
  }
  const handlerJoinRoom = (event: any) => {
    setState({ ...state, joinRoomField: event.target.value });
  }

  const handlerSubmitCreate = (event: any) => {
    createRoom(state.createRoomField)
    setState({ ...state, createRoomField: "", joinRoomField: "" })

    event.preventDefault();
  }

  const handlerSubmitJoin = (event: any) => {
    joinRoom(state.joinRoomField)
    setState({ ...state, createRoomField: "", joinRoomField: "" })

    event.preventDefault();
  }

  return (
    <Card radius={"10px"} minHeight={"100%"} minWidth={"100%"} bgColor={"#E0C3FC"} boxShadow={"7px 7px 3px #bea6d6, -7px -7px 3px #ffe0ff"}>
      {websocketState.room !== undefined ?
        <>
          <Row>
            <Col span={20}>
              <Text style={{ fontSize: "2em", fontFamily: "Source Code Pro" }}>Room: {websocketState.room.name}</Text>
            </Col>

            <Col span={4}>
              <Text style={{ textAlign: "right", fontSize: "2em", fontFamily: "Source Code Pro" }}>
                <CloseCircleFilled className="link" onClick={() => {
                  if (websocketState.room !== undefined) {
                    leaveRoom(websocketState.room.name)
                  }
                }} />
              </Text>
            </Col>

          </Row>
          < ChatComponent />
        </>
        :
        <>
          <Row style={{ marginBottom: "25px" }}>
            <Card radius={"10px"} bgColor={"#ffe0ff"} boxShadow={"7px 7px 3px #bea6d6, -1px -1px 1px #E0C3FC"}>
              <Input value={state.createRoomField}
                onChange={handlerCreateRoom}
                handlerSubmit={handlerSubmitCreate}
                onKeyDown={(event: any) => {
                  if (event.key === 'Enter') {
                    handlerSubmitCreate(event)
                  }
                }}
                height={50}
                placeholder="Create room" />
            </Card>
          </Row>

          <Row>
            <Card radius={"10px"} bgColor={"#ffe0ff"} boxShadow={"7px 7px 3px #bea6d6, -1px -1px 1px #E0C3FC"}>
              <Input value={state.joinRoomField}
                onChange={handlerJoinRoom}
                handlerSubmit={handlerSubmitJoin}
                onKeyDown={(event: any) => {
                  if (event.key === 'Enter') {
                    handlerSubmitJoin(event)
                  }
                }}
                height={50}
                placeholder="Join room" />
            </Card>
          </Row>

        </>
      }
    </Card>
  )
}

export default connect(mapStateToProps, mapDispatchToProps)(AppComponent);