import React, { FunctionComponent } from 'react'

type props = {
  icon?: any,
  iconSize?: string,
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
  width: style?.width || undefined,
  paddingRight: style?.paddingRight || undefined,
  paddingLeft: style?.paddingLeft || undefined,
  paddingTop: style?.paddingTop || undefined,
}) as React.CSSProperties;

const Text: FunctionComponent<props> = ({ children, icon, iconSize, style }) => {
  return (
    <div style={defaultStyle(style)}>
      {children}
      {icon ? <span style={{ fontSize: iconSize, textShadow: "none", marginLeft: "10px" }}>{icon}</span> : null}

    </div>
  )
}

export default Text;