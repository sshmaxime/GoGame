import React, { FunctionComponent } from 'react'

type props = {
    icon?: any,
    style?: React.CSSProperties
};

const defaultStyle = (style: React.CSSProperties | undefined) => ({
    fontFamily: style?.fontFamily || "Montserrat",
    fontWeight: style?.fontWeight || "500",
    fontSize: style?.fontSize || "1em",
    textShadow: style?.textShadow || "none",
    letterSpacing: style?.letterSpacing || "none",
    textAlign: style?.textAlign || "none",
    color: style?.color || "black",
    overflow: style?.overflow || undefined,
    whiteSpace: style?.whiteSpace || undefined,
    textOverflow: style?.textOverflow || undefined,
    width: style?.width || undefined
}) as React.CSSProperties;

const Text: FunctionComponent<props> = ({ children, icon, style }) => {
    return (
        <div style={defaultStyle(style)}>
            {children}
            <span style={{ textShadow: "none", marginLeft: "10px" }}>{icon}</span>
        </div>
    )
}

export default Text;