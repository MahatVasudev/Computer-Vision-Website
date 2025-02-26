import React, { useState, useEffect, useRef, useContext } from 'react';
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route, redirect, useNavigate, Outlet } from "react-router-dom";
import Home from "./pages/Home"
import About from "./pages/About"
import Login from './pages/Login'
import Register from './pages/Register'
import Navbar from './components/navbar';
import NavBarLeft from './components/navbar_left';
import CreatorPage from './pages/Creator';
import Post from './pages/Post';
import EditImage from './pages/EditImage';
import { Provider, useDispatch } from 'react-redux';
import store, { persistedStore } from './store';
import { PersistGate } from 'redux-persist/es/integration/react';
import { useVerifyQuery } from './features/user/auth';
import { delCreds, setCreds } from './features/user/auth.local';

export default function App() {
  const [isOpen, setIsOpen] = useState(false);
  const sidebarRef = useRef(null);
  const toggleSidebar = () => {
    setIsOpen(!isOpen);
  };



  useEffect(() => {
    const handleOutsideClick = (event) => {
      if (isOpen && sidebarRef.current && !sidebarRef.current.contains(event.target)) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleOutsideClick);
    return () => {
      document.removeEventListener("mousedown", handleOutsideClick);
    };
  }, [isOpen]);

  return (
    <>

      <Provider store={store}>
        <PersistGate loading={null} persistor={persistedStore}>
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
                    <Route path="u/:creator" element={<CreatorPage />} />
                    <Route element={<AuthCheck />}>
                      <Route path="/post/:post_id/" element={<Post />} />
                      <Route path="/:image_id/edit" element={<EditImage />} />
                    </Route>
                  </Route>
                </Routes>
              </div>
            </div>
          </BrowserRouter>

        </PersistGate>
      </Provider>
    </>

  );
}

const AuthCheck = () => {

  const { data, isLoading, isError } = useVerifyQuery()
  const dispatch = useDispatch()
  const navigate = useNavigate()

  useEffect(() => {

    if (!isLoading) {
      if (isError) {
        alert("User Not Logged in")
        dispatch(setCreds({ key: "status", value: false }))
        dispatch(delCreds({ key: "username" }))
        navigate("/login")
      } else {
        console.log(data)
        dispatch(setCreds({ key: "username", value: data.data?.username }))
      }

    }
  }, [isLoading])

  if (isLoading) {
    return (<div>Loading...</div>)
  }

  if (!isError) {

    return <Outlet />
  }
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <App />
);
