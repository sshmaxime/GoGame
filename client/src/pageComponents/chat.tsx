import React, { useState, RefObject } from 'react'
import Card from '../components/card';
import Input from '../components/input';
import Message from '../components/message';
import { RocketFilled, SendOutlined } from '@ant-design/icons'
import { List } from 'antd';
import { message } from '../store/types/types'

import { Dispatch, AnyAction, bindActionCreators } from "redux";
import { connect } from "react-redux";
import { IAppState } from "../store/reducers";
import { commonState } from "../store/reducers/common.reducer";
import { websocketState } from '../store/reducers/websocket.reducer';
import { websocketActions } from "../store/actions/websocket.actions";

const mapStateToProps = (state: IAppState): IAppState => { return { ...state, }; };
const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => bindActionCreators(
  {
    messageRoom: (roomName: string, msg: string) => websocketActions.MessageRoom(roomName, msg),
    createGame: (roomName: string, gameName: string) => websocketActions.CreateGame(roomName, gameName),
    joinGame: (roomName: string) => websocketActions.JoinGame(roomName),
  },
  dispatch
);

type props = {
  commonState: commonState,
  websocketState: websocketState,
  messageRoom: (roomName: string, msg: string) => void,
  createGame: (roomName: string, gameName: string) => void,
  joinGame: (roomName: string) => void,
  small?: boolean
};

const defaultState = (): {
  messageField: string,
  messagesEndRef: RefObject<any>,
} => {
  return {
    messageField: "",
    messagesEndRef: React.createRef()
  }
}

const ChatComponent: React.FC<props> = ({ commonState, websocketState, messageRoom, createGame, joinGame, small }) => {

  const [state, setState] = useState(defaultState())

  React.useEffect(() => {
    state.messagesEndRef.current.scrollIntoView({ behavior: "smooth" })
  }, [websocketState.messages]);

  const handlerMessage = (event: any) => {
    setState({ ...state, messageField: event.target.value });
    event.preventDefault();
  }

  const handlerSubmit = (event: any) => {
    if (websocketState.room !== undefined) {
      messageRoom(websocketState.room.name, state.messageField)
    }
    setState({ ...state, messageField: "" })
    event.preventDefault();
  }

  return (
    <Card radius={"10px"} minWidth={"100%"} bgColor={"#ffe0ff"} boxShadow={"7px 7px 3px #bea6d6, -1px -1px 1px #E0C3FC"}>
      <List
        split={false}
        style={{ height: small ? "150px" : "300px", overflow: "auto" }}
        itemLayout="horizontal"
        dataSource={websocketState.messages}
        renderItem={(item: message, index: number) => (
          <>
            {item.from === "SERVER" ?
              websocketState.inGame ?
                null
                :
                <List.Item style={item.from === websocketState.user?.username ? { padding: "0px", float: "right" } : { padding: "0px", float: "left" }}>
                  {item.msg}
                  <button onClick={() => joinGame(websocketState.room.name)}>JOIN</button>
                </List.Item>
              :
              <>
                <List.Item style={item.from === websocketState.user?.username ? { padding: "0px", float: "right" } : { padding: "0px", float: "left" }}>
                  <Message oldMessage={websocketState.messages[index - 1]} message={item} me={item.from === websocketState.user?.username} />
                </List.Item>
                <div className="fix"></div>
              </>
            }
          </>
        )}
      >
        <div ref={state.messagesEndRef} />
      </List>

      <div style={{ float: "right", display: "inline-block" }} >
        <span>
          <span onClick={() => createGame(websocketState.room.name, "tictactoe")} className="hoverPointer"  >
            <RocketFilled style={{ paddingRight: "10px", paddingLeft: "10px", fontSize: '2em', verticalAlign: "middle" }} />
          </span>
          <input onKeyDown={(event: any) => {
            if (event.key === 'Enter') {
              handlerSubmit(event)
            }
          }} value={state.messageField} onChange={handlerMessage} style={{
            fontFamily: "Source Code Pro",
            fontWeight: 500,
            fontSize: "1.3em",
            border: "2px solid black",
            borderRadius: "15px",
            padding: "10px 15px 10px 15px",
            verticalAlign: "middle",
            height: "40px"
          }} placeholder={"Write ..."} />
        </span>
        <span onClick={handlerSubmit} className="hoverPointer"  >
          <SendOutlined style={{ color: "rgb(100,120,255)", paddingRight: "10px", paddingLeft: "10px", fontSize: '2em', verticalAlign: "middle" }} />
        </span>
      </div>
    </Card>
  )
}

export default connect(mapStateToProps, mapDispatchToProps)(ChatComponent);