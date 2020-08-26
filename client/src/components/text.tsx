import React, { FunctionComponent } from 'react'

type props = {
    fontSize: string
    icon?: string
    fontWeight?: number
    fontFamily: string
};

const styles = (fontSize: string, fontWeight: number, fontFamily: string) => ({
    fontFamily: fontFamily,
    fontWeight: fontWeight,
    fontSize: fontSize,
    textShadow: "-3px -3px 0px #ffe0ff"
});

const Text: FunctionComponent<props> = ({ children, fontSize, icon, fontWeight = 700, fontFamily }) => {
    return (
        <div style={styles(fontSize, fontWeight, fontFamily)}>
            {children}
            <span style={{ textShadow: "none", marginLeft: "10px" }}>{icon}</span>
        </div>
    )
}

export default Text;