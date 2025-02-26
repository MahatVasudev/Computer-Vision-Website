import React, { useEffect, useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons";
import { useLoginMutation } from "../features/user/auth";
import { useNavigate } from "react-router-dom";
import { setCreds } from "../features/user/auth.local";
import { useDispatch } from "react-redux";

const Login = () => {
  const [passwordVisible, setPasswordVisible] = useState(false);
  const [err, setErr] = useState("")
  const navigate = useNavigate()
  const [login, { isLoading }] = useLoginMutation()
  const dispatch = useDispatch()
  useEffect(() => {
    const handleResize = () => {
      if (window.innerWidth >= 768) {
        document.body.style.overflow = "hidden";
      } else {
        document.body.style.overflow = "auto";
      }
    };

    handleResize();

    window.addEventListener("resize", handleResize);

    return () => {
      window.removeEventListener("resize", handleResize);
      document.body.style.overflow = "auto";
    };
  }, []);

  const togglePasswordVisibility = () => {
    setPasswordVisible((prev) => !prev);
  };

  const handleLoginUser = async () => {
    const username = document.getElementById("username")?.value
    const password = document.getElementById("password")?.value

    console.log(username, password)
    // Check whether it is a username or email
    //

    let sttus = "username"

    if (username.includes("@gmail.com")) {
      sttus = "email"
    }

    try {
      const userlog = await login({
        [sttus]: username,
        password: password,
      })

      const data = userlog.data.data.username

      if (userlog.error) {
        setErr(userlog.error.data.error)
        return
      }
      dispatch(setCreds({ key: "status", value: true }))
      navigate(`/u/${data}`)
    } catch (error) {
      setErr(error.message)
    }
  }

  return (
    <div className="relative h-screen w-full flex lg:flex-row flex-col-reverse lg:overflow-hidden md:overflow-hidden">
      {/* Half Image Section (Visible on large screens) */}
      <div className="w-full lg:w-1/2 hidden lg:block relative">
        <img
          src="/login_page.png"
          alt="Half Image"
          className="object-cover h-full w-full"
          style={{ objectPosition: "center top" }}
        />
      </div>

      {/* Image as Background on Small Screens */}
      <div
        className="block lg:hidden absolute top-0 left-0 w-full h-full bg-cover bg-center"
        style={{ backgroundImage: `url('/login_page.png')`, backgroundSize: 'contain' }}
      ></div>

      <div className="w-full lg:w-1/2 flex items-center justify-center bg-gray-50 lg:bg-transparent relative z-10 lg:z-auto">
        <div className="bg-white p-8 lg:p-12 rounded-lg shadow-lg w-full max-w-lg lg:mt-0 mt-8 md:mt-12">
          <h2 className="text-4xl font-bold text-center mb-4">Login</h2>
          <p className="text-center text-gray-500 mb-6">Welcome Back....</p>

          {err != "" ? (<div>
            error : {err}
          </div>) : (<></>)}
          <div className="mb-5">
            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="username">
              Username or Email
            </label>
            <input
              id="username"
              type="text"
              className="w-full px-4 py-3 rounded-lg bg-gray-200 mt-2 border focus:border-blue-500 focus:bg-white focus:outline-none"
            />
          </div>

          <div className="mb-6">
            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
              Password
            </label>
            <div className="relative flex items-center">
              <input
                id="password"
                type={passwordVisible ? "text" : "password"} // Toggle input type based on state
                className="w-full px-4 py-3 rounded-lg bg-gray-200 mt-2 border focus:border-blue-500 focus:bg-white focus:outline-none"
              />
              <button
                type="button"
                onClick={togglePasswordVisibility}
                className="absolute right-2 top-1/2 m-[0.5rem] transform -translate-y-1/2 focus:outline-none"
              >
                <FontAwesomeIcon icon={passwordVisible ? faEyeSlash : faEye} className="text-gray-500 h-5 w-5" />
              </button>
            </div>
          </div>

          {/* Forgot Password Link */}
          <div className="text-right mb-5">
            <a href="#" className="text-sm text-blue-500 hover:underline">
              Forgot Password?
            </a>
          </div>

          {/* Sign In Button */}
          <button
            className="w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 rounded-lg transition duration-200"
            onClick={handleLoginUser}
          >
            Sign In!
          </button>

          {/* Create Account Link */}
          <p className="mt-5 text-center text-gray-500">
            Do not have an Account?{" "}
            <a href="/register" className="text-blue-500 hover:underline">
              Create One
            </a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default Login;
