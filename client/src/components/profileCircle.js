import React from 'react'

const ProfileCircle = ({ to }) => {

  const url = `http://localhost:3000/u/${to}`
  return (
    <div className='bg-red-600 rounded-[50%] w-[3rem] text-center h-[3rem]' >
      <a className="text-white" href={url}>{to}</a>
    </div >

  )
}

export default ProfileCircle
