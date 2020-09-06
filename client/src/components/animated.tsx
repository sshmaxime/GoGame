import React, { FunctionComponent } from 'react'
import { animated } from 'react-spring'

type props = {
  animation: any
};

const Animated: FunctionComponent<props> = ({ children, animation }) => {
  return (
    <animated.span style={animation}>
      {children}
    </animated.span>
  )
}

export default Animated;
