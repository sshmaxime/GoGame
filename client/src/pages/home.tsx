import React from 'react'

import { Dispatch, AnyAction, bindActionCreators } from "redux";
import { connect } from "react-redux";
import { IAppState } from "../store/reducers";
import { websocketActions } from "../store/actions/websocket.actions";

import AuthComponent from "../pageComponents/auth"
import DashboardComponent from "../pageComponents/dashboard"

import HeaderComponent from '../pageComponents/header';
import AppComponent from '../pageComponents/app';
import { Row, Col } from 'antd';
import Text from '../components/text';

const mapStateToProps = (state: IAppState): IAppState => {
  return {
    ...state,
    commonState: state.commonState,
  };
};

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => bindActionCreators(
  {
    login: (username: string, password: string) => websocketActions.Login(username, password)
  },
  dispatch
);

class HomePage extends React.Component<ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>> {
  constructor(props: any) {
    super(props)
  }

  render() {
    return (
      <div style={{ height: "100%", width: "100%", padding: "50px 100px 50px 100px" }}>
        <HeaderComponent />
        {
          this.props.websocketState.ready ?
            <Row>
              <Col span={16}>
                {this.props.websocketState.user !== undefined ?
                  <AppComponent />
                  :
                  <AuthComponent />
                }
              </Col>
              <Col span={7} offset={1}>
                <DashboardComponent />
              </Col>
            </Row>
            :
            <Text style={{ fontSize: "1em" }}>Trying to connect to the server ...</Text>
        }
      </div >
    )
  }

}

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);