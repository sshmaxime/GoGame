import React, { FunctionComponent } from 'react'

type props = {
    radius?: string
    minWidth?: string
    bgColor: string
    boxShadow: string
};

const styles = (radius: string, minWidth: string, bgColor: string, boxShadow: string) => ({
    backgroundColor: "white",
    borderRadius: radius,
    display: "inline-block",
    background: bgColor,
    boxShadow: boxShadow,
    padding: "25px",
    minWidth: minWidth,
});

const Card: FunctionComponent<props> = ({ children, radius = "50px", minWidth = "", bgColor, boxShadow }) => {
    return (
        <div style={styles(radius, minWidth, bgColor, boxShadow)}>
            {children}
        </div>
    )
}

export default Card;