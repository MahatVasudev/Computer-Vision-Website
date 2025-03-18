import React from 'react'
import { useSelector } from 'react-redux'
import { Link } from 'react-router-dom'

const CustomButton = ({ text, href = "#", color = "black", bg_color = "#01A456", className = 'mr-5', onClick }) => {


  const preferedColor = useSelector((state) => state.settings.prefered_color)

  bg_color = bg_color === "self" ? preferedColor : bg_color

  return (
    <>
      <button onClick={onClick || null}
        className={`${className} py-2 px-4 rounded-full text-black font-bold`} style={{ backgroundColor: bg_color, color }}>
        <a href={href}>{text}</a>
      </button>
    </>
  )
}

export default CustomButton
