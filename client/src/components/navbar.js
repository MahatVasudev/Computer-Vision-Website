import React from 'react';

const Navbar = () => {
	return (
		<nav className="flex justify-between bg-[#2B2A2A] p-[1%]">
			<div className="text-white text-4xl font-inter font-bold ml-[3%]"> On-Sight </div>
			<div className='bg-red-600 rounded-[50%] w-[3rem] text-center h-[3rem]'>
				<a className="text-white" href='#'>user</a>
			</div>
		</nav >
	)
}

export default Navbar
