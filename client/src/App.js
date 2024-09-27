// src/App.js
import React from 'react';
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home"
import About from "./pages/About"
import Login from './pages/Login'
import Register from './pages/Register'

export default function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/">
					<Route index element={<Home />} />
					<Route path="About" element={<About />} />
					<Route path="Login" element={<Login />} />
					<Route path="Register" element={<Register />} />
				</Route>
			</Routes>
		</BrowserRouter>
	);
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);
