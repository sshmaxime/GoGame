import React, { useState, useRef, useEffect } from 'react'
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
import { fabric } from "fabric";
import { Row, Col } from 'antd';


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
  messageRoom: (roomName: string, msg: string) => void,
};

const board = [
  ["1", "0", "1"],
  ["1", "1", "1"],
  ["1", "1", "1"],
]

const defaultState = (): {
} => {
  return {
  }
}

const GameComponent: React.FC<props> = ({ commonState, websocketState, messageRoom }) => {

  const [state, setState] = useState(defaultState())
  const canvasRef = useRef(null)

  return (
    <Card radius={"10px"} minWidth={"400px"} minHeight={"400px"} bgColor={"#ffe0ff"} boxShadow={"7px 7px 3px #bea6d6, -1px -1px 1px #E0C3FC"}>
      <div style={{ margin: "0 auto", width: "70%", height: "100%" }}>
        {board.map((items, index1) => {
          const width = String(100 / items.length) + "%"
          const height = String(100 / board.length) + "%"
          const row = items.map((item, index) => {
            return <div key={index} style={{ display: "inline-block", width: width, height: "100%", fontSize: "5em", position: "relative" }}>
              <div style={{ width: "100%", height: "100%", borderRight: items.length == index + 1 ? "" : "1px solid black" }}>
                <div style={{ margin: "0", position: "absolute", top: "50%", left: "50%", transform: "translate(-50%, -50%)" }}>
                  {item}
                </div>
              </div>
            </div>
          })
          return (
            <div key={index1} style={{ height: height, borderBottom: board.length === index1 + 1 ? "" : "1px solid black" }}>
              {row}
            </div>
          )
        })}
      </div>
    </Card>
  )
}

export default connect(mapStateToProps, mapDispatchToProps)(GameComponent);