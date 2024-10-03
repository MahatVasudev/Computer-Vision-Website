import { faHome, faImage, faPlay, faPlus, faBars } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";
import { NavBarLeftConst } from "../constants/navbar_left_constants";

const NavBarLeft = ({ isOpen, toggleSidebar, sidebarRef }) => {
  return (
    <nav
      ref={sidebarRef}
      className={`fixed top-0 left-0 h-full w-64 bg-white border-r border-black transform 
      ${isOpen ? 'translate-x-0' : '-translate-x-full'} transition-transform duration-300 z-50 flex flex-col`}
    >
      {/* Logo Section */}
      <div className="p-4 flex items-center justify-center">
        <img src="https://via.placeholder.com/100" alt="Logo" className="w-24 h-auto" />
      </div>

      {/* faBars button to close sidebar */}
      <div className="p-4">
        <button onClick={toggleSidebar} className="flex items-center">
          <FontAwesomeIcon icon={faBars} className="text-black h-6 w-6" />
          <span className="ml-2 font-bold text-black">Close</span>
        </button>
      </div>


      {/* Scrollable area for navigation links */}
      <div className="flex-1 overflow-y-auto">
        {NavBarLeftConst.map((content, key) => {
          return (
            <>
              <a
                href={content.href}
                className="flex items-center p-4 transition-colors duration-200 hover:bg-gray-200"
                key={key}
              >
                <FontAwesomeIcon icon={content.icon} className="mr-2" />
                <span className="font-bold text-black">{content.name}</span>
              </a>

            </>)
        })}

      </div>

      {/* Create button at the bottom */}
      <div className="p-4 border-t border-gray-300">
        <a
          href="/create"
          className="flex items-center p-4 transition-colors duration-200 hover:bg-gray-200"
        >
          <FontAwesomeIcon icon={faPlus} className="mr-2" />
          <span className="font-bold text-black">Create</span>
        </a>
      </div>
    </nav >
  );
};

export default NavBarLeft;
