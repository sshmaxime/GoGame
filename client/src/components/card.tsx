import React, { FunctionComponent } from 'react'

type props = {
    radius?: string
    minWidth?: string
    bgColor: string
    boxShadow: string
    minHeight?: string
};

const styles = (radius: string, minWidth: string, bgColor: string, boxShadow: string, minHeight: string) => ({
    backgroundColor: "white",
    borderRadius: radius,
    display: "inline-block",
    background: bgColor,
    boxShadow: boxShadow,
    padding: "30px",
    width: minWidth,
    height: minHeight
});

const Card: FunctionComponent<props> = ({ children, radius = "50px", minWidth = "", minHeight = "", bgColor, boxShadow }) => {
    return (
        <div style={styles(radius, minWidth, bgColor, boxShadow, minHeight)}>
            {children}
        </div>
    )
}

export default Card;