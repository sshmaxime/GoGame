import React, { FunctionComponent } from 'react'

import Card from '../components/card';
import Input from "../components/input"

import { ArrowRightOutlined } from '@ant-design/icons'

import { Space } from 'antd';
import 'antd/dist/antd.less';

type props = {
    username: string,
    password: string,

    handlerUsername: any,
    handlerPassword: any,
    handlerSubmit: any,
};

const AuthComponent: FunctionComponent<props> = ({ children, username, password, handlerUsername, handlerPassword, handlerSubmit }) => {
    return (
        <Card bgColor={"#E0C3FC"} boxShadow={"7px 7px 3px #bea6d6, -7px -7px 3px #ffe0ff"}>
            <form onSubmit={handlerSubmit}>
                <Space size="middle" direction="vertical">
                    <Input value={username}
                        onChange={handlerUsername}
                        height={50}
                        placeholder="Username" />

                    <Input value={password}
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
        </Card>
    )
}

export default AuthComponent;



