import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBars } from '@fortawesome/free-solid-svg-icons';
import SearchBar from './search_bar';
import ProfileCircle from './profileCircle';
import CustomButton from './button';
import { useSelector } from 'react-redux';
import { useEffect } from 'react';
const Navbar = ({ toggleSidebar }) => {  // Receive toggleSidebar as a prop
  const isloggedin = useSelector((state) => state.auth.status) || false
  const username = useSelector((state) => state.auth.username)

  const theme = useSelector((state) => state.settings.dark)

  useEffect(() => {
    document.documentElement.classList.toggle("dark", theme === 1);
  }, [theme]);
  return (
    <nav className="flex sticky dark:text-white dark:bg-[#1C1C1C] z-40 top-0 items-center justify-between bg-[#FFFFFF] border-b-2 dark:border-white border-black p-[1%]">
      <button className='ml-4' onClick={toggleSidebar}>  {/* Trigger sidebar toggle */}
        <FontAwesomeIcon icon={faBars} className="h-[2rem]" />
      </button>
      <div className="mr-auto dark:text-white text-black text-4xl font-inter font-bold ml-[3%]"> On-Sight </div>
      <SearchBar top_style={"mr-auto"} />
      {
        !isloggedin ?
          (<>
            <CustomButton text={"SignUp"} href={"/register"} bg_color={"#D9D9D9"} style={"mr-2"} />
            <CustomButton text={"Login"} href={"/login"} color={"white"} bg_color={"black"} />
          </>) :
          <ProfileCircle to={username} />
      }
    </nav>
  )
}

export default Navbar;
