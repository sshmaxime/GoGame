import React, { FunctionComponent } from 'react'
import Card from '../components/card';
import Text from '../components/text';

import { Divider } from 'antd';
import 'antd/dist/antd.css';

type props = {
};

const DashboardComponent: FunctionComponent<props> = ({ children }) => {
    return (
        <Card radius={"10px"} width={"500px"} bgColor={"#ffe0ff"} boxShadow={"7px 7px 3px #bea6d6, -1px -1px 1px #E0C3FC"}>
            <Text style={{ fontSize: "3em", fontWeight: 900, fontFamily: "Source Code Pro", textShadow: "-2px -2px 0px #e0c3fc" }}>
                Dashboard
            </Text>

            <Divider plain orientation="left">
                <Text style={{ fontSize: "1.5em", fontFamily: "Source Code Pro" }}>
                    Users
                </Text>
            </Divider>

            <Divider plain orientation="left">
                <Text style={{ fontSize: "1.5em", fontFamily: "Source Code Pro" }}>
                    Rooms
                    </Text>
            </Divider>
        </Card >
    )
}

export default DashboardComponent;