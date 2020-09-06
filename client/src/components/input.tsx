import React, { FunctionComponent } from 'react'
import Animated from './animated'
import { useSpring } from 'react-spring'


type props = {
  height: number,
  width?: number,
  float?: string,
  placeholder: string,
  icon?: any,
  type?: string,
  value: string,
  onChange: any,
  onKeyDown?: any,
  handlerSubmit?: any
};

const styles = (height: number, width: number | undefined) => ({
  fontFamily: "Source Code Pro",
  width: width,
  fontWeight: 500,
  letterSpacing: "0px",
  fontSize: "1.3em",
  height: height,
  border: "2px solid black",
  borderRadius: "15px",
  padding: "10px 15px 10px 15px",
  verticalAlign: "middle",
}) as React.CSSProperties;

const Input: FunctionComponent<props> = ({ float = "none", height, width = undefined, type, placeholder, icon = null, value, onChange, onKeyDown, handlerSubmit }) => {
  return (
    <div style={{ float: float } as React.CSSProperties}>
      <span>
        <input onKeyDown={onKeyDown} value={value} onChange={onChange} type={type} style={styles(height, width)} placeholder={placeholder} />
      </span>
      {icon == null ? null : (
        <span onClick={handlerSubmit} className="hoverPointer"  >
          {icon}
        </span>)
      }


    </div >
  )
}

export default Input;
