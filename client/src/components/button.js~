import React from 'react'
import { Link } from 'react-router-dom'

const CustomButton = ({ text, color = "black", bg_color = "#01A456", style = 'mr-5' }) => {

  console.log("color", color)
  console.log("Sytle ", `${style} py-2 px-4 rounded-full text-black bg-[${bg_color}] text-[${color}] font-bold`)
  console.log("bg_color", bg_color)
  return (
    <>
      <button
        className={`${style} py-2 px-4 rounded-full text-black font-bold`} style={{ backgroundColor: bg_color, color }}>
        <a href={"/login"}>{text}</a>
      </button>
    </>
  )
}

export default CustomButton
