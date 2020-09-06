import React, { FunctionComponent } from 'react'
import Text from '../components/text';

import { Dispatch, AnyAction, bindActionCreators } from "redux";
import { connect } from "react-redux";
import { IAppState } from "../store/reducers";
import { commonState } from "../store/reducers/common.reducer";
import { websocketState } from '../store/reducers/websocket.reducer';

import { Row, Col } from 'antd';

const mapStateToProps = (state: IAppState): IAppState => { return { ...state, }; };
const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => bindActionCreators(
  {
  },
  dispatch
);

type props = {
  commonState: commonState,
  websocketState: websocketState
};

const HeaderComponent: React.FC<props> = ({ commonState, websocketState }) => {
  return (
    <Row style={{ marginBottom: "20px" }}>
      <Col span={18}>
        <Text style={{ fontSize: "4em", fontWeight: 900, fontFamily: "Montserrat", textShadow: "-3px -3px 0px #ffe0ff" }} icon="✌️">
          Hello
          </Text>
      </Col>

      <Col span={6} style={{ display: "flex", justifyContent: "flex-end", alignItems: "center" }}>
        <Text style={{ fontSize: "1.5em", fontWeight: 900, fontFamily: "Source Code Pro", color: "#ffffff", textShadow: "0px 0px 8px rgba(255,255,255,0.8)" }} iconSize="2em" icon="🧑‍💻">
          {websocketState.user?.username}
        </Text>
      </Col>
    </Row>
  )
}

export default connect(mapStateToProps, mapDispatchToProps)(HeaderComponent);