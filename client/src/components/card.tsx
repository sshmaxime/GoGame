import React, { FunctionComponent } from 'react'

type props = {
    radius?: string
    width?: string
    bgColor: string
    boxShadow: string
};

const styles = (radius: string, width: string, bgColor: string, boxShadow: string) => ({
    backgroundColor: "white",
    borderRadius: radius,
    display: "inline-block",
    background: bgColor,
    boxShadow: boxShadow,
    padding: "30px",
    width: width,
});

const Card: FunctionComponent<props> = ({ children, radius = "50px", width = "", bgColor, boxShadow }) => {
    return (
        <div style={styles(radius, width, bgColor, boxShadow)}>
            {children}
        </div>
    )
}

export default Card;