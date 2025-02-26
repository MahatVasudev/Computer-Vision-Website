import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons";
import CenteredModal from "../components/OTPMODAL";
import { useRegisterMutation, useSendEmailMutation, useVerifyOTPMutation, useVerifyQuery } from "../features/user/auth";
import { useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";

const Register = () => {
  const [passwordVisible, setPasswordVisible] = useState(false); // State for password visibility
  const [confirmPasswordVisible, setConfirmPasswordVisible] = useState(false); // State for confirm password visibility
  const [password, setPassword] = useState(""); // Password state
  const [confirmPassword, setConfirmPassword] = useState(""); // Confirm password state
  const [errorMessage, setErrorMessage] = useState(""); // Error message state
  const [email, setEmail] = useState("")
  const [availableMessage, setAvailableMessage] = useState({})
  const [otp, setOTP] = useState()
  const [sentSubmit, setSentSubmit] = useState(false)
  const [step, setStep] = useState(0)
  const [sendEmail, { }] = useSendEmailMutation()
  const [sendVerifyOTP, { }] = useVerifyOTPMutation()
  const [createUser, { }] = useRegisterMutation()
  const [additional_message, setAdditionalMessage] = useState()
  const navigate = useNavigate()
  const dispatch = useDispatch()

  useEffect(() => {

    if (email !== "" && !email.includes("@gmail.com")) {
      setAvailableMessage({ message: "NotAvailable", error: true })
      return
    } else if (email !== "") {
      setAvailableMessage({ message: "Available", error: false })
    }

  }, [email])
  useEffect(() => {
    if (password === "" || confirmPassword === "") {
      setErrorMessage("")
      return
    }

    if (confirmPassword !== password) {
      setErrorMessage("Passwords Donot Match")
    } else {
      setErrorMessage("")
    }

  }, [password, confirmPassword])

  useEffect(() => {
    const handleResize = () => {
      if (window.innerWidth >= 768) {
        document.body.style.overflowX = "hidden";
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

  const toggleConfirmPasswordVisibility = () => {
    setConfirmPasswordVisible((prev) => !prev);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (availableMessage.error) {
      setErrorMessage("email is not available or not valid");
      return
    }
    setErrorMessage("");

    if (password !== confirmPassword) {
      setErrorMessage("Passwords do not match.");
      return;
    }

    setStep(2)

    setAdditionalMessage("sending email...")
    const data = await sendEmail({ email: email })

    if (data.error) {
      setErrorMessage(data.error.message.data)
      setStep(0)
      return
    }

    setAdditionalMessage("Email Sent")

  };

  const handleOTPSubmit = async (e) => {

    e.preventDefault()
    console.log(otp)
    const data = await sendVerifyOTP({
      email: email,
      otp: Number(otp)
    })

    if (data.error) {
      setStep(0)
      setErrorMessage(data.error)
      return
    }

    const udata = await createUser({
      email: email,
      password: password,
      handshake: data.data.data.handshake
    })

    if (udata.error) {
      setStep(0)
      setErrorMessage(udata.error.message.data)
      return
    }

    const username = udata.data.data.username
    // navigate to profile page
    navigate(`/u/${username}`)
  }


  const handleInputChange = (e) => {
    const value = e.target.value.replace(/\D/g, "").slice(0, 4); // Allow only numbers, max 4 digits
    setOTP(value);
  };

  return (

    < div className="relative h-screen w-full flex lg:flex-row flex-col-reverse overflow-scroll" >

      {step === 2 ? (
        <>
          <CenteredModal value={otp} setValue={handleInputChange} onClose={() => { setStep(1) }} setSubmit={handleOTPSubmit} additional_message={additional_message} />
        </>) : (<></>)}
      {/* Half Image Section (Visible on large screens) */}
      < div className="w-full lg:w-1/2 lg:block sm:hidden relative" >
        <img
          src="/login_page.png"
          alt="Half Image"
          className="object-cover h-full w-full"
          style={{ objectPosition: "center top" }}
        />
      </div >
      {/* Register Form Section */}
      < div className="w-full lg:w-1/2 overflow-scroll flex items-center justify-center bg-gray-50 lg:bg-transparent relative z-10 lg:z-auto" >
        <div className="bg-white p-8 lg:p-12 rounded-lg shadow-lg w-full max-w-lg lg:mt-[-4rem] mt-6 md:mt-10"> {/* Adjusted margin-top here */}
          <h2 className="text-4xl font-bold text-center mb-4">Register</h2>
          <p className="text-center text-gray-500 mb-6">Create your account...</p>


          <form onSubmit={handleSubmit}>
            {/* Username/Email */}
            <div className="mb-5">
              <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="username">
                Email
              </label>
              <input
                id="email"
                type="email"
                className="w-full px-4 py-3 rounded-lg bg-gray-200 mt-2 border focus:border-blue-500 focus:bg-white focus:outline-none"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />

              <div className={`${availableMessage.error ? 'text-red-400' : 'text-green-400'}`}>{availableMessage.message}</div>
            </div>

            {/* Password */}
            <div className="mb-6">
              <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
                Password
              </label>
              <div className="relative flex items-center">
                <input
                  id="password"
                  type={passwordVisible ? "text" : "password"} // Toggle input type based on state
                  value={password}
                  onChange={(e) => setPassword(e.target.value)} // Update password state
                  className="w-full px-4 py-3 rounded-lg bg-gray-200 mt-2 border focus:border-blue-500 focus:bg-white focus:outline-none"
                  required
                />
                <button
                  type="button"
                  onClick={togglePasswordVisibility}
                  className="absolute right-2 top-1/2 transform -translate-y-1/2 focus:outline-none"
                  style={{ margin: '0.5rem' }} // Added margin here
                >
                  <FontAwesomeIcon icon={passwordVisible ? faEyeSlash : faEye} className="text-gray-500 h-5 w-5" />
                </button>
              </div>
            </div>

            {/* Confirm Password */}
            <div className="mb-6">
              <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="confirmPassword">
                Confirm Password
              </label>
              <div className="relative flex items-center">
                <input
                  id="confirmPassword"
                  type={confirmPasswordVisible ? "text" : "password"} // Toggle input type based on state
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)} // Update confirm password state
                  className="w-full px-4 py-3 rounded-lg bg-gray-200 mt-2 border focus:border-blue-500 focus:bg-white focus:outline-none"
                  required
                />
                <button
                  type="button"
                  onClick={toggleConfirmPasswordVisibility}
                  className="absolute right-2 top-1/2 transform -translate-y-1/2 focus:outline-none"
                  style={{ margin: '0.5rem' }} // Added margin here
                >
                  <FontAwesomeIcon icon={confirmPasswordVisible ? faEyeSlash : faEye} className="text-gray-500 h-5 w-5" />
                </button>
              </div>
            </div>

            {errorMessage && <p className="text-red-500 text-center mb-4">{errorMessage}</p>}
            {/* Sign Up Button */}
            <button
              type="submit"
              className="w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 rounded-lg transition duration-200"
            >
              Register
            </button>

            {/* Already have an account Link */}
            <p className="mt-5 text-center text-gray-500">
              Already have an Account?{" "}
              <a href="/login" className="text-blue-500 hover:underline">
                Login
              </a>
            </p>
          </form>
        </div>
      </div >
    </div >
  );
};

export default Register;
