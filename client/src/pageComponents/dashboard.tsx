import React, { FunctionComponent } from 'react'
import Card from '../components/card';
import Text from '../components/text';
import { List } from 'antd';

import { Divider } from 'antd';
import { Dispatch, AnyAction, bindActionCreators } from "redux";
import { connect } from "react-redux";
import { IAppState } from "../store/reducers";
import { commonState } from "../store/reducers/common.reducer";
import { websocketState } from '../store/reducers/websocket.reducer';
import { websocketActions } from "../store/actions/websocket.actions";
import { user, room } from '../store/types/types';


const mapStateToProps = (state: IAppState): IAppState => { return { ...state, }; };
const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => bindActionCreators(
  {
  },
  dispatch
);

type props = {
  commonState: commonState,
  websocketState: websocketState,
};

const DashboardComponent: FunctionComponent<props> = ({ children, commonState, websocketState }) => {
  return (
    <Card radius={"10px"} minWidth={"100%"} bgColor={"#ffe0ff"} boxShadow={"7px 7px 3px #bea6d6, -1px -1px 1px #E0C3FC"}>
      <Text style={{ fontSize: "2em", fontWeight: 900, fontFamily: "Source Code Pro", textShadow: "-2px -2px 0px #e0c3fc" }}>
        Dashboard
      </Text>

      <Divider plain orientation="left">
        <Text style={{ fontSize: "1.2em", fontFamily: "Source Code Pro" }}>
          Online Users
        </Text>
      </Divider>
      <List
        split={false}
        itemLayout="horizontal"
        locale={{ emptyText: <></> }}
        dataSource={websocketState.online_users}
        renderItem={(item: user, index: number) => (
          <>
            <List.Item style={{ fontFamily: "Source Code Pro", padding: "0px" }}>
              - {item.username}
            </List.Item>
            <div className="fix"></div>
          </>
        )}
      >
      </List>

      <Divider plain orientation="left">
        <Text style={{ fontSize: "1.2em", fontFamily: "Source Code Pro" }}>
          Available Rooms
        </Text>
      </Divider>
      <List
        split={false}
        itemLayout="horizontal"
        locale={{ emptyText: <></> }}
        dataSource={websocketState.online_rooms}
        renderItem={(item: room, index: number) => (
          <>
            <List.Item style={{ fontFamily: "Source Code Pro", padding: "0px" }}>
              - {item.name}
            </List.Item>
            <div className="fix"></div>
          </>
        )}
      >
      </List>
    </Card >
  )
}

export default connect(mapStateToProps, mapDispatchToProps)(DashboardComponent);