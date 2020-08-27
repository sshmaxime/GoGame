import React, { FunctionComponent } from 'react'
import Text from '../components/text';
import { List } from 'antd';

export type message = {
    from: string,
    msg: string,
    time: string,
}

type props = {
    message: message
    me: boolean

    style?: React.CSSProperties
};

const defaultStyle = (style: React.CSSProperties) => ({
    width: style?.width || "100%",
    maxWidth: style?.maxWidth || undefined,
    padding: style?.padding || "10px",
    marginBottom: style?.paddingBottom || "5px",
    paddingLeft: style?.paddingLeft || "15px",
    paddingRight: style?.paddingRight || "15px",
    marginLeft: style?.marginLeft || undefined,
    boxShadow: style?.boxShadow || "7px 7px 3px #bea6d6",
    textAlign: style?.textAlign || "left",
    borderRadius: style?.borderRadius || "10px",
    color: style?.color || "black",
    background: style?.background || "none",
    backgroundColor: style?.backgroundColor || "white",
    filter: style?.filter || "",
    wordWrap: "break-word",
    display: "block",
    overflowWrap: "break-word",
    wordBreak: "break-all"
}) as React.CSSProperties;

const meMessage = (style: React.CSSProperties): React.CSSProperties => {
    style.color = "white"
    style.background = "rgb(0,120,255)"
    style.background = "-moz-linear-gradient(180deg, rgba(0,120,255,1) 10%, rgba(0,198,255,1) 100%)"
    style.background = "-webkit-linear-gradient(180deg, rgba(0,120,255,1) 10%, rgba(0,198,255,1) 100%)"
    style.background = "linear-gradient(180deg, rgba(0,120,255,1) 10%, rgba(0,198,255,1) 100%)"
    style.filter = "progid:DXImageTransform.Microsoft.gradient(startColorstr='#0078ff',endColorstr='#00c6ff',GradientType=1)"
    return style
}

const otherMessage = (style: React.CSSProperties): React.CSSProperties => {
    style.backgroundColor = "#ffffff"
    style.color = "black"
    return style
}

const Message: FunctionComponent<props> = ({ message, style = {} as React.CSSProperties, me }) => {
    style = me ? { ...style, ...meMessage(style) } : { ...style, ...otherMessage(style) }

    console.log(message)
    const msgColor = me ? "white" : "black"
    const align = me ? "right" : "left"
    const paddingLeft = me ? "200px" : "10px"
    const paddingRight = me ? "10px" : "200px"
    return (
        <div style={{ marginBottom: "5px", paddingRight: paddingRight, paddingLeft: paddingLeft }}>
            <div style={defaultStyle(style)}>
                {!me ?
                    <Text style={{ fontSize: "1.4em", fontWeight: 900, color: "blue", fontFamily: "Source Code Pro" }}>
                        {message.from}
                    </Text>
                    :
                    null
                }

                <Text style={{ whiteSpace: "initial", fontSize: "1.2em", fontWeight: 500, color: msgColor, fontFamily: "Source Code Pro" }}>
                    {message.msg}
                </Text>
            </div>
            {/* <Text style={{ fontSize: "1em", width: "100%", fontWeight: 700, textAlign: align, color: "black", fontFamily: "Source Code Pro" }}>
                {message.time}
            </Text> */}

        </div>
    )
}

export default Message;

