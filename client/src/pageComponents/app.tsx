import React from 'react'
import Card from '../components/card';
import Text from '../components/text';

import 'antd/dist/antd.css';
import Input from '../components/input';

import { ArrowRightOutlined } from '@ant-design/icons'
import Ws from '../websocket/ws'

type state = {
    message: string
    messages: string[]
};

type props = {
    readonly ws: Ws
};

class AppComponent extends React.Component<props, state> {
    constructor(props: props) {
        super(props)

        this.state = {
            message: "",
            messages: []
        }
        // this.props.ws.joinRoom("demo")
        // this.props.ws.addListener("MESSAGE_ROOM", (data: any) => {
        //     this.setState({
        //         messages: [...this.state.messages, data]
        //     })
        // })
    }

    handlerMessage = (event: any) => {
        this.setState({ message: event.target.value });
    }
    handlerSubmit = (event: any) => {
        // this.props.ws.sendToRoom("demo", "hello")
        event.preventDefault();
    }

    render() {
        return (
            <>
                <Card minWidth={"1000px"} bgColor={"#E0C3FC"} boxShadow={"7px 7px 3px #bea6d6, -7px -7px 3px #ffe0ff"}>
                    <div >
                        <Text fontSize={"3em"} fontWeight={900} fontFamily={"Source Code Pro"}>
                            {this.state.messages.map((message: any) => {
                                return (
                                    <div>
                                        {message}
                                    </div>
                                )
                            })}
                        </Text>
                    </div>
                    <Input value={this.state.message}
                        handlerSubmit={this.handlerSubmit}
                        onKeyDown={(event: any) => {
                            if (event.key === 'Enter') {
                                this.handlerSubmit(event)
                            }
                        }}
                        onChange={this.handlerMessage}
                        height={50}
                        placeholder="Write something ..."
                        icon={<ArrowRightOutlined style={{ paddingRight: "10px", paddingLeft: "10px", fontSize: '2em', verticalAlign: "middle" }} />} />
                </Card>
            </>
        )
    }

}

export default AppComponent;