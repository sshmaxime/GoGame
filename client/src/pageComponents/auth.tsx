import React, { FunctionComponent, useState } from 'react'

import Card from '../components/card';
import Input from "../components/input"
import Text from "../components/text"

import { Dispatch, AnyAction, bindActionCreators } from "redux";
import { connect } from "react-redux";
import { IAppState } from "../store/reducers";
import { websocketActions } from "../store/actions/websocket.actions";
import { Divider } from 'antd';

import { ArrowRightOutlined } from '@ant-design/icons'
import { useTransition, animated } from 'react-spring'

import { Space } from 'antd';

const mapStateToProps = (state: IAppState): IAppState => { return { ...state, }; };
const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => bindActionCreators(
  {
    login: (username: string, password: string) => websocketActions.Login(username, password),
    register: (username: string, password: string) => websocketActions.Register(username, password)
  },
  dispatch
);

type props = {
  login: (username: string, password: string) => void
  register: (username: string, password: string) => void
};

const defaultState = (): {
  isRegister: boolean
  usernameField: string
  passwordField: string
} => {
  return {
    isRegister: false,
    usernameField: "",
    passwordField: ""
  }
}

const AuthComponent: FunctionComponent<props> = ({ children, login, register }) => {
  const [state, setState] = useState(defaultState())

  const handlerUsername = (event: any) => {
    setState({ ...state, usernameField: event.target.value });
    event.preventDefault();
  }
  const handlerPassword = (event: any) => {
    setState({ ...state, passwordField: event.target.value });
    event.preventDefault();
  }

  const handlerSubmit = (event: any) => {
    if (state.isRegister) {
      setState({ ...state, isRegister: false, usernameField: "", passwordField: "" })
      register(state.usernameField, state.passwordField)
    } else {
      login(state.usernameField, state.passwordField)
    }
    event.preventDefault();
  }

  return (
    <Card bgColor={"#E0C3FC"} boxShadow={"7px 7px 3px #bea6d6, -7px -7px 3px #ffe0ff"}>
      <>
        <Divider orientation="left" plain>
          <Text style={{ fontSize: "1.5em" }}>
            {state.isRegister ? <>Register</> : <>Login</>}
          </Text>
        </Divider>
        <Space size="middle" direction="vertical">

          <form onSubmit={handlerSubmit}>
            <Space size="middle" direction="vertical">
              <Input value={state.usernameField}
                onChange={handlerUsername}
                handlerSubmit={handlerSubmit}
                onKeyDown={(event: any) => {
                  if (event.key === 'Enter') {
                    handlerSubmit(event)
                  }
                }}
                height={50}
                placeholder="Username" />

              <Input value={state.passwordField}
                handlerSubmit={handlerSubmit}
                onKeyDown={(event: any) => {
                  if (event.key === 'Enter') {
                    handlerSubmit(event)
                  }
                }}
                onChange={handlerPassword}
                type="password" height={50}
                placeholder="Password"
                icon={<ArrowRightOutlined style={{ paddingRight: "10px", paddingLeft: "10px", fontSize: '2em', verticalAlign: "middle" }} />} />
            </Space >
          </form>
          <Text style={{ fontSize: "0.9em" }}>
            {
              state.isRegister ?
                <>Already have an account ? <span className="link" onClick={() => { setState({ ...state, isRegister: false, usernameField: "", passwordField: "" }) }} style={{ color: "blue" }}>Login</span></>
                :
                <>Don't have an account ? <span className="link" onClick={() => { setState({ ...state, isRegister: true, usernameField: "", passwordField: "" }) }} style={{ color: "blue" }}>Register</span></>
            }
          </Text>
        </Space >
      </>
    </Card>
  )
}

export default connect(mapStateToProps, mapDispatchToProps)(AuthComponent);


