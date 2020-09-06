import React, { useState, RefObject } from 'react'
import Card from '../components/card';
import Input from '../components/input';
import Message from '../components/message';
import { ArrowRightOutlined } from '@ant-design/icons'
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
    messageRoom: (roomName: string, msg: string) => websocketActions.MessageRoom(roomName, msg)
  },
  dispatch
);

type props = {
  commonState: commonState,
  websocketState: websocketState,
  messageRoom: (roomName: string, msg: string) => void
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

const ChatComponent: React.FC<props> = ({ commonState, websocketState, messageRoom }) => {

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
        style={{ height: "300px", overflow: "auto" }}
        itemLayout="horizontal"
        dataSource={websocketState.messages}
        renderItem={(item: message, index: number) => (
          <>
            <List.Item style={item.from === websocketState.user?.username ? { padding: "0px", float: "right" } : { padding: "0px", float: "left" }}>
              <Message oldMessage={websocketState.messages[index - 1]} message={item} me={item.from === websocketState.user?.username} />
            </List.Item>
            <div className="fix"></div>
          </>
        )}
      >
        <div ref={state.messagesEndRef} />
      </List>

      <Input value={state.messageField}
        handlerSubmit={handlerSubmit}
        onKeyDown={(event: any) => {
          if (event.key === 'Enter') {
            handlerSubmit(event)
          }
        }}
        float={"right"}
        onChange={handlerMessage}
        height={40}
        width={300}
        placeholder="Write ..."
        icon={<ArrowRightOutlined style={{ paddingRight: "10px", paddingLeft: "10px", fontSize: '2em', verticalAlign: "middle" }} />} />
    </Card>
  )
}

export default connect(mapStateToProps, mapDispatchToProps)(ChatComponent);