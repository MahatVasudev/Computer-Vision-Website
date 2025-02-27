import { useState } from "react"
import CenteredContainer from "../components/SetupContainer"
import InputTextField from "../components/InputTextFields";
import SelectField from "../components/SelectField";

const SetupProfile = () => {
  const [username, setUsername] = useState("")
  const max_step = 4
  const min_step = 1
  const [step, setStep] = useState(min_step)
  const [firstName, setFirstName] = useState("")
  const [gender, setGender] = useState("")
  const [lastName, setLastName] = useState("")
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
  return (
    <CenteredContainer prev_fn={() => prevStep()} next_fn={() => nextStep()}>
      {
        step === 1 && (<>
          <h2 className="mx-auto font-bold self-start text-2xl mb-5"> What do you want people to call you? </h2>
          <InputTextField label={"Username"} id="username" state={{ state: username, setState: setUsername }} />
        </>
        )
      } {
        step === 2 && (<>
          <h2 className="mx-auto font-bold self-start text-2xl mb-5">Whats Your Name?</h2>
          <InputTextField label={"First Name"} id="first_name" state={{ state: firstName, setState: setFirstName }} />
          <InputTextField label={"Last Name"} id="last_name" state={{ state: lastName, setState: setLastName }} />
        </>)
      } {
        step === 3 && (<>
          <h2 className="mx-auto font-bold self-start text-2xl mb-5">General Questions</h2>
          <p>Gender</p>
          <SelectField
            state={{ state: gender, setState: setGender }}
            options={[
              { value: '', label: 'Select Gender' },
              { value: 'male', label: "Male (he/him)" },
              { value: 'female', label: "Female (she/her)" },
              { value: 'none', label: "Don't want to specify" }]} />
          <p>Birth Year</p>

          <div className="mb-5">
            <input
              type="text"
              maxLength="4"
              className="mt-4 w-24 text-center border border-gray-300 rounded-lg p-2 text-xl outline-none focus:border-blue-500"
              placeholder="XXXX"
            />
          </div>
          <p>Preferred Mode</p>
          <p>Preferred Color</p>
          <div className="relative w-64">
            {/* Selected Color Display */}
            <div
              className="w-full h-12 flex items-center justify-between px-4 border border-gray-700 rounded-lg cursor-pointer"
              style={{ backgroundColor: selectedColor }}
            >
              <span className="text-white">{selectedColor.toUpperCase()}</span>
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
                value={selectedColor}
                onChange={(e) => handleColorChange(e.target.value)}
              />
            </div>
          </div>
        </>)
      }
    </CenteredContainer>
  )
}


export default SetupProfile
