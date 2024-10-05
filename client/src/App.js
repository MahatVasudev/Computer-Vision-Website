import React, { useState, useEffect, useRef } from 'react';
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home"
import About from "./pages/About"
import Login from './pages/Login'
import Register from './pages/Register'
import Navbar from './components/navbar';
import NavBarLeft from './components/navbar_left';
import CreatorPage from './pages/Creator';

export default function App() {
  const [isOpen, setIsOpen] = useState(false); // Sidebar state
  const sidebarRef = useRef(null); // Ref for sidebar

  const toggleSidebar = () => {
    setIsOpen(!isOpen); // Toggle sidebar state
  };

  // Handle closing the sidebar when clicking outside
  useEffect(() => {
    const handleOutsideClick = (event) => {
      if (isOpen && sidebarRef.current && !sidebarRef.current.contains(event.target)) {
        setIsOpen(false); // Close sidebar if clicked outside
      }
    };

    document.addEventListener("mousedown", handleOutsideClick); // Listen for clicks
    return () => {
      document.removeEventListener("mousedown", handleOutsideClick); // Cleanup
    };
  }, [isOpen]);

  return (
    <>
      <BrowserRouter>
        <div className='flex-1 flex flex-col'>
          {/* Pass toggleSidebar to the Navbar */}
          <Navbar toggleSidebar={toggleSidebar} />
          {/* Pass ref and isOpen state to the NavBarLeft */}
          <NavBarLeft isOpen={isOpen} toggleSidebar={toggleSidebar} sidebarRef={sidebarRef} />
          <div className='flex-1 p-4 overflow-y-auto h-[calc(100vh-64px)]'>
            <Routes>
              <Route path="/">
                <Route index element={<Home />} />
                <Route path="about" element={<About />} />
                <Route path="login" element={<Login />} />
                <Route path="register" element={<Register />} />
                <Route path=":creator" element={<CreatorPage />} />
              </Route>
            </Routes>
          </div>
        </div>
      </BrowserRouter>
    </>
  );
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);
