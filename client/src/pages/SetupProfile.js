import { useEffect, useState } from "react"
import CenteredContainer from "../components/SetupContainer"
import InputTextField from "../components/InputTextFields";
import SelectField from "../components/SelectField";
import { Moon, Sun } from "lucide-react";
import { setSettings, toggleThemes } from "../features/settings/settings.local";
import { useDispatch, useSelector } from "react-redux";
import { useSetupMutation } from "../features/user/auth";
import { useNavigate } from "react-router-dom";
import { useSeeUsernameExistsMutation } from "../features/user/user";
const SetupProfile = () => {
  const [username, setUsername] = useState("")
  const max_step = 3
  const min_step = 1
  const dispatch = useDispatch()
  const [step, setStep] = useState(min_step)
  const [firstName, setFirstName] = useState("")
  const [gender, setGender] = useState("")
  const [year, setYear] = useState()
  const theme = useSelector((state) => state.settings.dark)
  const preferedColor = useSelector((state) => state.settings.prefered_color)
  const [lastName, setLastName] = useState("")
  const [setup, { isLoading }] = useSetupMutation()
  const navigate = useNavigate()
  const [checkuser] = useSeeUsernameExistsMutation()
  const [checkMessage, setCheckMessage] = useState({})
  useEffect(() => {

    if (username.length <= 4 || username.length >= 20) {
      setCheckMessage({ message: "invalid username", error: true })
      return
    }


    setCheckMessage({ message: "", error: true })
  }, [username])

  const checkusername = async () => {
    const response = await checkuser({ username })

    if (!response.error) {
      setCheckMessage({ message: "username already taken!!", error: true })
      return
    }

    setCheckMessage({ message: "username available!!!", error: false })
  }
  document.getElementById("check-availability")?.addEventListener("click", () => {

    console.log()
    if (checkMessage.message === "" || checkMessage.error === false) {

      checkusername()
    }
  })



  const prevStep = () => {
    if (step > min_step) {
      setStep(step - 1)
    }

  }



  const nextStep = () => {
    console.log(step)
    if (step < max_step) {
      setStep(step + 1)
    }
  }


  const colors = ["#000000"]
  const [selectedColor, setSelectedColor] = useState(colors[0] || "#000000");


  const handleColorChange = (color) => {
    setSelectedColor(color);
  };


  const handleInputChange = (e) => {
    const value = e.target.value.replace(/\D/g, "").slice(0, 4); // Allow only numbers, max 4 digits
    setYear(value);
  };

  const submit_query = async () => {

    const data = {
      username,
      firstname: firstName,
      lastname: lastName === "" ? "*" : lastName,
      gender: gender,
      darkmode: theme,
      preferedColor: selectedColor,
      birthyear: Number(year)
    }

    console.log(data)
    try {
      const response = await setup(data)

      if (response.error) {
        console.log(response.error)
        alert("Error Faced: ", response.error.message?.error)
        return
      }
      navigate(`/u/${username}`)
    } catch (err) {

      alert("Error Faced: ", err)
    }


  }
  return (
    <CenteredContainer prev_fn={() => prevStep()} next_fn={() => nextStep()} step={step} finish_fn={() => submit_query()} final_step={max_step}>
      {
        step === 1 && (<>
          <h2 className="mx-auto font-bold self-start text-2xl mb-5"> What do you want people to call you? </h2>
          <InputTextField label={"Username"} id="username" state={{ state: username, setState: setUsername }} />
          <button id="check-availability" className="bg-green-400 rounded-[30%] text-white dark:border-white border-black border-2 font-bold p-3 mb-auto mt-1 ml-auto mr-5">
            check
          </button>
          <p className={checkMessage.error ? `text-red-400` : `text-green-400`}> {checkMessage?.message}</p>
        </>
        )
      } {
        step === 2 && (<>
          <h2 className="mx-auto font-bold self-start text-2xl mb-5">Whats Your Name?</h2>
          <InputTextField label={"First Name"} id="first_name" state={{ state: firstName, setState: setFirstName }} />
          <InputTextField label={"Last Name"} id="last_name" state={{ state: lastName, setState: setLastName }} />
        </>)
      } {
        step === 3 && (<div className="flex-1 flex-col">
          <h2 className="mx-auto font-bold self-start text-2xl mb-5">General Questions</h2>
          <p>Gender</p>
          <SelectField
            state={{ state: gender, setState: setGender }}
            options={[
              { value: '', label: 'Select Gender' },
              { value: 'M', label: "Male (he/him)" },
              { value: 'F', label: "Female (she/her)" },
              { value: 'NS', label: "Don't want to specify" }]} />
          <p>Birth Year</p>

          <div className="mb-5">
            <input
              type="text"
              maxLength="4"
              value={year}
              onChange={handleInputChange}
              className="mt-4 w-24 text-center border border-gray-300 rounded-lg p-2 text-xl outline-none focus:border-blue-500"
              placeholder="XXXX"
            />
          </div>
          <p>Preferred Mode</p>
          <button
            onClick={() => dispatch(toggleThemes())}
            className="p-2 rounded-full bg-gray-200 dark:bg-gray-700"
          >
            {theme === 0 ? <Sun size={20} /> : <Moon size={20} />}
          </button>
          <p>Preferred Color</p>
          <div className="relative w-64">
            {/* Selected Color Display */}
            <div
              className="w-full h-12 flex items-center justify-between px-4 border border-gray-700 rounded-lg cursor-pointer"
              style={{ backgroundColor: preferedColor }}
            >
              <span className="text-white">{preferedColor.toUpperCase()}</span>
            </div>

            {/* Color Options */}
            <div className="absolute mt-2 w-full bg-gray-900 p-3 rounded-lg shadow-lg flex flex-wrap gap-2">
              {colors.map((color, index) => (
                <div
                  key={index}
                  className="w-8 h-8 rounded-full cursor-pointer border-2 border-transparent hover:border-white"
                  style={{ backgroundColor: color }}
                  onClick={() => handleColorChange(color)}
                />
              ))}

              {/* Custom Color Input */}
              <input
                type="color"
                className="w-full h-10 cursor-pointer border-none outline-none bg-transparent"
                value={preferedColor}
                onChange={(e) => dispatch(setSettings({ key: "prefered_color", value: e.target.value }))}
              />
            </div>
          </div>
        </div>)
      }
    </CenteredContainer>
  )
}


export default SetupProfile
